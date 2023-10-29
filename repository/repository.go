package repository

import (
	"database/sql"

	"pingrobot-api.go/domain"
)

type WebServiceRepository interface {
	Create(userId int, webService domain.WebService) (int, error)
	GetAll(userId int) ([]domain.WebService, error)
	GetById(userId int, webServiceId int) (domain.WebService, error)
	Delete(userId int, webServiceId int) error
	Update(userId, webServiceId int, input domain.UpdateWebServiceInput) error
}

type Authorization interface {
	CreateUser(user domain.User) (int, error)
	GetUser(name, password string) (domain.User, error)
}

type Repositories struct {
	WebServices   WebServiceRepository
	Authorization Authorization
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		WebServices:   NewWebSericeRepo(db),
		Authorization: NewAuthorizationRepo(db),
	}
}
