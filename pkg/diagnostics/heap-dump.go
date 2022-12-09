package diagnostics

import (
    "fmt"
    "io"
    "net/http"
    "os"
    "runtime/debug"
    "strconv"
    "time"
)

func HeapDump(w http.ResponseWriter, r *http.Request) {
    filename := "heapdump-" +strconv.FormatInt(time.Now().Unix(), 10)
    f, err := os.CreateTemp("", filename)
    if err != nil {
        serveError(w, 500, err.Error())
        return
    }
    defer func() {
        err := f.Close()
        if err != nil {
            serveError(w, 500, err.Error())
            return
        }
        err = os.Remove(f.Name())
        if err != nil {
            serveError(w, 500, err.Error())
            return
        }
    }()

    debug.FreeOSMemory()
    debug.WriteHeapDump(f.Fd())
    if _, err := f.Seek(0, 0); err != nil {
        serveError(w, 500, err.Error())
        return
    }

    w.Header().Set("Content-Type", "application/octet-stream")
    w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))
    w.WriteHeader(http.StatusOK)
    _, err = io.Copy(w, f)
    if err != nil {
        serveError(w, 500, err.Error())
    }

    return
}

func serveError(w http.ResponseWriter, status int, txt string) {
    w.Header().Set("Content-Type", "text/plain; charset=utf-8")
    w.Header().Set("X-Go-Pprof", "1")
    w.Header().Del("Content-Disposition")
    w.WriteHeader(status)
}