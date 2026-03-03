package main

import (
	"context"
	"github.com/NLCaceres/goth-example/internal/route"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"log/slog"
	"os"
	"os/signal"
	"slices"
	"strings"
	"syscall"
	"time"
)

func main() {
	logOpts := &slog.HandlerOptions{ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
		if a.Key == slog.TimeKey {
			t := a.Value.Time()
			a.Value = slog.StringValue(t.Format(time.RFC822Z))
		}
		return a
	}}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, logOpts))
	if dotEnvErr := godotenv.Load(); dotEnvErr != nil {
		logger.Error("Environment not properly loaded")
	} // NOTE: "log" AND "fmt" print to the terminal BUT Echo's logger easily hides them

	app := echo.New()
	app.Logger = logger
	// `Use` must be used & declared BEFORE starting the app
	app.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus: true, LogURI: true,
		HandleError: true, // Forward errors to the global handler to decide status code
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
	}))

	app.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root: "static",
		Skipper: func(c *echo.Context) bool { // Skip if returning true
			isValidRef := strings.HasPrefix(c.Request().Referer(), "http://localhost:3000") ||
				strings.HasPrefix(c.Request().Referer(), "http://127.0.0.1:7331") ||
				strings.HasPrefix(c.Request().Referer(), "https://localhost:3000")
			isNavigation := slices.Contains(c.Request().Header["Sec-Fetch-Mode"], "navigate")
			isSameOrigin := slices.Contains(c.Request().Header["Sec-Fetch-Site"], "same-origin")
			if isValidRef && !isNavigation && isSameOrigin {
				return false
			}
			return true
		},
	}))

	route.Routes(app) // Routes must ALSO be declared before `app.Start` is called

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel() // Run on shutdown signal

	serverConfig := echo.StartConfig{Address: "localhost:3000", GracefulTimeout: time.Second * 5}
	if err := serverConfig.Start(ctx, app); err != nil {
		app.Logger.Error("Could not start the server", "ErrMsg", err)
	}
}
