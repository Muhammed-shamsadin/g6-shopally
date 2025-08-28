package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shopally-ai/internal/config"
)

func SetupRouter(cfg *config.Config) *gin.Engine {
	router := gin.Default()

	version1 := router.Group("/api/v1")

	version1.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	})

	//public

	// private
	return router
}
