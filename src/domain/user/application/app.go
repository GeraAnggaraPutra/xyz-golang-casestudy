package application

import (
	"kredit-plus/src/domain/user/service"
	"kredit-plus/src/kernel"
	"kredit-plus/src/middleware"

	"github.com/labstack/echo/v4"
)

func AddRoutes(g *echo.Echo, k *kernel.Kernel) {
	svc := service.NewService(k.GetDB())
	mdw := middleware.NewEnsureToken(k.GetCache(), k.GetDB())

	routes := g.Group("/user", mdw.ValidateToken)

	routes.GET("", readUserListApp(svc))
	routes.GET("/:guid", readUserDetailApp(svc))
	routes.POST("", createUserApp(svc))
	routes.PUT("/:guid", updateUserApp(svc))
	routes.DELETE("/:guid", deleteUserApp(svc))
}
