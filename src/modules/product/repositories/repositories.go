package repositories

import (
	"database/sql"
)

type Repository struct {
	DB      *sql.DB
	Product ProductRepository
}

func New(db *sql.DB) *Repository {
	return &Repository{
		DB:      db,
		Product: &productRepository{db},
	}
}
