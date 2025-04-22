package repository

import (
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func NewDB() (*sqlx.DB, error) {

	if err := godotenv.Load(); err != nil {
		logrus.Error("cannot read enviromental values")

	}
	db, err := NewPostgresDB(Config{
		Port:     viper.GetString("db.port"),
		Host:     viper.GetString("db.host"),
		UserName: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		return nil, err
	}

	return db, nil
}
