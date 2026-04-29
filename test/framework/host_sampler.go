package framework

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/onsi/ginkgo/v2"

	"github.com/kumahq/kuma/v2/test/framework/report"
)

func init() {
	// Register the suite-end baseline dumper. Runs whether or not any spec
	// failed — that's the whole point: a successful run still produces a
	// host-samples directory you can diff against a failing run.
	report.PostDumpHook = func(r ginkgo.Report) {
		DumpHostSamplesBaseline(r.SuiteDescription)
	}
}

// Why this exists:
//
// The post-mortem bundle (see debug.go) samples cluster state once, after the
// test has failed. When the failure mode is a multi-minute pause of the k3d
// node container (host overcommit, cgroup stall, kernel softlockup), every
// signal we collect inside that container is taken AFTER the pause ends, so
// the pause itself is invisible. We see a gap in timestamps and nothing else.
//
// This sampler runs OUT-OF-BAND on the test runner host. It does not depend on
// the k3d container making forward progress, so it keeps producing data while
// the container is frozen. Three signals matter most:
//
//   - "docker stats" per k3d container — proves whether the cgroup got CPU
//     cycles during the gap (host's view, not the frozen process's view).
//   - Host PSI (/proc/pressure/{cpu,memory,io}) — quantifies whether the gap
//     coincided with system-wide pressure.
//   - vmstat — procs in R vs B, free/swap, ctxt/s, %sys/%idle. Cheap, dense.
//
// Output goes to a per-process tmpdir; DumpHostSamplesTo copies them into the
// bundle when a test fails.

var (
	hostSamplerOnce   sync.Once
	hostSamplerCancel context.CancelFunc
	hostSamplerDir    string
)

// HostSamplerBaselineDir is where the suite-end snapshot of the rolling
// sampler files is written, regardless of pass/fail. It deliberately lives
// outside report.BaseDir so it isn't moved aside by DumpReport's startup
// rename. The CI artifact upload pulls this path explicitly.
const HostSamplerBaselineDir = "build/host-samples-baseline"

// hostSamplerInterval is the polling cadence for the in-process samplers
// (docker stats, host PSI). vmstat/iostat self-pace at 1s. Picked so a 4-minute
// freeze produces ~120 samples per signal — enough resolution to localise the
// edge of the gap.
const hostSamplerInterval = 2 * time.Second

// StartHostSampler starts the background sampling goroutines. Idempotent and
// safe to call from any cluster's lifecycle hook. The goroutines are not
// reaped explicitly; they exit when the test process exits.
func StartHostSampler() {
	hostSamplerOnce.Do(func() {
		dir, err := os.MkdirTemp("", fmt.Sprintf("kuma-e2e-samples-%d-*", os.Getpid()))
		if err != nil {
			Logf("[host-sampler] failed to create sample dir: %v", err)
			return
		}
		hostSamplerDir = dir
		Logf("[host-sampler] writing samples to %s", hostSamplerDir)

		ctx, cancel := context.WithCancel(context.Background())
		hostSamplerCancel = cancel

		// vmstat -n 1: timestamped 1s samples of procs/memory/swap/cpu.
		// -n suppresses repeating headers so the file is straightforward to grep.
		go runLongLived(ctx, filepath.Join(hostSamplerDir, "vmstat.txt"),
			"vmstat", "-t", "-n", "1")

		// iostat -x 1: per-device IO util%, await, queue depth. Disk stalls
		// during cluster bring-up (e2e clusters write a lot to the kine DB and
		// container layer storage) show up here.
		go runLongLived(ctx, filepath.Join(hostSamplerDir, "iostat.txt"),
			"iostat", "-x", "-t", "1")

		// Periodic samplers: docker stats and host PSI.
		go pollLoop(ctx, filepath.Join(hostSamplerDir, "docker-stats.tsv"),
			hostSamplerInterval, sampleDockerStats)

		go pollLoop(ctx, filepath.Join(hostSamplerDir, "host-pressure.txt"),
			hostSamplerInterval, sampleHostPressure)

		// Per-k3d-container PSI from inside the container, via docker exec.
		// Each docker exec is a separate sample so a frozen container just
		// causes a single sample to be missed, not the whole stream to stall.
		go pollLoop(ctx, filepath.Join(hostSamplerDir, "k3d-pressure.txt"),
			hostSamplerInterval, sampleK3dPressure)
	})
}

