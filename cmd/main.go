package main

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	pet_service "github.com/underbeers/PetService"
	"github.com/underbeers/PetService/pkg/handler"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	//db, err := repository.NewPostgresDB(repository.Config{
	//	Host:     viper.GetString("db.host"),
	//	Port:     viper.GetString("db.port"),
	//	Username: viper.GetString("db.username"),
	//	DBName:   viper.GetString("db.dbname"),
	//	SSLMode:  viper.GetString("db.sslmode"),
	//	Password: os.Getenv("DB_PASSWORD"),
	//})
	//if err != nil {
	//	logrus.Fatalf("failed to initialize db: %s", err.Error())
	//}

	//repos := repository.NewRepository(db)
	//services := service.NewService(repos)
	//handlers := handler.NewHandler(services)

	handlers := handler.Handler{}
	srv := new(pet_service.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
