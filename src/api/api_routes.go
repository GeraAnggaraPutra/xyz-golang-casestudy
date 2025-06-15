package api

import (
	"net/http"
	"os"

	"kredit-plus/src/kernel"
	"kredit-plus/src/middleware"

	"github.com/labstack/echo/v4"

	authDomain "kredit-plus/src/domain/auth/application"
	storageDomain "kredit-plus/src/domain/storage/application"
	userDomain "kredit-plus/src/domain/user/application"
	customerDomain "kredit-plus/src/domain/customer/application"
)

func routes(e *echo.Echo, k *kernel.Kernel) {
	// register middleware
	middleware.TimeoutMiddleware(e)
	middleware.RecoverMiddleware(e)
	middleware.RateLimiterMiddleware(e)
	middleware.CorsMiddleware(e)

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{
			"message": os.Getenv("APP_NAME") + "is Running",
		})
	})

	// domain routes
	authDomain.AddRoutes(e, k)
	storageDomain.AddRoutes(e, k)
	userDomain.AddRoutes(e, k)
	customerDomain.AddRoutes(e, k)
}
