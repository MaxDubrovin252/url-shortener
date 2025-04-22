package main

import (
	"url-shortener/config"
	"url-shortener/internal/http/handler"
	"url-shortener/internal/server"
	"url-shortener/pkg/repository"
	"url-shortener/pkg/service"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	if err := config.InitConfig(); err != nil {
		logrus.Errorf("cannot read config:%s", err)
	}

	db, err := repository.NewDB()

	if err != nil {
		logrus.Error(err.Error())
	}

	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handlers := handler.NewHandler(service)
	srv := new(server.SRV)

	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Error(err.Error())
	}

}
