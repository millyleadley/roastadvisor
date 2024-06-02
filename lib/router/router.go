package router

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/millyleadley/roastadvisor/lib/router/api"
)

func Start(ctx context.Context, db *sqlx.DB) {
	router := gin.Default()

	// TODO: Apply middlewares
	// https://github.com/gin-gonic/contrib?tab=readme-ov-file
	// router.Use(middlewares.GinLogger())

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.String(200, "ok")
	})

	router.GET("/reviews", api.GetReviews(db))

	// Listen for requests
	router.Run("localhost:8080")
}
