package middleware

import (
	"github.com/NLCaceres/goth-example/internal/util/stringy"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"net/http"
)

func Static() echo.MiddlewareFunc {
	return middleware.StaticWithConfig(middleware.StaticConfig{
		Root: "static",
		Skipper: func(c *echo.Context) bool { // Skip if returning true
			h := c.Request().Header
			return !(isFetchSiteSameOrigin(h) && !isNavigation(h) && isValidReferer(h))
		},
	})
}

// Checks that the request originated from the same origin (same scheme, host and port)
func isFetchSiteSameOrigin(h http.Header) bool {
	return h.Get("Sec-Fetch-Site") == "same-origin"
}

// Checks that the request originated from navigation between different HTML documents
// Usually set to "navigation" on first load (from address bar URL input) AND through
// normal anchor tags that receive whole HTML docs in response to their GET requests
func isNavigation(h http.Header) bool {
	return h.Get("Sec-Fetch-Mode") == "navigate"
}

// Checks that the referer is set to one of several origins (same scheme, host and port)
// It's set when browsing around a site, not on the first load.
// There should only ever be 1 Referer value, and it's set to the URL (including path)
// that fired the request off.
//
// Ex: "http://localhost:3000/foo" requests "http://localhost:3000/bar"
//
// then the Referer is "http://localhost:3000/foo"
func isValidReferer(h http.Header) bool {
	return stringy.HasAnyPrefix(h.Get("Referer"), "http://localhost:3000",
		"https://localhost:3000", "http://127.0.0.1:3000", "http://127.0.0.1:7331")
}
