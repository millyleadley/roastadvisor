package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/millyleadley/roastadvisor/lib/log"
	"github.com/samber/lo"
)

var (
	NotLogUrls = []string{"/favicon.ico"}
)

// Not used at the moment because the gin logs are pretty good, but something like this could be useful in the future.
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		c.Next()

		if !lo.Contains(NotLogUrls, path) {
			cost := time.Since(start)
			log.Info("📡 Received request", map[string]any{
				"status":     c.Writer.Status(),
				"method":     c.Request.Method,
				"path":       path,
				"query":      query,
				"ip":         c.ClientIP(),
				"user-agent": c.Request.UserAgent(),
				"errors":     c.Errors.ByType(gin.ErrorTypePrivate).String(),
				"cost":       cost,
			})
		}
	}
}
