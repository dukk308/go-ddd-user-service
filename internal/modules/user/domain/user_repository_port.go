package domain

import (
	"context"
	"user-service/internal/shared/common"
)

type UserRepositoryPort interface {
	CreateUser(ctx context.Context, user *User) *common.AppError
	GetListUser(ctx context.Context) ([]*User, *common.AppError)
	GetUserByID(ctx context.Context, id string) (*User, *common.AppError)
	GetUserByEmail(ctx context.Context, email string) (*User, *common.AppError)
	GetUserByUsername(ctx context.Context, username string) (*User, *common.AppError)
}
