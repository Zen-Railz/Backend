package database

import "database/sql"

type Repository struct {
	database *sql.DB
}
