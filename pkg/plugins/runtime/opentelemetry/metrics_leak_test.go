package opentelemetry

import (
	"fmt"
	"net"
	"runtime"
	"sync/atomic"
	"testing"
	"time"

	"github.com/go-logr/logr"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/kumahq/kuma/v2/pkg/core/runtime/component"
)

// TestMetricsPusherShutdownErrorCausesRestart demonstrates that metricsPusher.Start()
// returns shutdown errors instead of logging them (like the tracer does). When wrapped
// in ResilientComponent, this causes an unnecessary restart loop after the stop signal.
func TestMetricsPusherShutdownErrorCausesRestart(t *testing.T) {
	// Use an unreachable endpoint so the final export during shutdown fails.
	// grpc.NewClient is lazy, so exporter creation succeeds. The error surfaces
	// only when PeriodicReader.Shutdown tries to do a final export.
	t.Setenv("OTEL_EXPORTER_OTLP_ENDPOINT", "http://127.0.0.1:1")
	t.Setenv("OTEL_EXPORTER_OTLP_INSECURE", "true")

	registry := prometheus.NewRegistry()
	counter := prometheus.NewCounter(prometheus.CounterOpts{
		Name: "test_metric_total",
		Help: "a metric so the gatherer has something to export",
	})
	registry.MustRegister(counter)
	counter.Inc()

	mp := &metricsPusher{
		gatherer: registry,
		log:      logr.Discard(),
	}

	// Pre-close the stop channel. This simulates what happens during a
	// ResilientComponent restart: the stop channel is already closed from the
	// previous iteration, so Start() creates all resources, immediately reads
	// from stop, and then attempts shutdown.
	stop := make(chan struct{})
	close(stop)

	err := mp.Start(stop)

	if err != nil {
		// This is the bug: Start() should log the shutdown error and return nil
		// (like the tracer component in plugin.go does). Returning an error tells
		// ResilientComponent to retry, creating a pointless restart loop.
		t.Logf("BUG CONFIRMED: Start() returned error on shutdown: %v", err)
		t.Logf("The tracer in the same package logs shutdown errors and returns nil.")
		t.Logf("This causes ResilientComponent to call Start() again even though the CP is stopping.")
	} else {
		t.Log("Start() returned nil — no issue (shutdown export succeeded, e.g. no metrics to flush)")
	}
}

// TestMetricsPusherGoroutineLeakOnRestartCycles demonstrates that repeated
// Start()/shutdown cycles (as triggered by ResilientComponent restarts) accumulate
// goroutines. Each cycle creates a new gRPC exporter, PeriodicReader (which spawns
// a background goroutine), and MeterProvider.
func TestMetricsPusherGoroutineLeakOnRestartCycles(t *testing.T) {
	// Start a TCP listener that accepts connections but speaks no gRPC.
	// This makes the export hang until the shutdown context expires (5s),
	// leaving the gRPC connection in a half-open state.
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer ln.Close()

	// Accept connections in the background so the TCP handshake succeeds
	// but the gRPC/HTTP2 handshake never completes.
	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				return
			}
			// Hold the connection open, never respond
			go func(c net.Conn) {
				buf := make([]byte, 1024)
				for {
					_, err := c.Read(buf)
					if err != nil {
						return
					}
				}
			}(conn)
		}
	}()

	t.Setenv("OTEL_EXPORTER_OTLP_ENDPOINT", fmt.Sprintf("http://%s", ln.Addr().String()))
	t.Setenv("OTEL_EXPORTER_OTLP_INSECURE", "true")

	registry := prometheus.NewRegistry()
	counter := prometheus.NewCounter(prometheus.CounterOpts{
		Name: "test_leak_total",
		Help: "test counter",
	})
	registry.MustRegister(counter)
	counter.Inc()

	// Warm up: let the runtime settle
	runtime.GC()
	time.Sleep(200 * time.Millisecond)
	goroutinesBefore := runtime.NumGoroutine()

	const cycles = 3
	for i := 0; i < cycles; i++ {
		mp := &metricsPusher{
			gatherer: registry,
			log:      logr.Discard(),
		}

		stop := make(chan struct{})
		errCh := make(chan error, 1)
		go func() {
			errCh <- mp.Start(stop)
		}()

		// Let the PeriodicReader goroutine start
		time.Sleep(200 * time.Millisecond)

		// Signal stop — triggers shutdown with 5s context
		close(stop)

		select {
		case startErr := <-errCh:
			t.Logf("cycle %d: Start() returned: %v", i+1, startErr)
		case <-time.After(10 * time.Second):
			t.Fatalf("cycle %d: Start() did not return within 10s", i+1)
		}
	}

	// Give background goroutines a moment
	time.Sleep(time.Second)
	goroutinesAfter := runtime.NumGoroutine()
	leaked := goroutinesAfter - goroutinesBefore

	t.Logf("goroutines: before=%d after=%d delta=%+d", goroutinesBefore, goroutinesAfter, leaked)

	if leaked > 2 {
		t.Errorf("goroutine leak: %d goroutines accumulated over %d restart cycles", leaked, cycles)
		t.Log("Each cycle creates a gRPC connection and PeriodicReader goroutine.")
		t.Log("Without defer-based cleanup, these may not be released when shutdown fails.")
	}
}

