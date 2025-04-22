package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type UrlStorage struct {
	db *sqlx.DB
}

func NewURL(db *sqlx.DB) *UrlStorage {
	return &UrlStorage{db: db}
}

func (r *UrlStorage) Create(url string, alias string) (int64, error) {
	var id int64

	const operation = "repository.storage.create"

	query := fmt.Sprintf("INSERT INTO %s (url,alias) VALUES($1,$2) RETURNING id", UrlTable)

	row := r.db.QueryRow(query, url, alias)

	if err := row.Scan(&id); err != nil {

		return 0, fmt.Errorf("%s:%w", operation, err)
	}

	return id, nil

}

func (r *UrlStorage) Get(alias string) (string, error) {
	const operation = "repos.url.get"
	query := fmt.Sprintf(`SELECT url FROM %s  WHERE alias = $1`, UrlTable)

	stmt, err := r.db.Prepare(query)

	if err != nil {
		return "", err
	}

	var res string

	err = stmt.QueryRow(alias).Scan(&res)

	if err != nil {
		return "", fmt.Errorf("%s:%w", operation, ErrNotFound)
	}

	return res, nil

}
func (r *UrlStorage) Delete(alias string) error {
	const operation = "repos.url.delete"
	query := fmt.Sprintf("DELETE  FROM  %s  WHERE alias = $1", UrlTable)

	_, err := r.db.Exec(query, alias)

	if err != nil {
		return fmt.Errorf("%s:%w", operation, err)
	}

	return nil

}
