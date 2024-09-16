package models

import (
	"time"

	"github.com/google/uuid"
)

type Tender struct {
	ID              int       `db:"id" json:"id,omitempty"`
	Name            string    `db:"name" json:"name"`
	Description     string    `db:"description" json:"description"`
	ServiceType     string    `db:"service_type"  json:"serviceType"`
	Status          string    `db:"status" json:"status"`
	OrganizationID  uuid.UUID `db:"organization_id" json:"organizationId"`
	CreatorUsername string    `db:"creator_username" json:"creatorUsername"`
	Version         int       `db:"version" json:"version,omitempty"`
	CreatedAt       time.Time `db:"created_at" json:"createdAt,omitempty"`
	UpdatedAt       time.Time `db:"updated_at" json:"updatedAt,omitempty"`
}

type UpdateTenderRequest struct {
	Name            *string    `json:"name,omitempty"`
	Description     *string    `json:"description,omitempty"`
	ServiceType     *string    `json:"serviceType,omitempty"`
	Status          *string    `json:"status,omitempty"`
	OrganizationID  *uuid.UUID `json:"organizationId,omitempty"`
	CreatorUsername *string    `json:"creatorUsername,omitempty"`
}
