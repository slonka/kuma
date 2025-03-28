package log

import (
	"context"
	"io"
	"os"

	"github.com/go-logr/logr"
	"github.com/kumahq/kuma/pkg/multitenant"
	logger_extensions "github.com/kumahq/kuma/pkg/plugins/extensions/logger"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel/trace"
	"gopkg.in/natefinch/lumberjack.v2"
	"k8s.io/klog/v2/textlogger"
)

type LogLevel int

const (
	OffLevel LogLevel = iota
	InfoLevel
	DebugLevel
)

func (l LogLevel) String() string {
	switch l {
	case OffLevel:
		return "off"
	case InfoLevel:
		return "info"
	case DebugLevel:
		return "debug"
	default:
		return "unknown"
	}
}

func ParseLogLevel(text string) (LogLevel, error) {
	switch text {
	case "off":
		return OffLevel, nil
	case "info":
		return InfoLevel, nil
	case "debug":
		return DebugLevel, nil
	default:
		return OffLevel, errors.Errorf("unknown log level %q", text)
	}
}

func NewLogger(level LogLevel) logr.Logger {
	return NewLoggerTo(os.Stderr, level)
}

func NewLoggerWithRotation(level LogLevel, outputPath string, maxSize int, maxBackups int, maxAge int) logr.Logger {
	return NewLoggerTo(&lumberjack.Logger{
		Filename:   outputPath,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
	}, level)
}

func NewLoggerTo(destWriter io.Writer, level LogLevel) logr.Logger {
	config := textlogger.NewConfig(textlogger.Verbosity(int(level)), textlogger.Output(destWriter))
	logger := textlogger.NewLogger(config)
	return logger
}

const TenantLoggerKey = "tenantID"

// AddFieldsFromCtx will check if provided context contain tracing span and
// if the span is currently recording. If so, it will call spanLogValuesProcessor
// function if it's also present in the context. If not it will add trace_id
// and span_id to logged values. It will also add the tenant id to the logged
// values.
func AddFieldsFromCtx(
	logger logr.Logger,
	ctx context.Context,
	extensions context.Context,
) logr.Logger {
	if tenantId, ok := multitenant.TenantFromCtx(ctx); ok {
		logger = logger.WithValues(TenantLoggerKey, tenantId)
	}

	return addSpanValuesToLogger(logger, ctx, extensions)
}

func addSpanValuesToLogger(
	logger logr.Logger,
	ctx context.Context,
	extensions context.Context,
) logr.Logger {
	if span := trace.SpanFromContext(ctx); span.IsRecording() {
		if fn, ok := logger_extensions.FromSpanLogValuesProcessorContext(extensions); ok {
			return logger.WithValues(fn(span)...)
		}

		return logger.WithValues(
			"trace_id", span.SpanContext().TraceID(),
			"span_id", span.SpanContext().SpanID(),
		)
	}

	return logger
}
