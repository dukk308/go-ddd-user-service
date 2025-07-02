package queries

import (
	"context"
	"user-service/internal/modules/user/domain"
	"user-service/internal/shared/common"
)

type QueryUsersCommand interface {
	Execute(ctx context.Context) ([]domain.UserExposed, *common.AppError)
}

type QueryUser interface {
	GetListUser(ctx context.Context) ([]domain.UserExposed, *common.AppError)
}

type queryUsers struct {
	userRepository domain.UserRepositoryPort
}

func NewQueryUsers(userRepository domain.UserRepositoryPort) QueryUsersCommand {
	return &queryUsers{
		userRepository: userRepository,
	}
}

func (c *queryUsers) Execute(ctx context.Context) ([]domain.UserExposed, *common.AppError) {
	users, err := c.userRepository.GetListUser(ctx)
	if err != nil {
		return nil, err
	}

	publicUsers := make([]domain.UserExposed, len(users))

	for i, user := range users {
		dto := domain.UserExposed{}
		publicUsers[i] = *dto.From(user)
	}

	return publicUsers, nil
}
