package middleware

import (
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"slices"
	"strings"
)

func Static() echo.MiddlewareFunc {
	return middleware.StaticWithConfig(middleware.StaticConfig{
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
	})
}
