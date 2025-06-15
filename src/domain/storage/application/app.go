package application

import (
	"kredit-plus/src/domain/storage/service"
	"kredit-plus/src/kernel"
	"kredit-plus/src/middleware"

	"github.com/labstack/echo/v4"
)

func AddRoutes(g *echo.Echo, k *kernel.Kernel) {
	svc := service.NewService(k.GetDB(), k.GetGCS())
	mdw := middleware.NewEnsureToken(k.GetCache(), k.GetDB())

	routes := g.Group("/storage/file")

	routes.GET("/:path", getFileApp(svc))
	routes.POST("", createFileUploadApp(svc), mdw.ValidateToken)
}
