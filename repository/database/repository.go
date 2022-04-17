package database

import (
	"database/sql"
	"zenrailz/code"
	"zenrailz/errorr"
	"zenrailz/repository/common"
)

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		database: db,
	}
}

func (d *Repository) Ping() errorr.Entity {
	if pingErr := d.database.Ping(); pingErr != nil {
		err := common.ParseError(code.DatabasePingFailure, "Unable to reach database. Database did not respond.", pingErr)
		return err.Trace()
	}
	return nil
}
