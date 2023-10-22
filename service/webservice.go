package service

import (
	"context"

	"pingrobot-api.go/domain"
	"pingrobot-api.go/repository"
)

type WebSericeService struct {
	repo repository.WebServiceRepository
}

func NewWebSericeService(repo repository.WebServiceRepository) *WebSericeService {
	return &WebSericeService{repo: repo}
}

func (w *WebSericeService) GetAllWebServices(ctx context.Context) ([]domain.WebService, error) {
	return w.repo.GetAllWebServices(ctx)
}

//TODO: Connection to ping-functional with this package
