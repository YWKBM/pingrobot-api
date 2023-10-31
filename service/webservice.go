package service

import (
	"pingrobot-api.go/domain"
	"pingrobot-api.go/repository"
)

type WebServiceService struct {
	repo repository.WebServiceRepository
}

func NewWebSericeService(repo repository.WebServiceRepository) *WebServiceService {
	return &WebServiceService{repo: repo}
}

func (w *WebServiceService) Create(userId int, webService domain.WebService) (int, error) {
	if err := webService.Validate(); err != nil {
		return 0, err
	}

	return w.repo.Create(userId, webService)
}

func (w *WebServiceService) GetAll(userId int) ([]domain.WebService, error) {
	return w.repo.GetAll(userId)
}

func (w *WebServiceService) GetById(userId int, webServiceId int) (domain.WebService, error) {
	return w.repo.GetById(userId, webServiceId)
}

func (w *WebServiceService) Delete(userId, webServiceId int) error {
	return w.repo.Delete(userId, webServiceId)
}

func (w *WebServiceService) Update(userId, webServiceId int, input domain.UpdateWebServiceInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return w.repo.Update(userId, webServiceId, input)
}
