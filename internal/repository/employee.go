package repository

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type EmployeePostgres struct {
	db *sqlx.DB
}

func NewEmployeePostgres(db *sqlx.DB) *EmployeePostgres {
	return &EmployeePostgres{db: db}
}

func (r *EmployeePostgres) GetUserIdByUsername(username string) (uuid.UUID, error) {
	logrus.Info("debug")
	var userId uuid.UUID = uuid.Nil
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1", employeeTable)
	err := r.db.Get(&userId, query, username)

	if err != nil {
		logrus.Errorf("Error fetching user ID for username %s: %v", username, err)
		return uuid.Nil, err
	}
	if userId == uuid.Nil {
		return uuid.Nil, fmt.Errorf("user with username %s not found", username)
	}
	return userId, nil
}

func (r *EmployeePostgres) CheckUserOrganization(userId uuid.UUID, organizationId uuid.UUID) (bool, error) {
	var orgResp OrganizationResponsible

	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1 AND organization_id = $2", orgRespTable)
	err := r.db.Get(&orgResp, query, userId, organizationId)
	if err != nil {
		return false, err
	}
	logrus.Info(orgResp, "debug")
	if orgResp.ID == "" {
		return false, err
	}

	return true, err
}
