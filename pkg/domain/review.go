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

	CommentMeat     null.String `db:"comment_meat" validate:"optional"`
	CommentPotatoes null.String `db:"comment_potatoes" validate:"optional"`
	CommentVeg      null.String `db:"comment_veg" validate:"optional"`
	CommentAddOns   null.String `db:"comment_add_ons" validate:"optional"`
	CommentSize     null.String `db:"comment_size" validate:"optional"`
	CommentPrice    null.String `db:"comment_price" validate:"optional"`

	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
	ArchivedAt null.Time `db:"archived_at" validate:"optional"`
}
