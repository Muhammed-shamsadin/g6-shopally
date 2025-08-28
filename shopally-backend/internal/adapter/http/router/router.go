package router

import (
		"net/http"

		"github.com/gin-gonic/gin"
		"github.com/shopally-ai/internal/adapter/handler"
		"github.com/shopally-ai/internal/config"
	)

	// SetupRouter builds the Gin engine and registers application routes. It uses
	// the package-level handler function `handler.SearchFunc()` which relies on
	// the application to inject the usecase at startup via
	// `handler.InjectSearchUseCase(...)`.
	func SetupRouter(cfg *config.Config) *gin.Engine {
		router := gin.Default()

		version1 := router.Group("/api/v1")

		version1.GET("/health", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"status": "OK",
			})
		})

		// Public routes
		version1.GET("/search", handler.SearchFunc())

		return router
	}
