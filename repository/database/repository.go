package database

import (
	"database/sql"
	"zenrailz/anomaly"
	"zenrailz/code"
	"zenrailz/repository/common"
)

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		database: db,
	}
}

type Repository struct {
	database *sql.DB
}

func (d *Repository) Ping() *anomaly.ServiceError {
	if pingErr := d.database.Ping(); pingErr != nil {
		err := common.ParseError(code.DatabasePingFailure, "Unable to reach database. Database did not respond.", pingErr)
		return err.Trace()
	}
	return nil
}
