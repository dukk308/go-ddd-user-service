package commands

import (
	"context"
	"user-service/internal/modules/user/domain"
)

type CreateUserCommand interface {
	Execute(ctx context.Context, user *domain.UserCreateRequest) error
}

type createUserCommand struct {
	repository domain.UserRepositoryPort
}

func NewCreateUserCommand(repository domain.UserRepositoryPort) *createUserCommand {
	return &createUserCommand{repository: repository}
}

func (s *createUserCommand) Execute(ctx context.Context, user *domain.UserCreateRequest) error {
	if err := user.Validate(); err != nil {
		return err
	}

	userDomain := &domain.User{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
		Phone:    user.Phone,
		Role:     "user",
	}

	return s.repository.CreateUser(ctx, userDomain)
}
