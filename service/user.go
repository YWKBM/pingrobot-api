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

	err := u.repo.Create(ctx, &user)

	if err != nil {
		return err
	}

	return nil
}

func (u *UserService) SignIn(ctx context.Context, input UserSignInInput) (*domain.User, error) {
	email := input.Email
	password := input.Password
	user, err := u.repo.GetUser(ctx, email, password)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserService) CreateWebService(ctx context.Context, webService domain.WebSerice) error {
	return u.repo.CreateWebService(ctx, webService)
}
