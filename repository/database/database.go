package database

import (
	"database/sql"
	"zenrailz/anomaly"
	"zenrailz/code"
	"zenrailz/environment"
	"zenrailz/repository/common"

	_ "github.com/lib/pq"
)

func New() (*sql.DB, *anomaly.ServiceError) {
	uri, uriErr := environment.DatabaseUri()
	if uriErr != nil {
		return nil, uriErr.Trace()
	}

	db, sqlOpenErr := sql.Open("postgres", uri)
	if sqlOpenErr != nil {
		err := common.ParseError(code.DatabaseConnectionFailure, "Unable to open connection to database.", sqlOpenErr)
		return nil, err.Trace()
	}

	return db, nil
}
