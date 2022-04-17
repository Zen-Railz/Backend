package railway

import (
	"database/sql"
)

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		database: db,
	}
}
