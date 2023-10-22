package domain

type WebService struct {
	ID     int   `json:"id"`
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
	Link   string `json:"link"`
	Port   int    `json:"port"`
	Status string `json:"status"`
}
