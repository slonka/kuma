package report

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/ginkgo/v2/types"

	"github.com/kumahq/kuma/v2/pkg/util/files"
)

var (
	BaseDir       = "results"
	DumpOnSuccess = false
	// PostDumpHook runs at the end of DumpReport, regardless of whether any
	// spec failed. Used by the host-sampler to write a comparison baseline
	// even on fully-green suites. Default nil = no-op. Kept here (rather
	// than in the framework package) to avoid an import cycle: framework
	// already imports this package, so it can populate the hook from init().
	PostDumpHook func(report ginkgo.Report)
)

func stagingDir() string {
	return path.Join(BaseDir, "..", "kuma-test-staging")
}

// AddFileToReportEntry adds a file to the report.
//
// When called from a spec context (BeforeEach/It/AfterEach/etc.), the file is
// written directly to its final per-spec location under BaseDir so the bundle
// is usable even if the test process is killed before DumpReport runs - which
// is the common case on a job-level timeout-cancellation in CI. Suite-level
// callers (BeforeSuite/AfterSuite) fall back to staging because their spec
// context isn't known yet; DumpReport later moves them into place.
//
// The Ginkgo report entry is still registered so that DumpReport can write
// `combined.log`/`report.txt` per spec; for already-eager files the writeEntry
// path becomes a no-op rename (src == dst).
func AddFileToReportEntry(name string, content any) {
	outPath, err := pickReportPath(name)
	if err != nil {
		logf("[WARNING]: %v", err)
		return
	}
	if err := writeReportContent(outPath, content); err != nil {
		logf("[WARNING]: %v", err)
		return
	}
	ginkgo.AddReportEntry(name, outPath, ginkgo.ReportEntryVisibilityNever)
}

func pickReportPath(name string) (string, error) {
	specReport := ginkgo.CurrentSpecReport()
	if specReport.FullText() != "" {
		// Eager path: write directly to the final per-spec dir. This
		// matches the layout DumpReport produces, so the writeEntry no-op
		// rename later is safe and idempotent.
		specDir := filepath.Join(BaseDir, files.ToValidUnixFilename(specReport.FullText()))
		if err := os.MkdirAll(specDir, 0o755); err != nil {
			return "", fmt.Errorf("could not create spec dir %s: %w", specDir, err)
		}
		return filepath.Join(specDir, files.ToValidUnixFilename(name)), nil
	}
	// Suite-level call: spec is not yet known. Stage and let DumpReport
	// move it once it can attribute the entry to its spec.
	base := stagingDir()
	if err := os.MkdirAll(base, 0o755); err != nil {
		return "", fmt.Errorf("could not create staging dir %s: %w", base, err)
	}
	tmp, err := os.CreateTemp(base, "report-*")
	if err != nil {
		return "", fmt.Errorf("could not create temp file: %w", err)
	}
	if err := tmp.Close(); err != nil {
		return "", fmt.Errorf("could not close temp file %s: %w", tmp.Name(), err)
	}
	return tmp.Name(), nil
}

func writeReportContent(outPath string, content any) error {
	switch c := content.(type) {
	case string:
		if files.FileExists(c) {
			// Caller passed an existing file path. Copy into the bundle
			// so the bundle is self-contained even if the source file is
			// later cleaned up (e.g., session temp file on test exit).
			return files.CopyFile(c, outPath)
		}
		return os.WriteFile(outPath, []byte(c), 0o600)
	case []byte:
		return os.WriteFile(outPath, c, 0o600)
	default:
		return os.WriteFile(outPath, fmt.Appendf(nil, "%v", content), 0o600)
	}
}

// DumpReport dumps the report to the disk.
//
// With AddFileToReportEntry now writing eagerly into BaseDir/<spec>/, the
// per-entry writeEntry calls below become idempotent no-op renames for
// already-on-disk files. DumpReport's remaining job is to (a) write the
// per-spec metadata files (combined.log, report.txt) which depend on the
// spec's final state, and (b) flush any staged entries from suite-level
// callers into their final location once the failed spec is known.
//
// The previous "rename BaseDir to tmpDir to clear stale state" preamble has
// been dropped because it would discard the eagerly-written per-spec files
// that already accumulated during the suite. CI starts from a fresh
// checkout; local users should `rm -rf build/reports/e2e-debug` themselves.
func DumpReport(report ginkgo.Report) {
	ginkgo.GinkgoHelper()
	basePath := BaseDir
	logf("saving report to %q DumpOnSuccess: %v", basePath, DumpOnSuccess)
	writeEntry := func(path string, data string) {
		err := os.MkdirAll(filepath.Dir(path), 0o755)
		if err != nil {
			logf("[WARNING]: failed to create directory %q: %v", path, err)
			return
		}
		// If the value is a file that actually exists let's simply move it in
		if files.FileExists(data) {
			if strings.Contains(os.TempDir(), data) {
				// If the file is in the temp dir, let's copy it
				err = files.CopyFile(data, path)
			} else {
				err = os.Rename(data, path)
			}
		} else {
			err = os.WriteFile(path, []byte(data), 0o600)
		}
		if err != nil {
			logf("[WARNING]: failed to write file %q: %v", path, err)
		}
	}
	for _, entry := range report.SpecReports {
		if entry.State == types.SpecStatePending || entry.State == types.SpecStateSkipped {
			continue
		}
		if DumpOnSuccess || entry.Failed() {
			entryPath := path.Join(basePath, files.ToValidUnixFilename(entry.FullText()))
			writeEntry(path.Join(entryPath, "combined.log"), entry.CombinedOutput())
			f := &strings.Builder{}
			_, _ = fmt.Fprintf(f, "Entry[%s]: %s\n", entry.LeafNodeType, entry.FullText())
			_, _ = fmt.Fprintf(f, "State: %s\n", entry.State)
			_, _ = fmt.Fprintf(f, "Duration: %s\n", entry.RunTime)
			_, _ = fmt.Fprintf(f, "Start: %s End: %s\n", entry.StartTime, entry.EndTime)
			if entry.Failed() {
				_, _ = fmt.Fprintf(f, "Failure:%s\n", entry.Failure.FailureNodeLocation.String())
				_, _ = fmt.Fprintf(f, "Location:%s\n", entry.Failure.Location.String())
				_, _ = fmt.Fprintf(f, "---Failure--\n%v\n", entry.Failure.Message)
				_, _ = fmt.Fprintf(f, "---StackTrace---\n%s\n", entry.Failure.Location.FullStackTrace)
			}
			_, _ = fmt.Fprintf(f, "SpecEvents:\n")
			for _, e := range entry.SpecEvents {
				_, _ = fmt.Fprintf(f, "\t%s\n", e.GomegaString())
			}
			writeEntry(path.Join(entryPath, "report.txt"), f.String())

			for _, e := range entry.ReportEntries {
				writeEntry(path.Join(entryPath, files.ToValidUnixFilename(e.Name)), e.StringRepresentation())
			}
		}
	}
	if err := os.RemoveAll(stagingDir()); err != nil {
		logf("[WARNING]: failed to remove staging directory %q: %v", stagingDir(), err)
	}
	if PostDumpHook != nil {
		PostDumpHook(report)
	}
	logf("saved report to %q", basePath)
}

func logf(c string, args ...any) {
	ginkgo.GinkgoWriter.Printf(c, args...)
	ginkgo.GinkgoWriter.Println("")
}
