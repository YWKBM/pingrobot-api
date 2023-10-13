package service

import (
	"context"

	"pingrobot-api.go/domain"
	"pingrobot-api.go/repository"
)

type UserService struct {
	repo repository.UsersRepository
}

func NewUserService(repo repository.UsersRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (u *UserService) SingUp(ctx context.Context, input UserSignUpInput) error {
	user := domain.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}

	err := u.repo.Create(ctx, user)

	if err != nil {
		return err
	}

	return nil
}

func (u *UserService) SignIn(ctx context.Context, id int64) (*domain.User, error) {
	user, err := u.repo.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserService) CreateWebService(ctx context.Context, webService domain.WebSerice) error {
	return u.repo.CreateWebService(ctx, webService)
}
