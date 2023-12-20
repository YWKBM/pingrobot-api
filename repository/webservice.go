package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"pingrobot-api.go/domain"
)

type WebSericeRepo struct {
	db *sql.DB
}

func NewWebSericeRepo(db *sql.DB) *WebSericeRepo {
	return &WebSericeRepo{db}
}

func (w *WebSericeRepo) Create(userId int, webService domain.WebService) (int, error) {
	var webServiceId int
	tx, err := w.db.Begin()
	if err != nil {
		tx.Rollback()
	}

	var exists bool
	row := tx.QueryRow("SELECT EXISTS(SELECT user_id, link, port FROM web_services WHERE user_id = $1 AND link = $2 AND port = $3)", userId, webService.Link, webService.Port)
	if err := row.Scan(&exists); err != nil {
		tx.Rollback()
		return 0, err
	}

	if exists {
		err = errors.New("Already exists")
		return 0, err
	}

	var userEmail string
	row = tx.QueryRow("SELECT email FROM users WHERE id = $1", userId)
	if err := row.Scan(&userEmail); err != nil {
		tx.Rollback()
		return 0, err
	}

	createWebServiceQuery := fmt.Sprintf("INSERT INTO web_services (user_id, user_email, name, link, port) values ($1, $2, $3, $4, $5) RETURNING id")
	rows, err := tx.Query(createWebServiceQuery, userId, userEmail, webService.Name, webService.Link, webService.Port)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	rows.Next()
	{
		rows.Scan(&webServiceId)
	}

	return webServiceId, tx.Commit()
}

func (w *WebSericeRepo) GetAll(userId int) ([]domain.WebService, error) {
	var webServices []domain.WebService

	rows, err := w.db.Query("SELECT * FROM web_services WHERE user_id = $1", userId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var webService domain.WebService
		err := rows.Scan(&webService.ID, &webService.UserID, &webService.UserEmail, &webService.Name, &webService.Link, &webService.Port, &webService.Status, &webService.Alarm)
		if err != nil {
			return nil, err
		}
		webServices = append(webServices, webService)
	}
	return webServices, nil

}

func (w *WebSericeRepo) GetById(userId int, webServiceId int) (domain.WebService, error) {
	var webService domain.WebService

	err := w.db.QueryRow("SELECT id, user_id, user_email, name, link, port, status FROM web_services WHERE user_id = $1 AND id = $2", userId, webServiceId).
		Scan(&webService.ID, &webService.UserID, &webService.UserEmail, &webService.Name, &webService.Link, &webService.Port, &webService.Status, &webService.Alarm)

	return webService, err
}

func (w *WebSericeRepo) Delete(userId int, webServiceId int) error {
	_, err := w.db.Query("DELETE FROM web_services WHERE user_id = $1 AND id = $2", userId, webServiceId)

	return err
}

func (w *WebSericeRepo) Update(userId, webServiceId int, input domain.UpdateWebServiceInput) error {
	_, err := w.db.Query("UPDATE web_services SET name = $1, link = $2, port = $3 WHERE user_id = $4 AND id = $5",
		input.Name, input.Link, input.Port, userId, webServiceId)

	return err
}
