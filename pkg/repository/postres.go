package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	UrlTable = "url"
)

type Config struct {
	Port     string
	Host     string
	UserName string
	DBName   string
	Password string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("port=%s host=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Port, cfg.Host, cfg.UserName, cfg.DBName, cfg.Password, cfg.SSLMode))

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, err
}
