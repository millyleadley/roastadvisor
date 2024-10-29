package domain

import (
	"time"

	"gopkg.in/guregu/null.v3"
)

type Review struct {
	ID      string `db:"id"`
	UserID  string `db:"user_id"`
	Comment string `db:"comment"`

	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
	ArchivedAt null.Time `db:"archived_at"`
}
