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
	// failed - that's the whole point: a successful run still produces a
	// host-samples directory you can diff against a failing run.
	//
	// We deliberately do NOT cancel the sampler context here. A `make test/e2e`
	// invocation runs ginkgo across multiple suites in one binary process and
	// each suite ends with its own ReportAfterSuite -> DumpReport -> PostDumpHook
	// chain. If we cancel after the first suite's PostDumpHook fires, all
	// later suites run with no host sampling, so any failure that happens in
	// suite N>1 has no rolling host-samples data to inspect. (Observed on
	// PR #16455 run 25309274732: docker-stats stops at 09:02:37 from the
	// Federation_Suite end, but the Helm-suite failure is at 09:06:43 and
	// has no host data.) The cancel was needed previously to kill long-lived
	// vmstat/iostat subprocesses; those are in-process pollers now, so the
	// goroutines exit naturally on test-binary exit.
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
//   - "docker stats" per k3d container - proves whether the cgroup got CPU
//     cycles during the gap (host's view, not the frozen process's view).
//   - Host PSI (/proc/pressure/{cpu,memory,io}) - quantifies whether the gap
//     coincided with system-wide pressure.
//   - vmstat-equivalent samples from /proc/stat + /proc/meminfo + /proc/vmstat
//     + /proc/loadavg: procs running/blocked, ctxt switches, free/swap, paging.
//
// All samplers are in-process (no subprocess fork). Earlier revisions spawned
// `vmstat -t -n 1` / `iostat -x -t 1` as long-lived child processes; on a
// clean GitHub Actions runner those children inherit MAKE's jobserver pipe
// FDs (3 and 4) and survive the Go test binary's exit. Make then waits on the
// orphaned pipe FDs forever, which manifests as a 20-30 minute hang between
// the last "make: Leaving directory" line and the 60-minute job timeout. The
// in-process samplers below avoid that whole class of bug because no process
// is ever forked - the file descriptors stay inside the Go runtime.
//
// Output goes to a per-process tmpdir; DumpHostSamplesTo copies them into the
// bundle when a test fails.

var (
	hostSamplerOnce   sync.Once
	hostSamplerCancel context.CancelFunc
	hostSamplerDir    string
)

// hostSamplerBaselineDir resolves the suite-end snapshot directory at the
// moment of writing, deriving it from report.BaseDir so it follows the same
// $KUMA_DUMP_DIR-based path the rest of the bundle uses. It deliberately
// lives outside report.BaseDir so it isn't moved aside by DumpReport. The
// CI artifact upload pulls "build/host-samples-baseline" explicitly, so
// this must resolve to that absolute path on the runner.
//
// Layout assumption: report.BaseDir = "<workspace>/build/reports/e2e-debug"
// (set by config.go from $KUMA_DUMP_DIR). Going up two parents gives
// "<workspace>/build", which is where the artifact upload looks.
//
// Falls back to a relative path if report.BaseDir is empty (developer run
// without $KUMA_DUMP_DIR set), so local invocations still produce something.
func hostSamplerBaselineDir() string {
	if report.BaseDir == "" {
		return "build/host-samples-baseline"
	}
	return filepath.Join(filepath.Dir(filepath.Dir(report.BaseDir)), "host-samples-baseline")
}

