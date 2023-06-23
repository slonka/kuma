package logging

import (
    "github.com/kumahq/kuma/pkg/log"
    "go.uber.org/zap"
    "io"
    "net/http"
    "strings"
)

func ChangeLogLevel(writer http.ResponseWriter, request *http.Request) {
    defer request.Body.Close()
    body, err := io.ReadAll(request.Body)

    if err != nil {
        writer.WriteHeader(500)
        return
    }

    if strings.Contains(string(body), "error") {
        log.Lvl.SetLevel(zap.ErrorLevel)
    } else {
        log.Lvl.SetLevel(zap.DebugLevel)
    }
    writer.WriteHeader(200)
}
