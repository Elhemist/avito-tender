package repository

import (
	"avito-tender/internal/models"
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type historyTender struct {
	ID              int       `db:"id" `
	TenderId        int       `db:"tender_id"`
	Name            string    `db:"name" `
	Description     string    `db:"description"`
	ServiceType     string    `db:"service_type"`
	Status          string    `db:"status"`
	OrganizationID  uuid.UUID `db:"organization_id"`
	CreatorUsername string    `db:"creator_username"`
	Version         int       `db:"version"`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt       time.Time `db:"updated_at"`
}

type TenderPostgres struct {
	db *sqlx.DB
}

func NewTenderPostgres(db *sqlx.DB) *TenderPostgres {
	return &TenderPostgres{db: db}
}

func (r *TenderPostgres) GetAllTenders() ([]models.Tender, error) {
	var tenderList []models.Tender
	query := fmt.Sprintf("SELECT * FROM %s", tenderTable)
	err := r.db.Select(&tenderList, query)
	return tenderList, err
}

func (r *TenderPostgres) GetUserTenders(username string) ([]models.Tender, error) {
	var tenderList []models.Tender
	query := fmt.Sprintf("SELECT * FROM %s WHERE creator_username=$1", tenderTable)
	err := r.db.Select(&tenderList, query, username)
	return tenderList, err
}

func (r *TenderPostgres) CreateTender(tender models.Tender) (models.Tender, error) {
	// var newTenderId int
	// var tenderNew models.Tender
	query := fmt.Sprintf("INSERT INTO %s (name, description, service_type, status, organization_id, creator_username) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, version, created_at, updated_at", tenderTable)
	err := r.db.QueryRow(query, tender.Name, tender.Description, tender.ServiceType, tender.Status, tender.OrganizationID, tender.CreatorUsername).Scan(&tender.ID, &tender.Version, &tender.CreatedAt, &tender.UpdatedAt)
	if err != nil {
		return tender, err
	}

	logrus.Info("встувили ", "debug", err)

	err = r.AddTenderToHistory(tender)

	return tender, err
}

func (r *TenderPostgres) EditTender(tenderId int, tender models.UpdateTenderRequest) (models.Tender, error) {
	var tenderNew models.Tender
	version, err := r.GetTenderVersion(tenderId)
	if err != nil {
		return models.Tender{}, err
	}

	query := fmt.Sprintf("UPDATE %s SET", tenderTable)
	args := []interface{}{}
	if tender.Name != nil {
		query += ` name = '` + *tender.Name + `',`
	}
	if tender.Description != nil {
		query += ` description = '` + *tender.Description + `',`
	}
	if tender.ServiceType != nil {
		query += ` service_type = '` + *tender.ServiceType + `',`
	}
	if tender.Status != nil {
		query += ` status = '` + *tender.Status + `',`
	}
	if tender.OrganizationID != nil {
		strTender := *tender.OrganizationID
		query += ` organization_id = ` + strTender.String() + `,`
	}
	if tender.CreatorUsername != nil {
		query += ` creator_username = '` + *tender.CreatorUsername + `',`
	}
	newVersion := version + 1
	query += ` version = ` + strconv.Itoa(newVersion)

	query += `, updated_at = NOW()`

	query += ` WHERE id = $1`
	logrus.Info("debug3", query)
	logrus.Info("debug4", args)
	_, err = r.db.Exec(query, tenderId)
	if err != nil {
		logrus.Info("debug 123", query)
		return models.Tender{}, err
	}

	query = fmt.Sprintf("SELECT * FROM %s WHERE id = $1", tenderTable)
	err = r.db.Get(&tenderNew, query, tenderId)
	if err != nil {
		logrus.Info("debug 1234", query)
		return models.Tender{}, err
	}
	logrus.Info("debug 1235325", query)

	err = r.AddTenderToHistory(tenderNew)

	return tenderNew, err
}

func (r *TenderPostgres) AddTenderToHistory(tender models.Tender) error {
	query := fmt.Sprintf("INSERT INTO %s (tender_id, name, description, service_type, status, organization_id, creator_username, version, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)", tenderHistoryTable)
	_, err := r.db.Exec(query, tender.ID, tender.Name, tender.Description, tender.ServiceType, tender.Status, tender.OrganizationID, tender.CreatorUsername, tender.Version, tender.CreatedAt, tender.UpdatedAt)
	return err
}

func (r *TenderPostgres) RollbackTender(tender models.Tender) error {
	query := fmt.Sprintf("UPDATE %s SET name = $1, description = $2, service_type = $3, status = $4, organization_id = $5, creator_username = $6, version = $7, updated_at = NOW() WHERE id = $8", tenderTable)
	_, err := r.db.Exec(query, tender.Name, tender.Description, tender.ServiceType, tender.Status, tender.OrganizationID, tender.CreatorUsername, tender.Version, tender.ID)
	return err
}

func (r *TenderPostgres) GetTenderVersion(tenderId int) (int, error) {
	var tenderVersion int
	query := fmt.Sprintf(`SELECT COUNT(*) FROM %s WHERE tender_id =$1`, tenderHistoryTable)
	err := r.db.Get(&tenderVersion, query, tenderId)
	return tenderVersion, err
}

func (r *TenderPostgres) GetHistoryTender(tenderId int, version int) (models.Tender, error) {
	var historyTender historyTender
	query := fmt.Sprintf("SELECT * FROM %s WHERE tender_id=$1 AND version=$2", tenderHistoryTable)
	err := r.db.Get(&historyTender, query, tenderId, version)
	tender := models.Tender{
		ID:              historyTender.TenderId,
		Name:            historyTender.Name,
		Description:     historyTender.Description,
		ServiceType:     historyTender.ServiceType,
		Status:          historyTender.Status,
		OrganizationID:  historyTender.OrganizationID,
		CreatorUsername: historyTender.CreatorUsername,
		Version:         historyTender.Version,
		CreatedAt:       historyTender.CreatedAt,
		UpdatedAt:       historyTender.UpdatedAt}
	return tender, err
}