// hostSamplerInterval is the polling cadence for the in-process samplers
// (docker stats, host PSI). vmstat/iostat self-pace at 1s. Picked so a 4-minute
// freeze produces ~120 samples per signal — enough resolution to localize the
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

		// vmstat-equivalent: read /proc/stat + /proc/loadavg + /proc/meminfo
		// + /proc/vmstat directly each tick. No subprocess - see the comment
		// at the top of this file for why.
		go pollLoop(ctx, filepath.Join(hostSamplerDir, "vmstat.txt"),
			hostSamplerInterval, sampleVmstat)

		// iostat-equivalent: /proc/diskstats per device. Disk stalls during
		// cluster bring-up (e2e clusters write heavily to the kine DB and
		// container layer storage) show up as growing in-flight queue depth
		// and rising weighted-time-in-queue.
		go pollLoop(ctx, filepath.Join(hostSamplerDir, "iostat.txt"),
			hostSamplerInterval, sampleIostat)

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

		// Per-cgroup PSI for kubepods.slice. With CPU requests but no limits
		// we never get cgroup throttling stats — cpu.stat's nr_throttled stays
		// at 0 even when a container is starving on CFS shares. cpu.pressure
		// records that starvation directly: "full avg10=N" is the % of wall
		// time during which every task in the cgroup was waiting for CPU.
		// This is the single-best signal for the "stuck behind everything
		// else on a saturated host" failure mode.
		go pollLoop(ctx, filepath.Join(hostSamplerDir, "k3d-kubepods-pressure.txt"),
			hostSamplerInterval, sampleK3dKubepodsPressure)

		// Top-CPU processes on the host. The aggregate samplers (vmstat,
		// host PSI, docker stats) prove that CPU is saturated during a
		// freeze, but they don't say which process is consuming the cycles.
		// Per-process attribution is the missing piece for picking out the
		// "bloater" - the process whose periodic spike crowds out everyone
		// else's CFS share. Container processes appear in host ps output
		// (docker uses namespaces, not VMs), so one host-level sampler
		// captures k3s, kubelet, calico-node, calico-typha, felix, kuma-cp,
		// kuma-init etc. - distinguishable via the cgroup column.
		go pollLoop(ctx, filepath.Join(hostSamplerDir, "top-procs.txt"),
			hostSamplerInterval, sampleTopProcesses)

		// /proc/schedstat: per-CPU runqueue wait time. Field 7 of every
		// "cpu<N>" line is sum_sched_wait, the cumulative nanoseconds tasks
		// spent on the runqueue waiting to be scheduled. The delta between
		// two samples is the only direct quantification of "tasks were
		// queued and could not get CPU"; PSI's "% wall time stalled" is a
		// derived metric, schedstat is the raw counter.
		go pollLoop(ctx, filepath.Join(hostSamplerDir, "schedstat.txt"),
			hostSamplerInterval, sampleSchedstat)

		// Per-cgroup cpu.stat for each k3d container. Captures CFS throttle
		// counters (nr_throttled, throttled_usec, nr_periods) which complement
		// PSI: PSI tells us tasks waited; cpu.stat tells us whether they
		// waited because of CFS bandwidth control (limits) or shares
		// contention. With CPU requests but no limits the throttle counters
		// should stay at 0; if they don't, our diagnosis is wrong.
		go pollLoop(ctx, filepath.Join(hostSamplerDir, "cpu-stat.txt"),
			hostSamplerInterval, sampleK3dCgroupCPUStat)

		// Per-PID kernel-side wait info for hot or stuck procs. /proc/<pid>/wchan
		// is a one-line label naming the kernel function the task is sleeping
		// in (e.g. "futex_wait_queue", "io_schedule", "do_wait"); /proc/<pid>/stack
		// is the kernel call stack itself. Sampled only for the top-N CPU procs
		// and any D-state proc to keep volume bounded.
		go pollLoop(ctx, filepath.Join(hostSamplerDir, "proc-stack.txt"),
			hostSamplerInterval, sampleProcStack)

		// Per-interface network counters and conntrack table size. Calico's
		// veth pairs and iptables-driven NAT are not visible in docker stats;
		// drops/errors here would localize a calico-vs-flannel divergence to
		// the network stack rather than CPU. conntrack_count vs
		// conntrack_max exhaustion is a known calico failure mode.
		go pollLoop(ctx, filepath.Join(hostSamplerDir, "net-counters.txt"),
			hostSamplerInterval, sampleNetCounters)

		// Per-cgroup io.stat from cgroup-v2. docker stats's BlockIO
		// counter sums to a fraction of /proc/diskstats's total writes
		// (observed: 25 MB across all k3d containers vs sda1 bursts of
		// 40 MB in single 2-second windows), so the bulk of disk pressure
		// is attributed to other slices - dockerd/containerd metadata,
		// systemd-journald, kernel writeback. cgroup-v2 io.stat per slice
		// gives the actual per-cgroup byte/IO counters that docker stats
		// elides.
		go pollLoop(ctx, filepath.Join(hostSamplerDir, "cgroup-io.txt"),
			hostSamplerInterval, sampleCgroupIOStat)

		// Per-PID /proc/<pid>/io. read_bytes / write_bytes attribute the
		// disk traffic in /proc/diskstats back to specific processes -
		// kine inside k3s, dockerd, journald, helm.test, kumactl. Top-N
		// by write_bytes delta tells us who is fsyncing the disk into
		// the ground when IO PSI spikes.
		go pollLoop(ctx, filepath.Join(hostSamplerDir, "proc-io.txt"),
			hostSamplerInterval, sampleProcIO)
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
// hostSamplerBaselineDir(), regardless of test outcome. Intended to be called
// once per suite end so a fully-green run still produces a comparison
// baseline. Subsequent calls overwrite — the file is the full history of the
// process so the latest snapshot is always the most complete.
func DumpHostSamplesBaseline(suiteName string) {
	if hostSamplerDir == "" {
		return
	}
	dst := filepath.Join(hostSamplerBaselineDir(), sanitizeFilename(suiteName))
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
	for line := range strings.SplitSeq(string(out), "\n") {
		if strings.HasPrefix(line, "k3d-") {
			keep = append(keep, line)
		}
	}
	return strings.Join(keep, "\n")
}

