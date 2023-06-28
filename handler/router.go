package handler

import (
	"context"

	"github.com/caioeverest/fedits/adapter/http"
	"go.uber.org/fx"
)

// NewRouter creates a new router
func NewRouter(lc fx.Lifecycle, server *http.Server, providerHandler *Provider, methodHandler *Method, orquestratorHandler *Orquestrator) {
	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				// Provider
				{
					router := server.Group("/provider")
					router.POST("", providerHandler.Create)
					router.GET("/:slug", providerHandler.Get)
					router.PATCH("/:slug", providerHandler.Update)
					router.DELETE("/:slug", providerHandler.Delete)
					router.GET("/list/:method", providerHandler.List)
				}

				// Method
				{
					router := server.Group("/method")
					router.POST("", methodHandler.Create)
					router.GET("", methodHandler.List)
					router.GET("/:method", methodHandler.Get)
				}

				// Orquestrate
				{
					server.POST("/call", orquestratorHandler.Request)
				}
				return nil
			},
		},
	)
}
