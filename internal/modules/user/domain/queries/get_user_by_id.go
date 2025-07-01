package queries

import (
	"context"
	"errors"
	"user-service/internal/modules/user/domain"
)

type QueryUserById interface {
	Execute(ctx context.Context, id string) (domain.UserExposed, error)
}

type queryUserById struct {
	userRepository domain.UserRepositoryPort
}

func NewQueryUserById(userRepository domain.UserRepositoryPort) *queryUserById {
	return &queryUserById{
		userRepository: userRepository,
	}
}

func (c *queryUserById) Execute(ctx context.Context, id string) (domain.UserExposed, error) {
	user := c.userRepository.GetUserByID(ctx, id)
	userExposed := &domain.UserExposed{}

	if user == nil {
		return *userExposed, errors.New("user not found")
	}

	dto := domain.UserExposed{}
	return *dto.From(user), nil
}
