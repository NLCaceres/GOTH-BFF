package middleware

import (
	"github.com/NLCaceres/goth-example/internal/util/http"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func Static() echo.MiddlewareFunc {
	return middleware.StaticWithConfig(middleware.StaticConfig{
		Root: "static",
		Skipper: func(c *echo.Context) bool { // Skip if returning true
			h := http.Header{Header: c.Request().Header}
			return !(h.IsSiteAllSame() && !h.IsNavigationFetch() && h.IsRefererAny(http.SafeOrigins...))
		},
	})
}
