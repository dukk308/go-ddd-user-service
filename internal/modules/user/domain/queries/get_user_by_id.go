package queries

import (
	"context"
	"errors"
	"user-service/internal/modules/user/domain"
)

type QueryUserById interface {
	Execute(ctx context.Context, id string) (domain.ExposeUser, error)
}

type queryUserById struct {
	userRepository domain.UserRepositoryPort
}

func NewQueryUserById(userRepository domain.UserRepositoryPort) *queryUserById {
	return &queryUserById{
		userRepository: userRepository,
	}
}

func (c *queryUserById) Execute(ctx context.Context, id string) (domain.ExposeUser, error) {
	user := c.userRepository.GetUserByID(ctx, id)
	if user.ID == "" {
		return domain.ExposeUser{}, errors.New("user not found")
	}

	dto := domain.ExposeUser{}
	return *dto.From(&user), nil
}
