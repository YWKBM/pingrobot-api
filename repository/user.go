package repository

import (
	"database/sql"
	"context"
	"time"

	"pingrobot-api.go/domain"
)

type UserRepo struct {
	db *sql.DB
}

func NewUsersRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db}
}

func (u *UserRepo) Create(ctx context.Context, user domain.User) error {
	_, err := u.db.Exec("INSERT INTO users (name, email, password, registered_at, last_visit_at) values ($1, $2, $3, $4)",
	user.Name, user.Email, user.Password, time.Now, time.Now)
	
	return err
}

func (u *UserRepo) GetUser(ctx context.Context, email string, password string) (domain.User, error) {
	var user domain.User

	err := u.db.QueryRow("SELECT id, name, email, password, registered_at, last_visit_at WHERE email = $1 AND password = $2", email, password).
		Scan(&user.ID, user.Name, user.Email, user.Password, user.RegisteredAt, user.LastVisitAt)

	if err != nil{
		return user, err
	}

	user.WebServices, err = u.findWebServiceByUID(user.ID)
	if err != nil{
		return user, err
	}

	return user, nil 
}

func (u *UserRepo) findWebServiceByUID(UID int) ([]domain.WebService, error){
	var webServices []domain.WebService
	rows, err := u.db.Query("SELECT * FROM web_services WHRE user_id = $1", UID)
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

func (u *UserRepo) CreateWebService(ctx context.Context, webService domain.WebService) error {
	_, err := u.db.Query("SELECT * FROM users WHERE user_id = $1", webService.UserID)
	if err != nil{
		return err
	}

	_, err = u.db.Exec("INSERT INTO web_services (user_id, name, link, port, status) values ($1, $2, $3, $4, $5", 
	webService.UserID, webService.Name, webService.Link, webService.Port, webService.Status)

	return err
}
