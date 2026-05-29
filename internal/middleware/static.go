package middleware

import (
	"github.com/NLCaceres/goth-example/internal/util/stringy"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func Static() echo.MiddlewareFunc {
	return middleware.StaticWithConfig(middleware.StaticConfig{
		Root: "static",
		Skipper: func(c *echo.Context) bool { // Skip if returning true
			return !(c.Request().Header.Get("Sec-Fetch-Site") == "same-origin" &&
				!(c.Request().Header.Get("Sec-Fetch-Mode") == "navigate") &&
				stringy.HasAnyPrefix(c.Request().Referer(), "http://localhost:3000",
					"https://localhost:3000", "http://127.0.0.1:3000", "http://127.0.0.1:7331"))
		},
	})
}
