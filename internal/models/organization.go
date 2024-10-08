package models

import (
	"time"

	"github.com/google/uuid"
)

type OrganizationType string

const (
	IE  OrganizationType = "IE"
	LLC OrganizationType = "LLC"
	JSC OrganizationType = "JSC"
)

type Organization struct {
	ID          uuid.UUID        `db:"id"`
	Name        string           `db:"name"`
	Description string           `db:"description,omitempty"`
	Type        OrganizationType `db:"type"`
	CreatedAt   time.Time        `db:"created_at"`
	UpdatedAt   time.Time        `db:"updated_at"`
}

type OrganizationResponsible struct {
	ID             int `db:"id"`
	OrganizationID int `db:"organization_id"`
	UserID         int `db:"user_id"`
}
