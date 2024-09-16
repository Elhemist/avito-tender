package models

import (
	"time"

	"github.com/google/uuid"
)

type Employee struct {
	Id        uuid.UUID `db:"id"`
	Username  string    `db:"name"`
	FirstName string    `db:"first_name, omitempty"`
	LastName  string    `db:"last_name, omitempty"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
