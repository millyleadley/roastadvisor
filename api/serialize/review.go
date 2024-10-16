package serialize

import (
	pkgreviews "github.com/millyleadley/roastadvisor/api/gen/reviews"
	"github.com/millyleadley/roastadvisor/pkg/domain"
)

func Review(review domain.Review) *pkgreviews.Review {
	return &pkgreviews.Review{
		ID:           review.ID,
		UserID:       review.UserID,
		RestaurantID: review.RestaurantID,
	}
}
