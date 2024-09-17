package service

import (
	"avito-tender/internal/models"
	"avito-tender/internal/repository"
	"fmt"

	"github.com/google/uuid"
)

type TenderService struct {
	tenderRepo   repository.Tender
	employeeRepo repository.Employee
}

func NewTenderService(repoTender repository.Tender, repoUser repository.Employee) *TenderService {
	return &TenderService{tenderRepo: repoTender, employeeRepo: repoUser}
}

func (s *TenderService) GetAllTenders() ([]models.Tender, error) {
	return s.tenderRepo.GetAllTenders()
}

func (s *TenderService) GetUserTenders(username string) ([]models.Tender, error) {
	return s.tenderRepo.GetUserTenders(username)
}

func (s *TenderService) CreateTender(tender models.Tender) (models.Tender, error) {
	err := s.checkUserRights(tender.CreatorUsername, tender.OrganizationID)
	if err != nil {
		return models.Tender{}, err
	}
	return s.tenderRepo.CreateTender(tender)
}

func (s *TenderService) EditTender(tenderid int, tender models.UpdateTenderRequest) (models.Tender, error) {
	exist, err := s.tenderRepo.DoesTenderExists(tenderid)
	if err != nil {
		return models.Tender{}, err
	}
	if !exist {
		return models.Tender{}, NO_TENDER
	}

	return s.tenderRepo.EditTender(tenderid, tender)
}

func (s *TenderService) RollbackTender(tenderId int, versionId int) (models.Tender, error) {
	tender, err := s.tenderRepo.GetHistoryTender(tenderId, versionId)
	if err != nil {
		return tender, err
	}
	err = s.tenderRepo.RollbackTender(tender)

	return tender, err
}

func (s *TenderService) checkUserRights(usename string, organizationId uuid.UUID) error {

	userId, err := s.employeeRepo.GetUserIdByUsername(usename)
	if err != nil {
		return NO_USER
	}

	access, err := s.employeeRepo.CheckUserOrganization(userId, organizationId)
	if !access {
		return fmt.Errorf("access denied")
	}
	if err != nil {
		return err
	}
	return nil
}
