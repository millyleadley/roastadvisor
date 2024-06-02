package domain

import (
	"time"

	"gopkg.in/guregu/null.v3"
)

type Review struct {
	ID           string `db:"id"`
	UserID       string `db:"user_id"`
	RestaurantID string `db:"restaurant_id"`

	RatingMeat     int `db:"rating_meat"`
	RatingPotatoes int `db:"rating_potatoes"`
	RatingVeg      int `db:"rating_veg"`
	RatingAddOns   int `db:"rating_add_ons"`
	RatingSize     int `db:"rating_size"`
	RatingPrice    int `db:"rating_price"`

	CommentMeat     null.String `db:"comment_meat"`
	CommentPotatoes null.String `db:"comment_potatoes"`
	CommentVeg      null.String `db:"comment_veg"`
	CommentAddOns   null.String `db:"comment_add_ons"`
	CommentSize     null.String `db:"comment_size"`
	CommentPrice    null.String `db:"comment_price"`

	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
	ArchivedAt null.Time `db:"archived_at"`
}
