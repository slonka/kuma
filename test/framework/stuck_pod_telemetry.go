package framework

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
	"sync"
	"time"

	"github.com/gruntwork-io/terratest/modules/k8s"
	"github.com/gruntwork-io/terratest/modules/retry"
	terratest_testing "github.com/gruntwork-io/terratest/modules/testing"
	v1 "k8s.io/api/core/v1"

	"github.com/kumahq/kuma/v2/test/framework/report"
)

const (
	stuckPodTelemetryTimeout      = 8 * time.Second
	stuckPodTelemetryTriggerAfter = 25 * time.Second
)

var (
	stuckPodTelemetryMu   sync.Mutex
	stuckPodTelemetrySeen = map[string]struct{}{}
)

func WaitUntilPodAvailableWithTelemetryE(
	t terratest_testing.TestingT,
	options *k8s.KubectlOptions,
	podName string,
	retries int,
	sleepBetweenRetries time.Duration,
) error {
	statusMsg := fmt.Sprintf("Wait for pod %s to be provisioned.", podName)
	message, err := retry.DoWithRetryE(
		t,
		statusMsg,
		retries,
		sleepBetweenRetries,
		func() (string, error) {
			pod, err := k8s.GetPodE(t, options, podName)
			if err != nil {
				return "", err
			}
			captureStuckPodTelemetry(t, options, pod)
			if !k8s.IsPodAvailable(pod) {
				return "", k8s.NewPodNotAvailableError(pod)
			}
			return "Pod is now available", nil
		},
	)
	if err != nil {
		options.Logger.Logf(t, "Timedout waiting for Pod to be provisioned: %s", err)
		return err
	}
	options.Logger.Logf(t, "%s", message)
	return nil
}

func captureStuckPodTelemetry(
	t terratest_testing.TestingT,
	kubectlOptions *k8s.KubectlOptions,
	pod *v1.Pod,
) []string {
	if pod == nil || !shouldCaptureStuckPodTelemetry(t, kubectlOptions, pod) {
		return nil
	}

	key := pod.Namespace + "/" + pod.Name
	stuckPodTelemetryMu.Lock()
	if _, seen := stuckPodTelemetrySeen[key]; seen {
		stuckPodTelemetryMu.Unlock()
		return nil
	}
	stuckPodTelemetrySeen[key] = struct{}{}
	stuckPodTelemetryMu.Unlock()

	HostSamplerMark("STUCK_POD " + key + " node=" + pod.Spec.NodeName)

	var written []string
	if hostSnapshot := collectHostStuckPodSnapshot(pod); hostSnapshot != nil {
		reportName := path.Join("stuck-pod-telemetry", pod.Namespace, pod.Name, "host.txt")
		report.AddFileToReportEntry(reportName, hostSnapshot)
		written = append(written, reportName)
	}
	if nodeSnapshot := collectNodeStuckPodSnapshot(pod); nodeSnapshot != nil {
		reportName := path.Join("stuck-pod-telemetry", pod.Namespace, pod.Name, "node-runtime.txt")
		report.AddFileToReportEntry(reportName, nodeSnapshot)
		written = append(written, reportName)
	}

	return written
}

func shouldCaptureStuckPodTelemetry(
	t terratest_testing.TestingT,
	kubectlOptions *k8s.KubectlOptions,
	pod *v1.Pod,
) bool {
	if pod == nil || !hasInitContainerNamed(pod, "kuma-init") {
		return false
	}
	if podInitialized(pod) {
		return false
	}
	kumaInitStatus, found := findContainerStatus(pod.Status.InitContainerStatuses, "kuma-init")
	if !found || !containerRunningLongEnough(kumaInitStatus, stuckPodTelemetryTriggerAfter) {
		return false
	}
	logs, err := k8s.GetPodLogsE(t, kubectlOptions, pod, "kuma-init")
	if err != nil {
		return false
	}
	return strings.TrimSpace(logs) == ""
}

