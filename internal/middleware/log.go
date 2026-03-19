package middleware

import (
	"context"
	"github.com/NLCaceres/goth-example/internal/util/datetime"
	"github.com/NLCaceres/goth-example/internal/util/log"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"log/slog"
)

func AppLogger() *slog.Logger {
	return log.JSONLogger(defaultOpts())
}

func defaultOpts() *slog.HandlerOptions {
	return &slog.HandlerOptions{ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
		if a.Key == slog.TimeKey {
			t := a.Value.Time()
			a.Value = slog.StringValue(t.Format(datetime.LogTime))
		}
		return a
	}}
}

func RequestLogger(logger *slog.Logger) echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus: true, LogURI: true, HandleError: true,
		LogValuesFunc: func(c *echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error == nil {
				logger.LogAttrs(context.Background(), slog.LevelInfo, "REQUEST",
					slog.String("uri", v.URI), slog.Int("status", v.Status),
				)
			} else {
				logger.LogAttrs(context.Background(), slog.LevelError, "REQUEST_ERROR",
					slog.String("uri", v.URI), slog.Int("status", v.Status), slog.String("err", v.Error.Error()),
				)
			}
			return nil
		},
	})
}
