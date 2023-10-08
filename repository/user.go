package repository

import (
	"context"

	"github.com/jinzhu/gorm"
	"pingrobot-api.go/domain"
)

type UserRepo struct {
	db *gorm.DB
}

func (u *UserRepo) Create(ctx context.Context, user domain.User) error {
	return u.db.Create(user).Error
}

func (u *UserRepo) GetUserById(ctx context.Context, id uint) (*domain.User, error) {
	var user domain.User
	err := u.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
