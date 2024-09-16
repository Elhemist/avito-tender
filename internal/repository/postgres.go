package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	tenderTable        = "tender"
	tenderHistoryTable = "tender_history"
	employeeTable      = "employee"
	orgRespTable       = "organization_responsible"
	bidTable           = "bid"
	bidHistoryTable    = "bid_history"

// segmentTable        = "segments"
// userTable           = "users"
)

type Config struct {
	Host     string
	Port     string
	Usename  string
	Password string
	DBName   string
	SSLmode  string
}

type OrganizationResponsible struct {
	ID             string `db:"id"`
	OrganizationID string `db:"organization_id"`
	UserID         string `db:"user_id"`
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Usename, cfg.DBName, cfg.Password, cfg.SSLmode))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

// func InitTables(db *sqlx.DB) {
// 	query := `
//     DROP TABLE IF EXISTS active_segments;
// 	DROP TABLE IF EXISTS users;
//     DROP TABLE IF EXISTS segments;
// 	CREATE TABLE users(
// 		id         numeric         primary key
// 	);
// 	CREATE TABLE segments(
// 		id         serial          primary key,
// 		name       varchar(255)    UNIQUE not null
// 	);
// 	CREATE TABLE active_segments(
// 		userId     INTEGER         REFERENCES users (id) ON DELETE CASCADE,
// 		segmentId  INTEGER         REFERENCES segments (id) ON DELETE CASCADE,
// 		expiration TIMESTAMP,
// 		PRIMARY KEY(userId, segmentId)
// 	);
// 	`
// 	_, _ = db.Exec(query)
// }

// func FillTables(db *sqlx.DB) {
// 	query := `
// 	INSERT INTO users(id) VALUES (12);
// 	INSERT INTO users(id) VALUES (13);
// 	INSERT INTO users(id) VALUES (14);
// 	INSERT INTO users(id) VALUES (15);
// 	INSERT INTO users(id) VALUES (16);
// 	INSERT INTO users(id) VALUES (17);
// 	INSERT INTO users(id) VALUES (18);

// 	INSERT INTO segments(name) VALUES ('Sale50');
// 	INSERT INTO segments(name) VALUES ('Sale60');
// 	INSERT INTO segments(name) VALUES ('Sale70');
// 	INSERT INTO segments(name) VALUES ('Sale80');
// 	INSERT INTO segments(name) VALUES ('Sale90');
// 	INSERT INTO segments(name) VALUES ('Sale');
// 	`
// 	_, _ = db.Exec(query)
// }

// func DeleteExpired(db *sqlx.DB) {
// 	query := `DELETE FROM active_segments  WHERE expiration <=now()`
// 	for {
// 		_, err := db.Exec(query)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		time.Sleep(time.Second * 5)
// 	}
// }