// DumpHostSamplesTo copies the rolling sample files into the report bundle.
// Called from DumpState. The same files are emitted for every cluster in a
// failed test — they're host-wide, not cluster-scoped — but cluster-scoped
// paths keep the bundle layout consistent.
func DumpHostSamplesTo(cluster Cluster) {
	if hostSamplerDir == "" {
		return
	}
	entries, err := os.ReadDir(hostSamplerDir)
	if err != nil {
		Logf("[host-sampler] read dir %s: %v", hostSamplerDir, err)
		return
	}
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		full := filepath.Join(hostSamplerDir, e.Name())
		data, err := os.ReadFile(full)
		if err != nil {
			Logf("[host-sampler] read %s: %v", full, err)
			continue
		}
		report.AddFileToReportEntry(path.Join(cluster.Name(), "host-samples", e.Name()), data)
	}
}

// HostSamplerMark writes a marker line into every rolling sampler file. Used
// at spec boundaries so a successful run can be sliced out and compared
// against a failing run from the same suite. Cheap (one fprintf per file).
func HostSamplerMark(label string) {
	if hostSamplerDir == "" {
		return
	}
	entries, err := os.ReadDir(hostSamplerDir)
	if err != nil {
		return
	}
	now := time.Now().UTC().Format(time.RFC3339Nano)
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		full := filepath.Join(hostSamplerDir, e.Name())
		f, err := os.OpenFile(full, os.O_APPEND|os.O_WRONLY, 0)
		if err != nil {
			continue
		}
		fmt.Fprintf(f, "### MARK %s %s\n", now, label)
		_ = f.Close()
	}
}

// DumpHostSamplesBaseline snapshots the current rolling sampler files into
// HostSamplerBaselineDir, regardless of test outcome. Intended to be called
// once per suite end so a fully-green run still produces a comparison
// baseline. Subsequent calls overwrite — the file is the full history of the
// process so the latest snapshot is always the most complete.
func DumpHostSamplesBaseline(suiteName string) {
	if hostSamplerDir == "" {
		return
	}
	dst := filepath.Join(HostSamplerBaselineDir, sanitizeFilename(suiteName))
	if err := os.MkdirAll(dst, 0o755); err != nil {
		Logf("[host-sampler] mkdir %s: %v", dst, err)
		return
	}
	entries, err := os.ReadDir(hostSamplerDir)
	if err != nil {
		Logf("[host-sampler] read dir %s: %v", hostSamplerDir, err)
		return
	}
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		full := filepath.Join(hostSamplerDir, e.Name())
		data, err := os.ReadFile(full)
		if err != nil {
			continue
		}
		_ = os.WriteFile(filepath.Join(dst, e.Name()), data, 0o644)
	}
}

func sanitizeFilename(s string) string {
	if s == "" {
		return "suite"
	}
	var b strings.Builder
	for _, r := range s {
		switch {
		case r >= 'a' && r <= 'z', r >= 'A' && r <= 'Z', r >= '0' && r <= '9', r == '-', r == '_', r == '.':
			b.WriteRune(r)
		default:
			b.WriteRune('_')
		}
	}
	return b.String()
}

// runLongLived starts a long-running command and pipes stdout to a file until
// the context is cancelled. Used for vmstat/iostat, which self-pace.
func runLongLived(ctx context.Context, outPath string, name string, args ...string) {
	f, err := os.Create(outPath)
	if err != nil {
		Logf("[host-sampler] create %s: %v", outPath, err)
		return
	}
	defer f.Close()
	cmd := exec.CommandContext(ctx, name, args...)
	cmd.Stdout = f
	cmd.Stderr = f
	if err := cmd.Start(); err != nil {
		// vmstat/iostat may be missing on a developer machine; record once
		// and move on. Don't fail the test for a missing diagnostics tool.
		fmt.Fprintf(f, "failed to start %s: %v\n", name, err)
		return
	}
	_ = cmd.Wait()
}

