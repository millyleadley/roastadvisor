package domain

import (
	"time"

	"gopkg.in/guregu/null.v3"
)

type Restaurant struct {
	ID        string  `db:"id"`
	Name      string  `db:"name"`
	Address   string  `db:"address"`
	Latitude  float32 `db:"latitude"`
	Longitude float32 `db:"longitude"`

	// The unique reference on google maps
	// PlaceID string `db:"place_id"`
	// DataID string `db:"data_id"`

	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
	ArchivedAt null.Time `db:"archived_at" validate:"optional"`
}
