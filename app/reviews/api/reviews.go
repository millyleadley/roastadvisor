package apireviews

import (
	"context"

	"github.com/jmoiron/sqlx"
	pkgreviews "github.com/millyleadley/roastadvisor/api/gen/reviews"
	"github.com/millyleadley/roastadvisor/api/serialize"
	"github.com/millyleadley/roastadvisor/lib/errors"
	"github.com/millyleadley/roastadvisor/pkg/domain"
	"github.com/samber/lo"
)

type reviewssrvc struct {
	db *sqlx.DB
}

// NewReviews returns the Reviews service implementation.
func NewReviews(db *sqlx.DB) pkgreviews.Service {
	return &reviewssrvc{
		db: db,
	}
}

// List all reviews.
func (s *reviewssrvc) List(ctx context.Context) ([]*pkgreviews.Review, error) {
	reviews := []domain.Review{}
	err := s.db.Select(&reviews, "SELECT * FROM reviews")
	if err != nil {
		return nil, errors.Wrap(ctx, err, "listing reviews")
	}

	seralized := lo.Map(reviews, func(item domain.Review, _ int) *pkgreviews.Review {
		return serialize.Review(item)
	})

	return seralized, nil
}