func hasInitContainerNamed(pod *v1.Pod, name string) bool {
	for _, initContainer := range pod.Spec.InitContainers {
		if initContainer.Name == name {
			return true
		}
	}
	return false
}

func podInitialized(pod *v1.Pod) bool {
	for _, condition := range pod.Status.Conditions {
		if condition.Type == v1.PodInitialized {
			return condition.Status == v1.ConditionTrue
		}
	}
	return false
}

func findContainerStatus(statuses []v1.ContainerStatus, name string) (v1.ContainerStatus, bool) {
	for _, status := range statuses {
		if status.Name == name {
			return status, true
		}
	}
	return v1.ContainerStatus{}, false
}

func containerRunningLongEnough(status v1.ContainerStatus, minDuration time.Duration) bool {
	if status.State.Running == nil {
		return false
	}
	startedAt := status.State.Running.StartedAt.Time
	if startedAt.IsZero() {
		return false
	}
	return time.Since(startedAt) >= minDuration
}

func collectHostStuckPodSnapshot(pod *v1.Pod) []byte {
	var out bytes.Buffer
	fmt.Fprintf(&out, "captured_at=%s\n", time.Now().UTC().Format(time.RFC3339Nano))
	fmt.Fprintf(&out, "pod=%s/%s node=%s\n\n", pod.Namespace, pod.Name, pod.Spec.NodeName)

	appendHostFile(&out, "/proc/pressure/io")
	appendHostSubset(&out, "/proc/meminfo", []string{
		"Dirty:", "Writeback:", "WritebackTmp:", "DirtyThreshold:", "DirtyBackgroundThreshold:",
	})
	appendHostSubset(&out, "/proc/vmstat", []string{
		"nr_dirty ", "nr_writeback ", "nr_dirtied ", "nr_written ",
	})
	appendHostCommand(&out, 3*time.Second, "sh", "-c", "ps -eo pid,ppid,stat,wchan,cmd | awk 'NR==1 || $3 ~ /D/'")
	if pod.Spec.NodeName != "" {
		appendHostCommand(&out, 3*time.Second, "docker", "top", pod.Spec.NodeName, "-eo", "pid,ppid,stat,wchan,cmd")
	}

	return out.Bytes()
}

func appendHostFile(out *bytes.Buffer, filePath string) {
	fmt.Fprintf(out, "=== %s ===\n", filePath)
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Fprintf(out, "error: %v\n\n", err)
		return
	}
	out.Write(data)
	if len(data) == 0 || data[len(data)-1] != '\n' {
		out.WriteByte('\n')
	}
	out.WriteByte('\n')
}

func appendHostSubset(out *bytes.Buffer, filePath string, prefixes []string) {
	fmt.Fprintf(out, "=== %s subset ===\n", filePath)
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Fprintf(out, "error: %v\n\n", err)
		return
	}
	for _, line := range strings.Split(string(data), "\n") {
		for _, prefix := range prefixes {
			if strings.HasPrefix(line, prefix) {
				fmt.Fprintln(out, line)
				break
			}
		}
	}
	out.WriteByte('\n')
}

func appendHostCommand(out *bytes.Buffer, timeout time.Duration, name string, args ...string) {
	fmt.Fprintf(out, "=== %s %s ===\n", name, strings.Join(args, " "))
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	cmdOut, err := exec.CommandContext(ctx, name, args...).CombinedOutput()
	if len(cmdOut) > 0 {
		out.Write(cmdOut)
		if cmdOut[len(cmdOut)-1] != '\n' {
			out.WriteByte('\n')
		}
	}
	if err != nil {
		fmt.Fprintf(out, "error: %v\n", err)
	}
	out.WriteByte('\n')
}

