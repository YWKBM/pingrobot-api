package repository

import (
	"database/sql"
	"context"

	"pingrobot-api.go/domain"
)

type UsersRepository interface {
	Create(ctx context.Context, user domain.User) error
	GetUser(ctx context.Context, email string, password string) (domain.User, error)
	CreateWebService(ctx context.Context, webService domain.WebService) error
}

type WebServiceRepository interface {
	GetAllWebServices(ctx context.Context) ([]domain.WebService, error)
}

type Repositories struct {
	Users       UsersRepository
	WebServices WebServiceRepository
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		Users:       NewUsersRepo(db),
		WebServices: NewWebSericeRepo(db),
	}
}