// TestMetricsPusherResilientComponentRestartLoop demonstrates the full scenario:
// metricsPusher wrapped in ResilientComponent enters a restart loop when the
// OTLP endpoint is unreachable, because Start() returns shutdown errors.
//
// Each restart cycle takes ~5s (the hardcoded shutdown timeout), so we need to
// wait long enough to observe at least one restart.
func TestMetricsPusherResilientComponentRestartLoop(t *testing.T) {
	t.Setenv("OTEL_EXPORTER_OTLP_ENDPOINT", "http://127.0.0.1:1")
	t.Setenv("OTEL_EXPORTER_OTLP_INSECURE", "true")

	registry := prometheus.NewRegistry()
	counter := prometheus.NewCounter(prometheus.CounterOpts{
		Name: "test_restart_total",
		Help: "test counter",
	})
	registry.MustRegister(counter)
	counter.Inc()

	var startCount atomic.Int32

	// Wrap metricsPusher in a factory that counts Start() calls.
	// Each call creates a fresh metricsPusher (like ResilientComponent does
	// when it retries by calling the same component.Start() again).
	wrapper := &startCounter{
		inner: &metricsPusher{
			gatherer: registry,
			log:      logr.Discard(),
		},
		count: &startCount,
	}

	resilient := component.NewResilientComponent(
		logr.Discard(),
		wrapper,
		50*time.Millisecond,  // fast backoff for the test
		200*time.Millisecond, // short max backoff
	)

	stop := make(chan struct{})
	errCh := make(chan error, 1)
	go func() {
		errCh <- resilient.Start(stop)
	}()

	// Let the first Start() begin and block on <-stop.
	time.Sleep(200 * time.Millisecond)

	// Signal stop - this triggers the shutdown path.
	close(stop)

	// Wait long enough for the first Start() to fail shutdown (~5s) and for
	// ResilientComponent to attempt at least one restart. Each restart also
	// takes ~5s. Total: ~12s should see 1-2 restarts.
	select {
	case <-errCh:
	case <-time.After(20 * time.Second):
		t.Fatal("ResilientComponent did not stop within 20s")
	}

	count := startCount.Load()
	t.Logf("Start() was called %d times", count)

	if count > 1 {
		t.Logf("BUG CONFIRMED: metricsPusher was restarted %d times.", count)
		t.Log("The first Start() call is expected. Subsequent calls are unnecessary restarts")
		t.Log("caused by Start() returning the shutdown error instead of nil.")
		t.Log("The tracer in the same package avoids this by logging shutdown errors and returning nil.")
	} else {
		// ResilientComponent uses a random select between <-stop and errCh.
		// Sometimes <-stop wins and no restart occurs. This is not a test
		// failure - the bug is proven by TestMetricsPusherShutdownErrorCausesRestart.
		t.Log("No restart observed (ResilientComponent selected <-stop before errCh).")
		t.Log("The shutdown error bug is still present — see TestMetricsPusherShutdownErrorCausesRestart.")
	}
}

// startCounter wraps a Component and counts how many times Start() is called.
type startCounter struct {
	inner component.Component
	count *atomic.Int32
}

func (s *startCounter) Start(stop <-chan struct{}) error {
	s.count.Add(1)
	return s.inner.Start(stop)
}

func (s *startCounter) NeedLeaderElection() bool {
	return s.inner.NeedLeaderElection()
}
