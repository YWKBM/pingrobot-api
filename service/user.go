package service

import (
	"context"
	"time"

	"pingrobot-api.go/domain"
)

type UsersRepository interface {
	Create(ctx context.Context, user domain.User) error
	GetUserById(ctx context.Context, id int64) (domain.User, error)
}

type User struct {
	repo UsersRepository
}

func NewUserService(repo UsersRepository) *User {
	return &User{
		repo: repo,
	}
}

func (u *User) Create(ctx context.Context, user domain.User) error {
	user.CreatedAt = time.Now()
	user.LastVisitAt = time.Now()

	return u.repo.Create(ctx, user)
}

func (u *User) GetUserById(ctx context.Context, id int64) (domain.User, error) {
	return u.repo.GetUserById(ctx, id)
}
