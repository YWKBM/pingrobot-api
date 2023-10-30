package repository

import (
	"database/sql"
	"time"

	"pingrobot-api.go/domain"
)

type Auth struct {
	db *sql.DB
}

func NewAuthorizationRepo(db *sql.DB) *Auth {
	return &Auth{db}
}

func (a *Auth) CreateUser(user domain.User) (int, error) {
	_, err := a.db.Exec("INSERT INTO users (name, email, password, registered_at, last_visit_at) values ($1, $2, $3, $4, $5)",
		user.Name, user.Email, user.Password, time.Now(), time.Now())

	if err != nil {
		return 0, err
	}
	return user.ID, err
}

func (a *Auth) GetUser(name, password string) (domain.User, error) {
	var user domain.User

	err := a.db.QueryRow("SELECT id, name, email, password, registered_at, last_visit_at FROM users WHERE name = $1 AND password = $2", name, password).
		Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.RegisteredAt, &user.LastVisitAt)

	if err == sql.ErrNoRows {
		return user, err
	}

	return user, nil
}
