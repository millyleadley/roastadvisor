package router

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/millyleadley/roastadvisor/app/api"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// Need this to find the swagger docs
	_ "github.com/millyleadley/roastadvisor/docs"
)

func Start(ctx context.Context, db *sqlx.DB) {
	router := gin.Default()

	// TODO: Apply middlewares
	// https://github.com/gin-gonic/contrib?tab=readme-ov-file
	// router.Use(middlewares.GinLogger())

	url := ginSwagger.URL("http://localhost:8000/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.String(200, "ok")
	})

	router.GET("/reviews", api.GetReviews(db))

	// Listen for requests
	router.Run("localhost:8000")
}
