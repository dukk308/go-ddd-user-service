package domain

import (
	"context"
)

type UserRepositoryPort interface {
	CreateUser(ctx context.Context, user *User) error
	GetListUser(ctx context.Context) ([]User, error)
	GetUserByID(ctx context.Context, id string) User
}
