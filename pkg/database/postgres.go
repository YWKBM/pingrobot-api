package pkg

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"pingrobot-api.go/domain"
)

type ConnectionInfo struct {
	Host     string
	Port     int
	Username string
	DBName   string
	SSLMode  string
	Password string
}

func NewPostgresConnection(info ConnectionInfo) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open((fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=%s password=%s",
		info.Host, info.Port, info.Username, info.DBName, info.SSLMode, info.Password))), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, _ := db.DB()

	err = sqlDB.Ping()
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(domain.User{}, domain.WebSerice{})

	return db, nil
}
