package service

import (
	"avito-tender/internal/models"
	"avito-tender/internal/repository"
)

const NO_TENDER = NotExistError("No tenders with this id exist")
const NO_BID = NotExistError("No bid with this id exist")
const NO_USER = NotExistError("No such user exist ")

type NotExistError string

type Tender interface {
	GetAllTenders() ([]models.Tender, error)
	GetUserTenders(username string) ([]models.Tender, error)
	CreateTender(tender models.Tender) (models.Tender, error)
	EditTender(tenderid int, tender models.UpdateTenderRequest) (models.Tender, error)
	RollbackTender(tenderId int, versionId int) (models.Tender, error)
}
type Bid interface {
	CreateBid(bid models.Bid) (models.Bid, error)
	GetUserBids(username string) ([]models.Bid, error)
	GetTenderBids(tenderid int) ([]models.Bid, error)
	EditBid(bidId int, bid models.UpdateBidRequest) (models.Bid, error)
	RollbackBid(bidId int, versionId int) (models.Bid, error)
}
type Service struct {
	Tender
	Bid
}

func (e NotExistError) Error() string {
	return string(e)
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Tender: NewTenderService(repos.Tender, repos.Employee),
		Bid:    NewBidService(repos.Bid, repos.Employee, repos.Tender),
	}
}
