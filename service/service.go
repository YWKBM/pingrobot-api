package service

import (
	"pingrobot-api.go/domain"
	"pingrobot-api.go/repository"
)

type WebServices interface {
	Create(usderId int, webSevice domain.WebService) (int, error)
	GetAll(userId int) ([]domain.WebService, error)
	GetById(userId int, webServiceId int) (domain.WebService, error)
	Delete(userId int, webServiceId int) error
	Update(userId int, webSeviceId int, input domain.UpdateWebServiceInput) error
}

type Authorization interface {
	CreateUser(user domain.User) (int, error)
	GenerateToken(name, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type Services struct {
	WebServices   WebServices
	Authorization Authorization
}

type Deps struct {
	Repos *repository.Repositories
}

func NewServices(deps Deps) *Services {
	webServiceService := NewWebSericeService(deps.Repos.WebServices)
	authService := NewAuthService(deps.Repos.Authorization)

	return &Services{
		WebServices:   webServiceService,
		Authorization: authService,
	}
}
