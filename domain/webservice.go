package domain

import (
	"github.com/jinzhu/gorm"
)

type WebSerice struct {
	gorm.Model
	UserId int64  `json:"id"`
	Name   string `json:"name"`
	Link   string `json:"link"`
	Port   string `json:"port"`
	Status string `jsong:"status"`
}
