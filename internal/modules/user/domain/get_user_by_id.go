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

func (c *queryUserById) Execute(ctx context.Context, id string) (UserExposed, error) {
	user := c.userRepository.GetUserByID(ctx, id)
	userExposed := UserExposed{}

	if user == nil {
		return userExposed, errors.New("user not found")
	}

	return userExposed.From(user), nil
}
