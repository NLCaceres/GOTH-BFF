package middleware

import (
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"time"
)

func RateLimiter() echo.MiddlewareFunc {
	c := middleware.RateLimiterMemoryStoreConfig{
		Rate: 15.0, Burst: 15, ExpiresIn: time.Minute * 15,
	}
	return middleware.RateLimiter(middleware.NewRateLimiterMemoryStoreWithConfig(c))
}
