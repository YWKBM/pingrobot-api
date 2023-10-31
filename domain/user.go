package domain

import (
	"errors"
	"net/mail"
	"time"
)

type User struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	RegisteredAt time.Time `json:"registered_at"`
	LastVisitAt  time.Time `json:"last_visit_at"`
}

func (u User) Validate() error {
	if u.Name == "" {
		return errors.New("Invalid username")
	}

	if len(u.Password) < 6 {
		return errors.New("Password should be at least 6 characters")
	}

	if !validEmail(u.Email) {
		return errors.New("Invalid email")
	}

	return nil
}

func validEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
