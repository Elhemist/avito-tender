package models

import (
	"time"

	"github.com/google/uuid"
)

type Bid struct {
	ID              int       `db:"id" json:"id,omitempty"`
	Name            string    `db:"name" json:"name"`
	Description     string    `db:"description" json:"description"`
	Status          string    `db:"status" json:"status"`
	TenderID        int       `db:"tender_id" json:"tenderId"`
	OrganizationID  uuid.UUID `db:"organization_id" json:"organizationId"`
	CreatorUsername string    `db:"creator_username" json:"creatorUsername"`
	Version         int       `db:"version" json:"version"`
	CreatedAt       time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt       time.Time `db:"updated_at" json:"updatedAt"`
}
type UpdateBidRequest struct {
	Name            *string    `db:"name" json:"name,omitempty"`
	Description     *string    `db:"description" json:"description,omitempty"`
	Status          *string    `db:"status" json:"status,omitempty"`
	TenderID        *int       `db:"tender_id" json:"tenderId,omitempty"`
	OrganizationID  *uuid.UUID `db:"organization_id" json:"organizationId,omitempty"`
	CreatorUsername *string    `db:"creator_username" json:"creatorUsername,omitempty"`
}

type Review struct {
	ID        int       `db:"id" json:"id"`
	BidID     int       `db:"bid_id" json:"bidId"`
	UserID    int       `db:"user_id" json:"userId"`
	Comment   string    `db:"comment" json:"comment"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
}

type Decision struct {
	ID        int       `db:"id" json:"id"`
	BidID     int       `db:"bid_id" json:"bidId"`
	UserID    int       `db:"user_id" json:"userId"`
	Action    string    `db:"action" json:"action"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
}
