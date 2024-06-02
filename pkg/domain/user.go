package domain

import (
	"time"

	"gopkg.in/guregu/null.v3"
)

type User struct {
	ID       string      `db:"id"`
	Username null.String `db:"username"`

	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
	ArchivedAt null.Time `db:"archived_at"`
}
