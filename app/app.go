package app

import (
	"log"

	"github.com/spf13/viper"
	"pingrobot-api.go/pingrobot"
	pkg "pingrobot-api.go/pkg/database"
	"pingrobot-api.go/repository"
	"pingrobot-api.go/service"
	"pingrobot-api.go/transport"
)

func Run() {

	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	viper.ReadInConfig()

	connectionInfo := pkg.ConnectionInfo{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		DBName:   viper.GetString("db.dbname"),
		Username: viper.GetString("db.username"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: viper.GetString("db.password"),
	}

	db, err := pkg.NewPostgresConnection(connectionInfo)
	if err != nil {
		log.Fatal(err)
	}

	repository := repository.NewRepositories(db)

	deps := service.Deps{
		Repos: repository,
	}

	services := service.NewServices(deps)

	handler := transport.NewHadnler(services.WebServices, services.Authorization)
	pingrobot.Run(db)
	handler.Init()
}

//Host     string
//Port     int
//Username string
//DBName   string
//SSLMode  string
//Password string