// sampleVmstat reads /proc files to produce a vmstat-equivalent sample
// without forking a subprocess. The keys we care about, mapping back to the
// columns vmstat would print:
//
//	procs r/b           - /proc/loadavg's "running/total" field, plus
//	                      procs_running/procs_blocked from /proc/stat
//	memory free/cached  - MemFree, Buffers, Cached, SwapFree from /proc/meminfo
//	swap si/so          - pswpin/pswpout from /proc/vmstat (cumulative; deltas
//	                      between samples = pages/sec swapped in/out)
//	io bi/bo            - pgpgin/pgpgout from /proc/vmstat (sectors read/written)
//	system in/cs        - intr/ctxt from /proc/stat
//	cpu us/sy/id/wa/st  - cpu line from /proc/stat (jiffies; consumer computes deltas)
//
// We emit raw counters and let downstream consumers compute rates - same
// strategy as docker-stats.tsv. One sample per call, no statefulness here.
func sampleVmstat() string {
	var b strings.Builder
	if data, err := os.ReadFile("/proc/loadavg"); err == nil {
		fmt.Fprintf(&b, "loadavg %s", string(data))
	}
	if data, err := os.ReadFile("/proc/stat"); err == nil {
		for line := range strings.SplitSeq(string(data), "\n") {
			switch {
			case strings.HasPrefix(line, "cpu "),
				strings.HasPrefix(line, "ctxt "),
				strings.HasPrefix(line, "intr "),
				strings.HasPrefix(line, "procs_running "),
				strings.HasPrefix(line, "procs_blocked "):
				fmt.Fprintf(&b, "%s\n", line)
			}
		}
	}
	if data, err := os.ReadFile("/proc/meminfo"); err == nil {
		for line := range strings.SplitSeq(string(data), "\n") {
			switch {
			case strings.HasPrefix(line, "MemFree:"),
				strings.HasPrefix(line, "MemAvailable:"),
				strings.HasPrefix(line, "Buffers:"),
				strings.HasPrefix(line, "Cached:"),
				strings.HasPrefix(line, "SwapTotal:"),
				strings.HasPrefix(line, "SwapFree:"),
				strings.HasPrefix(line, "Dirty:"),
				strings.HasPrefix(line, "Writeback:"):
				fmt.Fprintf(&b, "%s\n", line)
			}
		}
	}
	if data, err := os.ReadFile("/proc/vmstat"); err == nil {
		for line := range strings.SplitSeq(string(data), "\n") {
			switch {
			case strings.HasPrefix(line, "pgpgin "),
				strings.HasPrefix(line, "pgpgout "),
				strings.HasPrefix(line, "pswpin "),
				strings.HasPrefix(line, "pswpout "),
				strings.HasPrefix(line, "pgmajfault "),
				strings.HasPrefix(line, "pgfault "):
				fmt.Fprintf(&b, "%s\n", line)
			}
		}
	}
	return b.String()
}