// pollLoop runs sample at hostSamplerInterval and appends results to outPath.
// Each line is prefixed with an RFC3339Nano timestamp to make correlation with
// the k3s/CP/kuma-init timestamps in the bundle straightforward.
func pollLoop(ctx context.Context, outPath string, interval time.Duration, sample func() string) {
	f, err := os.Create(outPath)
	if err != nil {
		Logf("[host-sampler] create %s: %v", outPath, err)
		return
	}
	defer f.Close()
	t := time.NewTicker(interval)
	defer t.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case now := <-t.C:
			out := sample()
			if out == "" {
				continue
			}
			fmt.Fprintf(f, "=== %s ===\n%s\n", now.UTC().Format(time.RFC3339Nano), out)
			if err := f.Sync(); err != nil {
				return
			}
		}
	}
}

// sampleDockerStats: one snapshot of CPU%/mem%/blockio/netio per k3d-*
// container. --no-stream returns immediately; --format keeps the output
// machine-parseable.
func sampleDockerStats() string {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	out, err := exec.CommandContext(ctx, "docker", "stats", "--no-stream",
		"--format", "{{.Name}}\t{{.CPUPerc}}\t{{.MemUsage}}\t{{.MemPerc}}\t{{.BlockIO}}\t{{.NetIO}}\t{{.PIDs}}",
	).Output()
	if err != nil {
		return fmt.Sprintf("docker stats failed: %v", err)
	}
	// Filter to k3d-* lines to keep the file focused.
	var keep []string
	for _, line := range strings.Split(string(out), "\n") {
		if strings.HasPrefix(line, "k3d-") {
			keep = append(keep, line)
		}
	}
	return strings.Join(keep, "\n")
}

// sampleHostPressure reads the host's PSI counters. Available on cgroup-v2
// hosts (Linux 4.20+, all current GitHub runners). Each counter has avg10/60/300
// percentages plus a monotonic "total" of stalled microseconds — the deltas
// between samples tell us how much wall-clock time was lost to pressure.
func sampleHostPressure() string {
	var b strings.Builder
	for _, p := range []string{"/proc/pressure/cpu", "/proc/pressure/memory", "/proc/pressure/io"} {
		data, err := os.ReadFile(p)
		if err != nil {
			fmt.Fprintf(&b, "--- %s ---\n%v\n", p, err)
			continue
		}
		fmt.Fprintf(&b, "--- %s ---\n%s", p, string(data))
	}
	loadavg, err := os.ReadFile("/proc/loadavg")
	if err == nil {
		fmt.Fprintf(&b, "--- /proc/loadavg ---\n%s", string(loadavg))
	}
	return b.String()
}

// sampleK3dPressure reads PSI from inside each k3d-*-server-0 container.
// In cgroup-v2, the container's view of /proc/pressure/* reflects the host's
// counters, but reading them via docker exec also tells us whether the
// container is responsive at all — a hung exec is itself a signal.
func sampleK3dPressure() string {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	listOut, err := exec.CommandContext(ctx, "docker", "ps", "--filter", "name=k3d-",
		"--filter", "status=running", "--format", "{{.Names}}").Output()
	if err != nil {
		return fmt.Sprintf("docker ps failed: %v", err)
	}
	var b strings.Builder
	for _, name := range strings.Fields(string(listOut)) {
		execCtx, execCancel := context.WithTimeout(context.Background(), 3*time.Second)
		out, err := exec.CommandContext(execCtx, "docker", "exec", name, "sh", "-c",
			"cat /proc/pressure/cpu /proc/pressure/memory /proc/pressure/io /proc/loadavg /proc/stat 2>/dev/null | head -40",
		).CombinedOutput()
		execCancel()
		fmt.Fprintf(&b, "--- %s ---\n", name)
		if err != nil {
			// A timeout here is itself the signal we're after — record it
			// instead of dropping the sample.
			fmt.Fprintf(&b, "exec error: %v\n", err)
		}
		if len(out) > 0 {
			b.Write(out)
			if !strings.HasSuffix(string(out), "\n") {
				b.WriteString("\n")
			}
		}
	}
	return b.String()
}
