package domain

import (
	"time"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	ID           int64       `json:"id"`
	Name         string      `json:"name"`
	Email        string      `json:"email"`
	Password     string      `json:"password"`
	RegisteredAt time.Time   `json:"registeredAt"`
	LastVisitAt  time.Time   `json:"lastVisitAt"`
	Webservices  []WebSerice `json:"webservices"`
	//Verification
}
