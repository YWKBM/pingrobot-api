package transport

import (
	"context"

	"pingrobot-api.go/domain"
)

type User interface {
	Create(ctx context.Context, user domain.User) error
	GetUserById(ctx context.Context, id int64) (domain.User, error)
}


