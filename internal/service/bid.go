package service

import (
	"avito-tender/internal/models"
	"avito-tender/internal/repository"
	"fmt"

	"github.com/google/uuid"
)

type BidService struct {
	employeeRepo repository.Employee
	bidRepo      repository.Bid
}

func NewBidService(repoBid repository.Bid, repoUser repository.Employee) *BidService {
	return &BidService{bidRepo: repoBid, employeeRepo: repoUser}
}

func (s *BidService) CreateBid(bid models.Bid) (models.Bid, error) {
	err := s.checkUserRights(bid.CreatorUsername, bid.OrganizationID)
	if err != nil {
		return models.Bid{}, err
	}

	return s.bidRepo.CreateBid(bid)
}

func (s *BidService) GetUserBids(username string) ([]models.Bid, error) {
	return s.bidRepo.GetUserBids(username)
}

func (s *BidService) GetTenderBids(tenderid int) ([]models.Bid, error) {
	return s.bidRepo.GetTenderBids(tenderid)
}

func (s *BidService) EditBid(bidId int, bid models.UpdateBidRequest) (models.Bid, error) {

	return s.bidRepo.EditBid(bidId, bid)
}

func (s *BidService) RollbackBid(bidId int, versionId int) (models.Bid, error) {
	bid, err := s.bidRepo.GetHistoryBid(bidId, versionId)
	if err != nil {
		return bid, err
	}
	err = s.bidRepo.RollbackBid(bid)

	return bid, err
}

func (s *BidService) checkUserRights(usename string, organizationId uuid.UUID) error {
	userId, err := s.employeeRepo.GetUserIdByUsername(usename)
	if err != nil {
		return err
	}

	access, err := s.employeeRepo.CheckUserOrganization(userId, organizationId)
	if err != nil {
		return err
	}

	if !access {
		return fmt.Errorf("access denied")
	}
	return nil
}
