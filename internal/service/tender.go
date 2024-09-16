package service

import (
	"avito-tender/internal/models"
	"avito-tender/internal/repository"
	"fmt"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
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
	logrus.Info("проверка прав ", "debug")
	err := s.checkUserRights(tender.CreatorUsername, tender.OrganizationID)
	if err != nil {
		return models.Tender{}, err
	}
	logrus.Info("проверили права ", "debug", err)

	return s.tenderRepo.CreateTender(tender)
}

func (s *TenderService) EditTender(tenderid int, tender models.UpdateTenderRequest) (models.Tender, error) {

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

	logrus.Info("проверяем имя ", "debug ")
	userId, err := s.employeeRepo.GetUserIdByUsername(usename)
	if err != nil {
		return err
	}

	logrus.Info("проверяем доступ", "debug")
	access, err := s.employeeRepo.CheckUserOrganization(userId, organizationId)
	if !access {
		return fmt.Errorf("access denied")
	}
	if err != nil {
		return err
	}

	logrus.Info("сверяем доступ", "debug")
	if !access {
		return fmt.Errorf("access denied")
	}
	return nil
}
