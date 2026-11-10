package logger

import (
	"context"
	"io"
	"log/slog"
	"os"
	"strings"
	"time"
)

type contextKey string

var (
	TraceIDKey  contextKey = "trace_id"
	saoPauloLoc *time.Location
)

type ContextHandler struct {
	slog.Handler
}

func init() {
	loc, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		loc = time.FixedZone("BRT", -3*3600)
	}
	saoPauloLoc = loc
}

func SetupDefault(format string, level slog.Level) *slog.Logger {
	l := NewLogger(os.Stdout, format, level)
	slog.SetDefault(l)

	return l
}

func NewLogger(w io.Writer, format string, level slog.Level) *slog.Logger {
	if w == nil {
		w = os.Stdout
	}

	opts := &slog.HandlerOptions{
		AddSource: true,
		Level:     level,
		ReplaceAttr: func(_ []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				t := a.Value.Time().In(saoPauloLoc)
				return slog.String(slog.TimeKey, t.Format("2006-01-02T15:04:05.000-0700"))
			}
			return a
		},
	}
	handler := slog.NewJSONHandler(os.Stdout, opts)

	customLogger := slog.New(handler)
	slog.SetDefault(customLogger)

	var baseHandler slog.Handler
	if strings.EqualFold(format, "text") {
		baseHandler = slog.NewTextHandler(w, opts)
	} else {
		baseHandler = slog.NewJSONHandler(w, opts)
	}

	return slog.New(&ContextHandler{baseHandler})
}

func WithTraceID(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, TraceIDKey, traceID)
}
