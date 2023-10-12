package repository

import (
	"context"

	"github.com/jinzhu/gorm"
	"pingrobot-api.go/domain"
)

type UsersRepository interface {
	Create(ctx context.Context, user domain.User) error
	GetUserById(ctx context.Context, id int64) (domain.User, error)
}

type WebServiceRepository interface {
	Create(ctx context.Context, webService domain.WebSerice) error
	GetWebServiceByUserId(ctx context.Context, id int64) ([]domain.WebSerice, error)
}

type Repositories struct {
	Users       UsersRepository
	WebServices WebServiceRepository
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		Users:       NewUsersRepo(db),
		WebServices: NewWebSericeRepo(db),
	}
}
