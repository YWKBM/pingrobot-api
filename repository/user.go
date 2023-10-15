package repository

import (
	"context"

	"gorm.io/gorm"
	"pingrobot-api.go/domain"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUsersRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db}
}

func (u *UserRepo) Create(ctx context.Context, user *domain.User) error {
	return u.db.Create(user).Error
}

func (u *UserRepo) GetUser(ctx context.Context, email string, password string) (*domain.User, error) {
	var user domain.User

	err := u.db.Where("email = ? AND password = ?", email, password).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserRepo) CreateWebService(ctx context.Context, webService domain.WebSerice) error {
	return u.db.Create(webService).Error
}
