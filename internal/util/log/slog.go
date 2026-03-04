package log

import (
	"log/slog"
	"os"
	"time"
)

func AppLogger() *slog.Logger {
	return JSONLogger(defaultOpts())
}

func JSONLogger(logOpts *slog.HandlerOptions) *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, logOpts))
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
