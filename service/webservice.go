package service

import (
	"context"

	"pingrobot-api.go/domain"
)

type WebServiceRepository interface {
	Create(ctx context.Context, webService domain.WebSerice) error
	GetWebServiceByUserId(ctx context.Context, id int64) (domain.WebSerice, error)
}

type WebSerice struct {
	repo WebServiceRepository
}

func NewWebSericeService(repo WebServiceRepository) *WebSerice {
	return &WebSerice{repo: repo}
}

func (w *WebSerice) Create(ctx context.Context, webService domain.WebSerice) error {
	return w.repo.Create(ctx, webService)
}

func (w *WebSerice) GetWebServiceByUserId(ctx context.Context, id int64) (domain.WebSerice, error) {
	return w.repo.GetWebServiceByUserId(ctx, id)
}

//TODO: Connection to ping-functional with this package
