package api

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/millyleadley/roastadvisor/lib/log"
)

func GetReviews(c *gin.Context) {
	log.Error(errors.New("not implemented"))

	c.JSON(401, gin.H{
		"message": "Not implemented",
	})
}
