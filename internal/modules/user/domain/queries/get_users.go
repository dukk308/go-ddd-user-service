package queries

import (
	"context"
	"user-service/internal/modules/user/domain"
)

type QueryUsersCommand interface {
	Execute(ctx context.Context) ([]domain.ExposeUser, error)
}

type QueryUser interface {
	GetListUser(ctx context.Context) ([]domain.ExposeUser, error)
}

type queryUsers struct {
	userRepository domain.UserRepositoryPort
}

func NewQueryUsers(userRepository domain.UserRepositoryPort) QueryUsersCommand {
	return &queryUsers{
		userRepository: userRepository,
	}
}

func (c *queryUsers) Execute(ctx context.Context) ([]domain.ExposeUser, error) {
	users, err := c.userRepository.GetListUser(ctx)
	if err != nil {
		return nil, err
	}

	publicUsers := make([]domain.ExposeUser, len(users))

	for i, user := range users {
		dto := domain.ExposeUser{}
		publicUsers[i] = *dto.From(&user)
	}

	return publicUsers, nil
}
