package configuration

import "database/sql"

type Repository struct {
	database *sql.DB
}
