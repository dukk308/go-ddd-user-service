package queries

import (
	"context"
	"user-service/internal/modules/user/domain"
)

type QueryUsersCommand interface {
	Execute(ctx context.Context) ([]domain.UserExposed, error)
}

type QueryUser interface {
	GetListUser(ctx context.Context) ([]domain.UserExposed, error)
}

type queryUsers struct {
	userRepository domain.UserRepositoryPort
}

func NewQueryUsers(userRepository domain.UserRepositoryPort) QueryUsersCommand {
	return &queryUsers{
		userRepository: userRepository,
	}
}

func (c *queryUsers) Execute(ctx context.Context) ([]domain.UserExposed, error) {
	users := c.userRepository.GetListUser(ctx)
	publicUsers := make([]domain.UserExposed, len(users))

	for i, user := range users {
		dto := domain.UserExposed{}
		publicUsers[i] = dto.From(user)
	}

	return publicUsers, nil
}
