package repository

import (
	"context"

	"github.com/jinzhu/gorm"
	"pingrobot-api.go/domain"
)

type WebSericeRepo struct {
	db *gorm.DB
}

func NewWebSericeRepo(db *gorm.DB) *WebSericeRepo {
	return &WebSericeRepo{db}
}

func (w *WebSericeRepo) CreateWebService(ctx context.Context, webService *domain.WebSerice) error {
	return w.db.Create(webService).Error
}

func (w *WebSericeRepo) GetWebServiceByUserId(ctx context.Context, userId uint) ([]domain.WebSerice, error){
	var webServices []domain.WebSerice
	err := w.db.Where("user_id = ?", userId).Find(*&webServices).Error
	if err != nil{
		return nil, err
	}

	return webServices, nil
}
