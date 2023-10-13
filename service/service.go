package service

import (
	"context"

	"pingrobot-api.go/domain"
	"pingrobot-api.go/repository"
)

type UserSignUpInput struct {
	Name     string
	Email    string
	Password string
}

//type UserSignInInput struct {
//	Email    string
//	Password string
//}

type Users interface {
	SingUp(ctx context.Context, input UserSignUpInput) error
	SignIn(ctx context.Context, id int64) (*domain.User, error) //TODO: User verif, id only for testing - use UserSignInput, returning user only for testing
	CreateWebService(ctx context.Context, webService domain.WebSerice) error
}

type WebServices interface {
	GetWebServiceByUserId(ctx context.Context, id int64) ([]domain.WebSerice, error)
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
