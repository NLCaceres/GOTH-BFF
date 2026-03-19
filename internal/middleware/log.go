package middleware

import (
	"github.com/NLCaceres/goth-example/internal/util/log"
	"log/slog"
	"time"
)

func AppLogger() *slog.Logger {
	return log.JSONLogger(defaultOpts())
}

func defaultOpts() *slog.HandlerOptions {
	return &slog.HandlerOptions{ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
		if a.Key == slog.TimeKey {
			t := a.Value.Time()
			a.Value = slog.StringValue(t.Format(time.RFC822Z))
		}
		return a
	}}
}
