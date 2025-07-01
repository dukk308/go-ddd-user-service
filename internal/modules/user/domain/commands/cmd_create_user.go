package commands

import (
	"context"
	"errors"
	"user-service/internal/modules/user/domain"

	"go.uber.org/zap"
)

type CreateUserCommand interface {
	Execute(ctx context.Context, user *domain.UserCreateRequest) error
}

type createUserCommand struct {
	repository domain.UserRepositoryPort
	logger     *zap.Logger
}

func NewCreateUserCommand(repository domain.UserRepositoryPort, logger *zap.Logger) *createUserCommand {
	return &createUserCommand{repository: repository, logger: logger}
}

func (s *createUserCommand) Execute(ctx context.Context, user *domain.UserCreateRequest) error {
	if err := user.Validate(); err != nil {
		s.logger.Error("Validate user", zap.Error(err))
		return err
	}

	if user := s.repository.GetUserByUsername(ctx, user.Username); user != nil {
		s.logger.Error("Username already exists", zap.String("username", user.Username))
		return errors.New("username already exists")
	}

	userDomain := &domain.User{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
		Phone:    user.Phone,
		Role:     "user",
	}

	if err := s.repository.CreateUser(ctx, userDomain); err != nil {
		s.logger.Error("Create user Failed", zap.Error(err))
		return err
	}

	return nil
}
