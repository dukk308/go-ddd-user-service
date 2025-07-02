package commands

import (
	"context"
	"net/http"
	"user-service/internal/modules/user/domain"
	biz_errors "user-service/internal/modules/user/domain/errors"
	"user-service/internal/shared/common"

	"go.uber.org/zap"
)

type CreateUserHandler interface {
	Execute(ctx context.Context, user *domain.UserCreation) *common.AppError
}

type createUserHandler struct {
	repository domain.UserRepositoryPort
	logger     *zap.Logger
}

func NewCreateUserHandler(repository domain.UserRepositoryPort, logger *zap.Logger) *createUserHandler {
	return &createUserHandler{repository: repository, logger: logger}
}

func (s *createUserHandler) Execute(ctx context.Context, user *domain.UserCreation) *common.AppError {
	if err := user.Validate(); err != nil {
		s.logger.Error("Validate user", zap.Error(err))
		return common.NewAppError(err, "VALIDATE_USER_FAILED", "Validate user failed", http.StatusBadRequest, "")
	}

	if user, err := s.repository.GetUserByUsername(ctx, user.Username); user != nil {
		s.logger.Error("Username already exists", zap.String("username", user.Username))
		return biz_errors.UsernameIsExistedError(err)
	}

	userBiz := &domain.User{}
	if err := s.repository.CreateUser(ctx, userBiz.From(user)); err != nil {
		s.logger.Error("Create user Failed", zap.Error(err))
		return err
	}

	return nil
}
