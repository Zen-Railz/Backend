package database

import (
	"database/sql"
	"zenrailz/anomaly"
	"zenrailz/code"
	"zenrailz/environment"

	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, *anomaly.ServiceError) {
	uri, uriErr := environment.DatabaseUri()
	if uriErr != nil {
		return nil, uriErr.Trace()
	}

	db, sqlOpenErr := sql.Open("postgres", uri)
	if sqlOpenErr != nil {
		err := parseError(code.DatabaseConnectionFailure, "Unable to open connection to database.", sqlOpenErr)
		return nil, err.Trace()
	}

	return db, nil
}

func Health(db *sql.DB) *anomaly.ServiceError {
	if pingErr := db.Ping(); pingErr != nil {
		err := parseError(code.DatabasePingFailure, "Unable to reach database. Database did not respond.", pingErr)
		return err.Trace()
	}
	return nil
}
