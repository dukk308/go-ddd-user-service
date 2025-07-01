package relational

import (
	"context"
	"user-service/internal/modules/user/domain"

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

func (r *UserRepository) CreateUser(ctx context.Context, user *domain.User) error {
	userSQL := ToEntity(*user)
	userSQL.ID = uuid.New().String()
	return r.db.Create(&userSQL).Error
}

func (r *UserRepository) GetListUser(ctx context.Context) ([]domain.User, error) {
	var usersSQL []UserSQL
	if err := r.db.Find(&usersSQL).Error; err != nil {
		return nil, err
	}

	users := make([]domain.User, len(usersSQL))
	for i, userSQL := range usersSQL {
		users[i] = ToDomain(userSQL)
	}

	return users, nil
}

func (r *UserRepository) GetUserByID(ctx context.Context, id string) domain.User {
	var userSQL UserSQL

	if err := r.db.Where("id = ?", id).First(&userSQL).Error; err != nil {
		return domain.User{}
	}

	return ToDomain(userSQL)
}
