package log

import (
	"log/slog"
	"os"
)

func JSONLogger(logOpts *slog.HandlerOptions) *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, logOpts))
}
