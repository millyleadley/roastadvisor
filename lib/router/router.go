package router

import (
	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.String(200, "ok")
	})

	// Listens on localhost:8080 by default
	router.Run()
}
