package queries

import (
	"context"
	"user-service/internal/modules/user/domain"
	biz_errors "user-service/internal/modules/user/domain/errors"
	"user-service/internal/shared/common"
)

type QueryUserById interface {
	Execute(ctx context.Context, id string) (*domain.UserExposed, *common.AppError)
}

type queryUserById struct {
	userRepository domain.UserRepositoryPort
}

func NewQueryUserById(userRepository domain.UserRepositoryPort) *queryUserById {
	return &queryUserById{
		userRepository: userRepository,
	}
}

func (c *queryUserById) Execute(ctx context.Context, id string) (*domain.UserExposed, *common.AppError) {
	user, err := c.userRepository.GetUserByID(ctx, id)

	if err != nil {
		if err.Is(common.DbRecordNotFoundError(err)) {
			return nil, biz_errors.UserNotFoundError(err)
		}

		return nil, err
	}

	dto := domain.UserExposed{}
	return dto.From(user), nil
}
