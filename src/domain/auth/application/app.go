package application

import (
	"kredit-plus/src/domain/auth/service"
	"kredit-plus/src/kernel"
	"kredit-plus/src/middleware"

	"github.com/labstack/echo/v4"
)

func AddRoutes(g *echo.Echo, k *kernel.Kernel) {
	svc := service.NewService(k.GetDB(), k.GetCache())
	mdw := middleware.NewEnsureToken(k.GetCache(), k.GetDB())

	routes := g.Group("/auth")

	routes.POST("/login", loginApp(svc))
	routes.POST("/refresh-token", refreshTokenApp(svc))
	routes.POST("/logout", logoutApp(svc), mdw.ValidateToken)
}
