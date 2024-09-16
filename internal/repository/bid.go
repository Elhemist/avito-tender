package repository

import (
	"avito-tender/internal/models"
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type BidPostgres struct {
	db *sqlx.DB
}

type historyBid struct {
	ID              int       `db:"id"`
	BidId           int       `db:"bid_id"`
	Name            string    `db:"name"`
	Description     string    `db:"description"`
	Status          string    `db:"status"`
	TenderID        int       `db:"tender_id" `
	OrganizationID  uuid.UUID `db:"organization_id"`
	CreatorUsername string    `db:"creator_username"`
	Version         int       `db:"version" `
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt       time.Time `db:"updated_at"`
}

func NewBidPostgres(db *sqlx.DB) *BidPostgres {
	return &BidPostgres{db: db}
}

func (r *BidPostgres) GetUserBids(username string) ([]models.Bid, error) {
	var bidList []models.Bid
	query := fmt.Sprintf("SELECT * FROM %s WHERE creator_username=$1", bidTable)
	err := r.db.Select(&bidList, query, username)
	return bidList, err
}

func (r *BidPostgres) GetTenderBids(tenderId int) ([]models.Bid, error) {
	var bidList []models.Bid
	query := fmt.Sprintf("SELECT * FROM %s WHERE tender_id=$1", bidTable)
	err := r.db.Select(&bidList, query, tenderId)
	return bidList, err
}

func (r *BidPostgres) CreateBid(bid models.Bid) (models.Bid, error) {
	query := fmt.Sprintf("INSERT INTO %s (name, description, status, tender_id,  organization_id, creator_username) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, version, created_at, updated_at", bidTable)
	err := r.db.QueryRow(
		query,
		bid.Name,
		bid.Description,
		bid.Status,
		bid.TenderID,
		bid.OrganizationID,
		bid.CreatorUsername,
	).Scan(&bid.ID, &bid.Version, &bid.CreatedAt, &bid.UpdatedAt)
	if err != nil {
		return bid, err
	}

	err = r.AddBidToHistory(bid)

	return bid, err
}

func (r *BidPostgres) AddBidToHistory(bid models.Bid) error {
	query := fmt.Sprintf("INSERT INTO %s (bid_id, name, description, status, tender_id, organization_id, creator_username, version, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)", bidHistoryTable)
	_, err := r.db.Exec(query, bid.ID, bid.Name, bid.Description, bid.Status, bid.TenderID, bid.OrganizationID, bid.CreatorUsername, bid.Version, bid.CreatedAt, bid.UpdatedAt)
	return err
}

func (r *BidPostgres) EditBid(bidId int, bid models.UpdateBidRequest) (models.Bid, error) {
	var bidNew models.Bid
	version, err := r.GetBidVersion(bidId)
	if err != nil {
		return bidNew, err
	}

	query := fmt.Sprintf("UPDATE %s SET", bidTable)
	if bid.Name != nil {
		query += ` name = '` + *bid.Name + `',`
	}
	if bid.Description != nil {
		query += ` description = '` + *bid.Description + `',`
	}
	if bid.Status != nil {
		query += ` status = '` + *bid.Status + `',`
	}
	if bid.OrganizationID != nil {
		strTender := *bid.OrganizationID
		query += ` organization_id = '` + strTender.String() + `',`
	}
	if bid.CreatorUsername != nil {
		query += ` creator_username = '` + *bid.CreatorUsername + `',`
	}
	newVersion := version + 1
	query += ` version = ` + strconv.Itoa(newVersion)

	query += `, updated_at = NOW()`

	query += ` WHERE id = $1`

	_, err = r.db.Exec(query, bidId)
	if err != nil {
		return bidNew, err
	}

	query = fmt.Sprintf("SELECT * FROM %s WHERE id = $1", bidTable)
	err = r.db.Get(&bidNew, query, bidId)
	if err != nil {
		return models.Bid{}, err
	}

	err = r.AddBidToHistory(bidNew)

	return bidNew, err
}

func (r *BidPostgres) GetBidVersion(bidId int) (int, error) {
	var bidVersion int
	query := fmt.Sprintf(`SELECT COUNT(*) FROM %s WHERE bid_id = $1`, bidHistoryTable)
	err := r.db.Get(&bidVersion, query, bidId)
	return bidVersion, err
}

func (r *BidPostgres) GetHistoryBid(bidId int, version int) (models.Bid, error) {
	var historyBid historyBid
	query := fmt.Sprintf("SELECT * FROM %s WHERE bid_id=$1 AND version=$2", bidHistoryTable)
	err := r.db.Get(&historyBid, query, bidId, version)
	return models.Bid{
		ID:              historyBid.BidId,
		Name:            historyBid.Name,
		Description:     historyBid.Description,
		Status:          historyBid.Status,
		TenderID:        historyBid.TenderID,
		OrganizationID:  historyBid.OrganizationID,
		CreatorUsername: historyBid.CreatorUsername,
		Version:         historyBid.Version,
		CreatedAt:       historyBid.CreatedAt,
		UpdatedAt:       historyBid.UpdatedAt,
	}, err
}

func (r *BidPostgres) RollbackBid(bid models.Bid) error {
	query := fmt.Sprintf("UPDATE %s SET name = $1, description = $2, status = $3, organization_id = $4, creator_username = $5, tender_id = $6, version = $7, updated_at = NOW() WHERE id = $8", bidTable)
	_, err := r.db.Exec(query, bid.Name, bid.Description, bid.Status, bid.OrganizationID, bid.CreatorUsername, bid.TenderID, bid.Version, bid.ID)
	return err
}
