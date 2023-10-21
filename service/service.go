package service

import (
	"context"

	"pingrobot-api.go/domain"
	"pingrobot-api.go/repository"
)

type UserSignUpInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `jsong:"password"`
}

type UserSignInInput struct {
	Email    string `json:"email"`
	Password string `jsong:"password"`
}

type Users interface {
	SingUp(ctx context.Context, input UserSignUpInput) error
	SignIn(ctx context.Context, input UserSignInInput) (domain.User, error) //TODO: User verif, id only for testing - use UserSignInput, returning user only for testing
	CreateWebService(ctx context.Context, webService domain.WebService) error
}

type WebServices interface {
	GetWebServiceByUserId(ctx context.Context) ([]domain.WebService, error)
}

type Services struct {
	Users      Users
	WebSerices WebServices
}

type Deps struct {
	Repos *repository.Repositories
}

func NewServices(deps Deps) *Services {
	userService := NewUserService(deps.Repos.Users)
	webSericeService := NewWebSericeService(deps.Repos.WebServices)

	return &Services{
		Users:      userService,
		WebSerices: webSericeService,
	}
}
