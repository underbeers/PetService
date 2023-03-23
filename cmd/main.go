package main

import (
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/underbeers/PetService/pkg/handler"
	"github.com/underbeers/PetService/pkg/repository"
	pet_service "github.com/underbeers/PetService/pkg/server"
	"github.com/underbeers/PetService/pkg/service"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	cfg := repository.GetConfig()
	db, err := repository.NewPostgresDB(*cfg.DB)
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos, cfg)
	handlers := handler.NewHandler(services)

	srv := new(pet_service.Server)
	if err := srv.Run(handlers.InitRoutes(), cfg); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}

}