// sampleIostat reads /proc/diskstats for an iostat -x equivalent. Each line
// has 14+ fields (since kernel 4.18); the relevant ones for spotting stalls:
//
//	field 9  - I/Os currently in progress
//	field 10 - time spent doing I/Os (ms; cumulative)
//	field 11 - weighted time spent in queue (ms; cumulative; main stall signal)
//
// Loop and ram devices are skipped to keep the file focused on real disks.
func sampleIostat() string {
	data, err := os.ReadFile("/proc/diskstats")
	if err != nil {
		return fmt.Sprintf("diskstats failed: %v", err)
	}
	var b strings.Builder
	for line := range strings.SplitSeq(string(data), "\n") {
		fields := strings.Fields(line)
		if len(fields) < 14 {
			continue
		}
		name := fields[2]
		if strings.HasPrefix(name, "loop") || strings.HasPrefix(name, "ram") {
			continue
		}
		fmt.Fprintf(&b, "%s\n", strings.TrimSpace(line))
	}
	return b.String()
}

// sampleHostPressure reads the host's PSI counters. Available on cgroup-v2
// hosts (Linux 4.20+, all current GitHub runners). Each counter has avg10/60/300
// percentages plus a monotonic "total" of stalled microseconds - the deltas
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
	for name := range strings.FieldsSeq(string(listOut)) {
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

// sampleK3dKubepodsPressure walks every cpu.pressure file under any kubepods
// hierarchy inside each k3d node and emits its contents. With CPU requests
// but no limits, throttling stats stay at zero even under heavy contention;
// per-cgroup PSI is the only direct quantification of "this container's
// tasks were stuck waiting for CPU." The "full avg10/avg60/avg300" fields
// read as percentages of wall time during which every task in the cgroup
// was waiting on CPU. Stable above ~5% indicates real starvation, numbers
// approaching 100% indicate a fully blocked cgroup.
//
// Path-probing matters: cgroup-v2 layouts vary by distro and runtime. The
// previous revision hardcoded /sys/fs/cgroup/kubepods.slice and silently
// returned nothing on this k3s setup (find produced 0 results), so we had
// no per-pod pressure data despite thinking we did. The script below probes
// known kubepods locations and prints a marker line so the bundle records
// which path was actually used - if that line says "no kubepods cgroup
// found", we know to look elsewhere instead of assuming the cgroups are
// just quiet.
//
// Memory and IO pressure files are skipped to keep the sample size bounded -
// host-pressure already captures system-wide memory/IO PSI, and the failure
// mode under investigation is CPU. Add them back if the suspect changes.
func sampleK3dKubepodsPressure() string {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	listOut, err := exec.CommandContext(ctx, "docker", "ps", "--filter", "name=k3d-",
		"--filter", "status=running", "--format", "{{.Names}}").Output()
	if err != nil {
		return fmt.Sprintf("docker ps failed: %v", err)
	}
	// Try a few known kubepods locations in order. Whichever matches first
	// wins; the marker line lets us see from the bundle which path was used.
	// Filter out cgroups whose "full" line is exactly zero - they have
	// nothing interesting to say and dominate the file size on a quiet test.
	// Top-level kubepods slice is always included so we have a baseline.
	const script = `
roots="/sys/fs/cgroup/kubepods.slice /sys/fs/cgroup/kubepods /sys/fs/cgroup/kubepods-besteffort.slice /sys/fs/cgroup/kubepods-burstable.slice"
chosen=""
for r in $roots; do
  [ -d "$r" ] || continue
  count=$(find "$r" -name 'cpu.pressure' 2>/dev/null | wc -l)
  [ "$count" -gt 0 ] || continue
  chosen="$r"
  echo "### kubepods root: $r ($count cpu.pressure files)"
  break
done
if [ -z "$chosen" ]; then
  # Last-ditch search across the whole cgroup tree for any kubepods*
  # directory containing cpu.pressure. Slower but defensive.
  matched=$(find /sys/fs/cgroup -maxdepth 4 -path '*kubepods*' -name 'cpu.pressure' 2>/dev/null | head -1)
  if [ -n "$matched" ]; then
    chosen=$(dirname "$matched" | sed 's,/[^/]*$,,')
    echo "### kubepods root: $chosen (via wildcard search)"
  else
    echo "### kubepods root: NONE FOUND (probed: $roots)"
  fi
fi
[ -n "$chosen" ] || exit 0
find "$chosen" -name 'cpu.pressure' 2>/dev/null | while read f; do
  contents=$(cat "$f" 2>/dev/null)
  case "$f" in
    "$chosen/cpu.pressure") include=1 ;;
    *) include=0 ;;
  esac
  case "$contents" in
    *"full avg10=0.00"*"full avg60=0.00"*"full avg300=0.00"*"some avg10=0.00"*) ;;
    *) include=1 ;;
  esac
  [ "$include" = "1" ] || continue
  echo "=== $f ==="
  echo "$contents"
done
`
	var b strings.Builder
	for name := range strings.FieldsSeq(string(listOut)) {
		execCtx, execCancel := context.WithTimeout(context.Background(), 4*time.Second)
		out, err := exec.CommandContext(execCtx, "docker", "exec", name, "sh", "-c", script).CombinedOutput()
		execCancel()
		fmt.Fprintf(&b, "--- %s ---\n", name)
		if err != nil {
			// Timeout here is the signal — the docker exec itself didn't
			// return within 4s, suggesting either the container or docker
			// daemon is overloaded.
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

// sampleTopProcesses captures the top 30 processes by CPU on the host.
// pcpu is the % of one CPU averaged over the process's lifetime - imperfect
// for spike-detection but good enough at 2s sample cadence to spot a
// process that recently burst (its lifetime average jumps when a long-lived
// process starts a heavy phase). The cgroup column attributes each row to
// the docker container or systemd slice that owns it; cross-referencing
// against docker-stats.tsv pinpoints the responsible container, and against
// the comm/args column pinpoints the binary inside.
//
// We deliberately do not use `top -b -n1` because it spends ~1s computing
// instantaneous CPU% and would hold the goroutine for that long; ps is
// near-instant and writes the same data we need.
func sampleTopProcesses() string {
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()
	out, err := exec.CommandContext(ctx, "ps",
		"-eo", "pid,ppid,pcpu,pmem,rss,etime,stat,cgroup,comm,args",
		"--sort=-pcpu", "--no-headers",
	).Output()
	if err != nil {
		return fmt.Sprintf("ps failed: %v", err)
	}
	const keepN = 30
	lines := strings.Split(string(out), "\n")
	if len(lines) > keepN {
		lines = lines[:keepN]
	}
	return strings.Join(lines, "\n")
}

// sampleSchedstat reads /proc/schedstat. The "cpu<N>" lines record per-CPU
// scheduler counters; field 7 (1-indexed) is sum_sched_wait, the cumulative
// nanoseconds tasks spent on the runqueue waiting to be scheduled. Deltas
// between samples convert to "ns of runqueue wait incurred during this
// 2-second tick" - the cleanest single signal for "the host couldn't keep
// up with the work submitted to it". Saturated host: numbers grow by
// hundreds-of-millions of ns/sec per CPU. Quiet host: near zero.
//
// We emit the raw counters; downstream consumers compute rates. Fields
// are documented in the kernel's Documentation/scheduler/sched-stats.rst.
// Format version is on the first line; if it changes we'll see the older
// data anyway.
func sampleSchedstat() string {
	data, err := os.ReadFile("/proc/schedstat")
	if err != nil {
		return fmt.Sprintf("schedstat failed: %v", err)
	}
	return string(data)
}

// sampleK3dCgroupCPUStat reads cpu.stat from each k3d node's kubepods
// hierarchy via docker exec. cpu.stat fields of interest:
//
//	usage_usec      - cumulative CPU time used (cgroup-v2)
//	user_usec       - user-mode portion of usage_usec
//	system_usec     - kernel-mode portion of usage_usec
//	nr_periods      - number of CFS bandwidth periods elapsed
//	nr_throttled    - number of periods this cgroup was throttled
//	throttled_usec  - cumulative time throttled
//
// With CPU requests but no limits the throttle counters should stay at 0.
// PSI captures the symptom (tasks waited); cpu.stat distinguishes the
// cause (CFS bandwidth control vs shares contention). If throttled_usec
// is non-zero the diagnosis "shares contention only" is wrong.
//
// Same path-probe pattern as sampleK3dKubepodsPressure; record the path
// chosen so the bundle is self-explanatory.
func sampleK3dCgroupCPUStat() string {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	listOut, err := exec.CommandContext(ctx, "docker", "ps", "--filter", "name=k3d-",
		"--filter", "status=running", "--format", "{{.Names}}").Output()
	if err != nil {
		return fmt.Sprintf("docker ps failed: %v", err)
	}
	const script = `
roots="/sys/fs/cgroup/kubepods.slice /sys/fs/cgroup/kubepods /sys/fs/cgroup/kubepods-besteffort.slice /sys/fs/cgroup/kubepods-burstable.slice"
echo "=== /sys/fs/cgroup/cpu.stat ==="
cat /sys/fs/cgroup/cpu.stat 2>/dev/null
chosen=""
for r in $roots; do
  [ -d "$r" ] || continue
  [ -f "$r/cpu.stat" ] || continue
  chosen="$r"
  break
done
if [ -z "$chosen" ]; then
  echo "### no kubepods cpu.stat found"
  exit 0
fi
echo "### kubepods root: $chosen"
echo "=== $chosen/cpu.stat ==="
cat "$chosen/cpu.stat" 2>/dev/null
# Per-pod cpu.stat. Limit to the first 30 to keep volume bounded; more
# than that and we're looking at the wrong metric anyway.
find "$chosen" -mindepth 1 -maxdepth 3 -name 'cpu.stat' 2>/dev/null | head -30 | while read f; do
  echo "=== $f ==="
  cat "$f" 2>/dev/null
done
`
	var b strings.Builder
	for name := range strings.FieldsSeq(string(listOut)) {
		execCtx, execCancel := context.WithTimeout(context.Background(), 4*time.Second)
		out, err := exec.CommandContext(execCtx, "docker", "exec", name, "sh", "-c", script).CombinedOutput()
		execCancel()
		fmt.Fprintf(&b, "--- %s ---\n", name)
		if err != nil {
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

// sampleProcStack captures /proc/<pid>/wchan and /proc/<pid>/stack for the
// top-N CPU procs and any D-state proc on the host. wchan is a one-line
// label naming the kernel function the task is sleeping in (e.g.
// "futex_wait_queue", "io_schedule", "do_wait", "ep_poll"); stack is the
// full kernel call stack. Together they answer "what is this process
// stuck on?" without needing perf or bpftrace.
//
// Bounded: top 5 by CPU plus all D-state procs, capped at 15 total. Both
// /proc/<pid>/stack and wchan require root (we run as root in CI). On a
// developer machine without root the stack reads will produce "0xffff..."
// or empty - that's fine, just less useful.
func sampleProcStack() string {
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()
	out, err := exec.CommandContext(ctx, "ps",
		"-eo", "pid,pcpu,stat,comm",
		"--sort=-pcpu", "--no-headers",
	).Output()
	if err != nil {
		return fmt.Sprintf("ps failed: %v", err)
	}
	type proc struct{ pid, pcpu, stat, comm string }
	var top, dState []proc
	for line := range strings.SplitSeq(string(out), "\n") {
		fields := strings.Fields(line)
		if len(fields) < 4 {
			continue
		}
		p := proc{pid: fields[0], pcpu: fields[1], stat: fields[2], comm: fields[3]}
		if len(top) < 5 {
			top = append(top, p)
		}
		if strings.HasPrefix(p.stat, "D") && len(dState) < 10 {
			dState = append(dState, p)
		}
	}
	picks := append([]proc{}, top...)
	for _, d := range dState {
		dup := false
		for _, t := range top {
			if t.pid == d.pid {
				dup = true
				break
			}
		}
		if !dup {
			picks = append(picks, d)
		}
	}
	if len(picks) > 15 {
		picks = picks[:15]
	}
	var b strings.Builder
	for _, p := range picks {
		fmt.Fprintf(&b, "=== pid=%s pcpu=%s stat=%s comm=%s ===\n", p.pid, p.pcpu, p.stat, p.comm)
		if data, err := os.ReadFile("/proc/" + p.pid + "/wchan"); err == nil {
			fmt.Fprintf(&b, "wchan: %s\n", strings.TrimSpace(string(data)))
		}
		if data, err := os.ReadFile("/proc/" + p.pid + "/stack"); err == nil {
			b.WriteString("stack:\n")
			b.Write(data)
			if !strings.HasSuffix(string(data), "\n") {
				b.WriteString("\n")
			}
		}
	}
	return b.String()
}

// sampleNetCounters reads /proc/net/dev for per-interface bytes/packets/
// drops/errors and a few netfilter counters relevant to the calico-vs-flannel
// gap. Calico creates one veth pair per pod plus iptables-driven NAT; if
// drops or errors spike on those interfaces during a freeze, the freeze
// is at least partly network-bound. conntrack_count vs conntrack_max
// exhaustion is a known calico failure mode.
//
// /proc/net/dev format (tab-separated after the iface name):
//
//	iface: rx_bytes rx_packets rx_errs rx_drop ... tx_bytes tx_packets tx_errs tx_drop ...
//
// We keep all interfaces; on a CI runner there are usually fewer than 30,
// and filtering would risk excluding the calico veth pair we want.
func sampleNetCounters() string {
	var b strings.Builder
	if data, err := os.ReadFile("/proc/net/dev"); err == nil {
		b.WriteString("--- /proc/net/dev ---\n")
		b.Write(data)
	}
	for _, p := range []string{
		"/proc/sys/net/netfilter/nf_conntrack_count",
		"/proc/sys/net/netfilter/nf_conntrack_max",
		"/proc/sys/net/nf_conntrack_max",
	} {
		if data, err := os.ReadFile(p); err == nil {
			fmt.Fprintf(&b, "--- %s ---\n%s", p, string(data))
		}
	}
	if data, err := os.ReadFile("/proc/net/sockstat"); err == nil {
		fmt.Fprintf(&b, "--- /proc/net/sockstat ---\n%s", string(data))
	}
	return b.String()
}

// sampleCgroupIOStat walks the cgroup-v2 hierarchy one level deep and
// emits the io.stat file for each top-level slice plus every direct
// docker/containerd scope under system.slice. Format per cgroup-v2 docs:
//
//	<major>:<minor> rbytes=N wbytes=N rios=N wios=N dbytes=N dios=N
//
// One line per device per cgroup. Counters are cumulative; deltas
// between samples convert to bytes/sec per cgroup. This is the missing
// piece for "the disk is busy but docker stats says only 25 MB total" -
// kernel-side writes (jbd2, kworker writeback) charge to root io.stat,
// dockerd/containerd to system.slice, kine inside the k3d container to
// /system.slice/docker-<id>.scope.
//
// Bounded enumeration: walk depth 2 from /sys/fs/cgroup, skipping
// nested kubepods because they're already covered by the per-pod
// kubepods cpu.pressure sampler. Final file is large but bounded -
// most cgroups have one line per active device.
func sampleCgroupIOStat() string {
	var b strings.Builder
	const root = "/sys/fs/cgroup"
	emit := func(path string) {
		ioPath := filepath.Join(path, "io.stat")
		data, err := os.ReadFile(ioPath)
		if err != nil || len(data) == 0 {
			return
		}
		fmt.Fprintf(&b, "=== %s ===\n%s", ioPath, string(data))
		if !strings.HasSuffix(string(data), "\n") {
			b.WriteString("\n")
		}
	}
	// Root + first-level slices.
	emit(root)
	if entries, err := os.ReadDir(root); err == nil {
		for _, e := range entries {
			if !e.IsDir() {
				continue
			}
			emit(filepath.Join(root, e.Name()))
			// Drill one more level into system.slice and docker
			// containers, which is where the real attribution lives.
			if e.Name() == "system.slice" || e.Name() == "user.slice" {
				sub := filepath.Join(root, e.Name())
				if subs, err := os.ReadDir(sub); err == nil {
					for _, se := range subs {
						if !se.IsDir() {
							continue
						}
						// Skip kubepods - covered by k3d-kubepods-pressure.
						if strings.Contains(se.Name(), "kubepods") {
							continue
						}
						emit(filepath.Join(sub, se.Name()))
					}
				}
			}
		}
	}
	return b.String()
}

// sampleProcIO reads /proc/<pid>/io for every process and keeps the
// top-30 by write_bytes (cumulative). Fields per kernel docs:
//
//	rchar / wchar              - bytes the process tried to read/write
//	syscr / syscw              - read/write syscalls issued
//	read_bytes / write_bytes   - bytes that hit the storage layer
//	cancelled_write_bytes      - bytes from truncated/unlinked dirty pages
//
// write_bytes is the most useful for "who is fsync-storming the disk".
// rchar - read_bytes >> 0 means the process is page-cache-friendly;
// roughly equal means it's hitting the disk on every read. Cumulative
// counters; downstream consumer computes deltas.
//
// /proc/<pid>/io requires either being the process owner or having
// CAP_SYS_PTRACE. CI runs as root so we get everything; on a developer
// machine without root we'll silently get fewer rows.
func sampleProcIO() string {
	entries, err := os.ReadDir("/proc")
	if err != nil {
		return fmt.Sprintf("read /proc failed: %v", err)
	}
	type rec struct {
		pid, comm, body string
		writeBytes      int64
	}
	var rows []rec
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		pid := e.Name()
		if pid == "" || pid[0] < '0' || pid[0] > '9' {
			continue
		}
		data, err := os.ReadFile("/proc/" + pid + "/io")
		if err != nil || len(data) == 0 {
			continue
		}
		var wb int64
		for line := range strings.SplitSeq(string(data), "\n") {
			if v, ok := strings.CutPrefix(line, "write_bytes: "); ok {
				_, _ = fmt.Sscanf(v, "%d", &wb)
				break
			}
		}
		comm := ""
		if c, err := os.ReadFile("/proc/" + pid + "/comm"); err == nil {
			comm = strings.TrimSpace(string(c))
		}
		rows = append(rows, rec{pid: pid, comm: comm, body: string(data), writeBytes: wb})
	}
	// Sort descending by write_bytes. Avoid pulling in sort.Slice -
	// keep the function dependency-light; partial sort for top-30 is
	// fine since we only need the head.
	const keepN = 30
	for i := 0; i < len(rows) && i < keepN; i++ {
		maxIdx := i
		for j := i + 1; j < len(rows); j++ {
			if rows[j].writeBytes > rows[maxIdx].writeBytes {
				maxIdx = j
			}
		}
		rows[i], rows[maxIdx] = rows[maxIdx], rows[i]
	}
	if len(rows) > keepN {
		rows = rows[:keepN]
	}
	var b strings.Builder
	for _, r := range rows {
		fmt.Fprintf(&b, "=== pid=%s comm=%s ===\n%s", r.pid, r.comm, r.body)
		if !strings.HasSuffix(r.body, "\n") {
			b.WriteString("\n")
		}
	}
	return b.String()
}
