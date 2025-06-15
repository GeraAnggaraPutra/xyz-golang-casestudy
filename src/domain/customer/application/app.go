package application

import (
	"kredit-plus/src/domain/customer/service"
	"kredit-plus/src/kernel"
	"kredit-plus/src/middleware"

	"github.com/labstack/echo/v4"
)

func AddRoutes(g *echo.Echo, k *kernel.Kernel) {
	svc := service.NewService(k.GetDB())
	mdw := middleware.NewEnsureToken(k.GetCache(), k.GetDB())

	routes := g.Group("/customer", mdw.ValidateToken)

	routes.GET("", readCustomerListApp(svc))
	routes.GET("/:guid", readCustomerDetailApp(svc))
	routes.POST("", createCustomerApp(svc))
	routes.PUT("/:guid", updateCustomerApp(svc))
	routes.DELETE("/:guid", deleteCustomerApp(svc))
}
