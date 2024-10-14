package app

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/millyleadley/roastadvisor/lib/log"
	"github.com/millyleadley/roastadvisor/pkg/domain"
)

// @Summary      Get array of reviews
// @Description  Responds with the list of all reviews as JSON.
// @Produce      json
// @Success      200  {array}  domain.Review
// @Router       /reviews [get]
func GetReviews(db *sqlx.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		reviews := []domain.Review{}
		err := db.Select(&reviews, "SELECT * FROM reviews")
		if err != nil {
			log.Error(err)
			c.JSON(500, gin.H{
				"message": "Internal server error",
			})
			return
		}

		c.JSON(200, reviews)
	}
}
