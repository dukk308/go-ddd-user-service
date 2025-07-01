package domain

import (
	"context"
)

type UserRepositoryPort interface {
	CreateUser(ctx context.Context, user *User) error
	GetListUser(ctx context.Context) []*User
	GetUserByID(ctx context.Context, id string) *User
	GetUserByEmail(ctx context.Context, email string) *User
	GetUserByUsername(ctx context.Context, username string) *User
}
