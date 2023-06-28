package http

import (
	"context"
	"fmt"

	_ "github.com/caioeverest/fedits/docs"

	"github.com/caioeverest/fedits/internal/config"
	"github.com/caioeverest/fedits/internal/logger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go.uber.org/fx"
)

type Server struct {
	*echo.Echo
}

// New builds an HTTP server that will begin serving requests
// when the Fx application starts.
func New(lc fx.Lifecycle, cfg *config.Config, log *logger.Logger) *Server {
	e := echo.New()
	e.HideBanner = true

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := e.Start(fmt.Sprintf(":%d", cfg.HTTPPort)); err != nil {
					log.Errorf("shutting down the server: %v", err)
					return
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return e.Shutdown(ctx)
		},
	})
	return &Server{e}
}
