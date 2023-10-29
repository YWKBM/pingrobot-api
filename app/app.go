package app

import (
	"log"

	pkg "pingrobot-api.go/pkg/database"
	"pingrobot-api.go/repository"
	"pingrobot-api.go/service"
	"pingrobot-api.go/transport"
)

func Run() {

	connectionInfo := pkg.ConnectionInfo{
		Host:     "localhost",
		Port:     5432,
		DBName:   "pingrobot-api",
		Username: "postgres",
		SSLMode:  "disable",
		Password: "0591",
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

	handler.Init()
}

//Host     string
//Port     int
//Username string
//DBName   string
//SSLMode  string
//Password string
