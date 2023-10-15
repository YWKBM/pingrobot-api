package domain

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	ID           int64       `json:"id"`
	Name         string      `json:"name" gorm:"not null;size:255"`
	Email        string      `json:"email" gorm:"not null;size:255"`
	Password     string      `json:"password" gorm:"not null;size:255"`
	RegisteredAt time.Time   `json:"registeredAt"`
	LastVisitAt  time.Time   `json:"lastVisitAt"`
	Webservices  []WebSerice `json:"webservices"`
	//Verification
}
