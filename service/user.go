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

	return u.repo.Create(ctx, user)
}

func (u *UserService) SignIn(ctx context.Context, input UserSignInInput) (domain.User, error) {
	email := input.Email
	password := input.Password
	return u.repo.GetUser(ctx, email, password)
}

func (u *UserService) CreateWebService(ctx context.Context, webService domain.WebService) error {
	return u.repo.CreateWebService(ctx, webService)
}
