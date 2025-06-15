package application

import (
	"kredit-plus/src/domain/transaction/service"
	"kredit-plus/src/kernel"
	"kredit-plus/src/middleware"

	"github.com/labstack/echo/v4"
)

func AddRoutes(g *echo.Echo, k *kernel.Kernel) {
	svc := service.NewService(k.GetDB())
	mdw := middleware.NewEnsureToken(k.GetCache(), k.GetDB())

	routes := g.Group("/transaction", mdw.ValidateToken)

	routes.GET("", readTransactionListApp(svc))
	routes.GET("/:guid", readTransactionDetailApp(svc))
	routes.POST("", createTransactionApp(svc))
}
