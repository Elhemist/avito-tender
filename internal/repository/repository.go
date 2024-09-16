package repository

import (
	"avito-tender/internal/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Tender interface {
	GetAllTenders() ([]models.Tender, error)
	GetUserTenders(username string) ([]models.Tender, error)
	CreateTender(tender models.Tender) (models.Tender, error)
	EditTender(tenderId int, tender models.UpdateTenderRequest) (models.Tender, error)
	RollbackTender(tender models.Tender) error
	GetHistoryTender(tenderId int, version int) (models.Tender, error)
}

type Bid interface {
	GetUserBids(username string) ([]models.Bid, error)
	GetTenderBids(tenderId int) ([]models.Bid, error)
	CreateBid(tender models.Bid) (models.Bid, error)
	AddBidToHistory(bid models.Bid) error
	EditBid(bidId int, bid models.UpdateBidRequest) (models.Bid, error)
	GetHistoryBid(bidId int, version int) (models.Bid, error)
	RollbackBid(bid models.Bid) error
}

type Employee interface {
	GetUserIdByUsername(username string) (uuid.UUID, error)
	CheckUserOrganization(userId uuid.UUID, organizationId uuid.UUID) (bool, error)
}

type Repository struct {
	Tender
	Bid
	Employee
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Tender:   NewTenderPostgres(db),
		Bid:      NewBidPostgres(db),
		Employee: NewEmployeePostgres(db),
	}
}
