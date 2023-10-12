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
	return nil
}

func (u *UserService) SignIn(ctx context.Context, input UserSignInInput) error {
	return nil
}

func (u *UserService) CreateWebService(ctx context.Context, webService domain.WebSerice) {

}
