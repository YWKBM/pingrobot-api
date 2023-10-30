package domain

import (
	"time"
)

type User struct {
	ID           int
	Name         string
	Email        string
	Password     string
	RegisteredAt time.Time
	LastVisitAt  time.Time
}
