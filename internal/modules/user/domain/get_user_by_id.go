package domain

import (
	"context"
	"errors"
)

type QueryUserById interface {
	Execute(ctx context.Context, id string) User
}

type queryUserById struct {
	userRepository UserRepositoryPort
}

func NewQueryUserById(userRepository UserRepositoryPort) *queryUserById {
	return &queryUserById{
		userRepository: userRepository,
	}
}

func (c *queryUserById) Execute(ctx context.Context, id string) (ExposeUser, error) {
	user := c.userRepository.GetUserByID(ctx, id)
	if user.ID == "" {
		return ExposeUser{}, errors.New("user not found")
	}

	dto := ExposeUser{}
	return *dto.From(&user), nil
}
