package domain

import (
	"errors"
	"net/url"
)

type WebService struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	UserEmail string `json:"user_email"`
	Name      string `json:"name"`
	Link      string `json:"link"`
	Port      int    `json:"port"`
	Status    string `json:"status"`
}

type UpdateWebServiceInput struct {
	Name string `json:"name"`
	Link string `json:"link"`
	Port int    `json:"port"`
}

func (w WebService) Validate() error {
	if validURL(w.Link) {
		return errors.New("Invalid url")
	}

	return nil
}

func (w UpdateWebServiceInput) Validate() error {
	if validURL(w.Link) {
		return errors.New("Invalid url")
	}

	return nil
}

func validURL(link string) bool {
	//TODO: rewrite for ping-app
	_, err := url.ParseRequestURI(link)
	return err == nil
}
