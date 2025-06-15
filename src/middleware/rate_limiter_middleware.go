package middleware

import (
	"kredit-plus/src/constant"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RateLimiterMiddleware(e *echo.Echo) {
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(constant.DefaultMdwRateLimiter)))
}