func collectNodeStuckPodSnapshot(pod *v1.Pod) []byte {
	if pod.Spec.NodeName == "" {
		return nil
	}
	switch Config.K8sType {
	case K3dK8sType, K3dCalicoK8sType, KindK8sType:
	default:
		return nil
	}

	script := fmt.Sprintf(`POD_NAME=%q
echo "captured_at=%s"
echo "pod=%s/%s node=%s"
echo
echo "=== crictl pods --name $POD_NAME ==="
crictl pods -a --name "$POD_NAME" -o json 2>&1 || true
POD_ID="$(crictl pods -a --name "$POD_NAME" -q 2>/dev/null | head -n 1)"
if [ -n "$POD_ID" ]; then
  echo
  echo "=== crictl inspectp $POD_ID ==="
  crictl inspectp "$POD_ID" 2>&1 || true
  echo
  echo "=== crictl ps -a --pod $POD_ID ==="
  crictl ps -a --pod "$POD_ID" -o json 2>&1 || true
  CTR_ID="$(crictl ps -a --pod "$POD_ID" -q 2>/dev/null | head -n 1)"
  if [ -n "$CTR_ID" ]; then
    echo
    echo "=== crictl inspect $CTR_ID ==="
    crictl inspect "$CTR_ID" 2>&1 || true
  fi
fi
echo
echo "=== ctr -n k8s.io tasks ls ==="
ctr -n k8s.io tasks ls 2>&1 || true
echo
echo "=== process list matching kuma-init / transparent-proxy / iptables ==="
ps -eo pid,ppid,stat,wchan,args 2>/dev/null | grep -E 'kumactl install transparent-proxy|kuma-init|iptables|ip6tables' | grep -v grep || true
for pid in $(ps -eo pid=,args= 2>/dev/null | grep -E 'kumactl install transparent-proxy|iptables|ip6tables' | grep -v grep | awk '{print $1}' | head -n 8); do
  echo
  echo "=== /proc/$pid/wchan ==="
  cat /proc/$pid/wchan 2>/dev/null || true
  echo
  echo "=== /proc/$pid/syscall ==="
  cat /proc/$pid/syscall 2>/dev/null || true
  echo
  echo "=== /proc/$pid/status ==="
  cat /proc/$pid/status 2>/dev/null || true
  echo
  echo "=== /proc/$pid/sched ==="
  cat /proc/$pid/sched 2>/dev/null || true
  echo
  echo "=== /proc/$pid/io ==="
  cat /proc/$pid/io 2>/dev/null || true
  echo
  echo "=== /proc/$pid/stack ==="
  cat /proc/$pid/stack 2>/dev/null || true
done
echo
echo "=== /proc/pressure/io ==="
cat /proc/pressure/io 2>/dev/null || true
echo
echo "=== /proc/meminfo dirty/writeback ==="
grep -E '^(Dirty|Writeback|WritebackTmp|DirtyThreshold|DirtyBackgroundThreshold)' /proc/meminfo 2>/dev/null || true
echo
echo "=== /proc/vmstat dirty/writeback ==="
grep -E '^(nr_dirty|nr_writeback|nr_dirtied|nr_written) ' /proc/vmstat 2>/dev/null || true
echo
echo "=== xtables lock ==="
ls -l /run/xtables.lock 2>/dev/null || true
`, pod.Name, time.Now().UTC().Format(time.RFC3339Nano), pod.Namespace, pod.Name, pod.Spec.NodeName)

	ctx, cancel := context.WithTimeout(context.Background(), stuckPodTelemetryTimeout)
	defer cancel()
	out, err := exec.CommandContext(ctx, "docker", "exec", pod.Spec.NodeName, "sh", "-c", script).CombinedOutput()
	if err == nil {
		return out
	}

	var wrapped bytes.Buffer
	wrapped.Write(out)
	if len(out) > 0 && out[len(out)-1] != '\n' {
		wrapped.WriteByte('\n')
	}
	fmt.Fprintf(&wrapped, "error: %v\n", err)
	return wrapped.Bytes()
}
