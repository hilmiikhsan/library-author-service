package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Author struct {
	ID        uuid.UUID    `db:"id"`
	Name      string       `db:"name"`
	Bio       string       `db:"bio"`
	BirthDate time.Time    `db:"birth_date"`
	DeathDate sql.NullTime `db:"death_date"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
}
