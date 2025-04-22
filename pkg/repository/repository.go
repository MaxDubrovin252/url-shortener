package repository

import (
	"github.com/jmoiron/sqlx"
)

type URL interface {
	Create(url string, alias string) (int64, error)
	Get(alias string) (string, error)
	Delete(alias string) error
}

type Repository struct {
	URL
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		URL: NewURL(db),
	}
}
