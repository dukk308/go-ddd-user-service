package relational

import (
	"context"
	"errors"
	"user-service/internal/modules/user/domain"
	biz_errors "user-service/internal/modules/user/domain/errors"
	"user-service/internal/shared/common"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var _ domain.UserRepositoryPort = (*UserRepository)(nil)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *domain.User) *common.AppError {
	userSQL := ToEntity(*user)
	userSQL.ID = uuid.New().String()

	if err := r.db.Create(&userSQL).Error; err != nil {
		return common.DbUnexpectedError(err, "CREATE_USER_FAILED", "Create user failed")
	}

	return nil
}

func (r *UserRepository) GetListUser(ctx context.Context) ([]*domain.User, *common.AppError) {
	var usersSQL []UserSQL
	if err := r.db.Find(&usersSQL).Error; err != nil {
		return nil, common.DbUnexpectedError(err, "GET_LIST_USER_FAILED", "Get list user failed")
	}

	users := make([]*domain.User, len(usersSQL))

	for i, userSQL := range usersSQL {
		users[i] = ToDomain(userSQL)
	}

	return users, nil
}

func (r *UserRepository) GetUserByID(ctx context.Context, id string) (*domain.User, *common.AppError) {
	var userSQL UserSQL

	if err := r.db.Where("id = ?", id).First(&userSQL).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.DbRecordNotFoundError(err)
		}

		return nil, common.DbUnexpectedError(err, "GET_USER_BY_ID_FAILED", "Get user is failed")
	}

	return ToDomain(userSQL), nil
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, *common.AppError) {
	var userSQL UserSQL

	if err := r.db.Where("email = ?", email).First(&userSQL).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.DbRecordNotFoundError(err)
		}

		return nil, common.DbUnexpectedError(err, "GET_USER_BY_EMAIL_FAILED", "Get user is failed")
	}

	return ToDomain(userSQL), nil
}

func (r *UserRepository) GetUserByUsername(ctx context.Context, username string) (*domain.User, *common.AppError) {
	var userSQL UserSQL

	if err := r.db.Where("username = ?", username).First(&userSQL).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, biz_errors.UserNotFoundError(err)
		}

		return nil, common.DbUnexpectedError(err, "GET_USER_BY_USERNAME_FAILED", "Get user is failed")
	}

	return ToDomain(userSQL), nil
}
