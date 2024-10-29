package domain

import (
	"time"

	"gopkg.in/guregu/null.v3"
)

type User struct {
	ID string `db:"id"`
	// Email must be unique; so no two users can have the same email
	Email          string `db:"email"`
	HashedPassword string `db:"hashed_password"`
	// Sessiontoken stores cookie info for the user
	SessionToken null.String `db:"session_token"`

	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
	ArchivedAt null.Time `db:"archived_at"`
}
