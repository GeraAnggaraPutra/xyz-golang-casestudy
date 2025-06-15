package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"kredit-plus/src/kernel"
	"kredit-plus/toolkit/config"
	"kredit-plus/toolkit/logger"

	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

func RunEchoServer(ctx context.Context, k *kernel.Kernel) {
	cfg := config.NewRuntimeConfig()

	e := echo.New()
	e.HideBanner = true
	e.Validator = config.NewValidator()
	e.HTTPErrorHandler = errorHandler()

	// shutdown gracefully
	go func() {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeoutDuration)
		defer cancel()

		<-time.After(cfg.ShutdownWaitDuration)

		if err := e.Shutdown(ctx); err != nil {
			log.Printf("ERROR shutdown server : %s", err.Error())
		}
	}()

	// register prometheus metrics
	if cfg.Prometheus {
		e.Use(echoprometheus.NewMiddleware(cfg.Name))
		e.GET("/metrics", echoprometheus.NewHandler())
	}

	// register routes
	routes(e, k)

	log.Printf("serving REST HTTP server : %s", logger.ParseJSON(cfg))

	if err := e.Start(fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Printf("ERROR starting http server : %s", err.Error())
	}
}
