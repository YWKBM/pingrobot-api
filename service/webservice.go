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

func (w *WebSericeService) GetWebServiceByUserId(ctx context.Context, id int64) ([]domain.WebSerice, error) {
	return w.repo.GetWebServiceByUserId(ctx, id)
}

//TODO: Connection to ping-functional with this package
