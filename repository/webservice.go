package repository

import (
	"context"

	"database/sql"
	"pingrobot-api.go/domain"
)

type WebSericeRepo struct {
	db *sql.DB
}

func NewWebSericeRepo(db *sql.DB) *WebSericeRepo {
	return &WebSericeRepo{db}
}

func (w *WebSericeRepo) GetAllWebServices(ctx context.Context) ([]domain.WebService, error) {
	var webServices []domain.WebService

	rows, err := w.db.Query("SELECT * FROM web_services")
	if err != nil{
		return nil, err
	}

	for rows.Next(){
		var webService domain.WebService
		err := rows.Scan(&webService.ID, &webService.UserID, &webService.Name, &webService.Link, &webService.Port, &webService.Status)
		if err != nil{
			return nil, err
		}
		webServices = append(webServices, webService)
	}
	return webServices, nil

}
