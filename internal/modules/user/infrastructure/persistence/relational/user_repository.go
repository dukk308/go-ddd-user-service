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

func (r *UserRepository) GetListUser(ctx context.Context) []*domain.User {
	var usersSQL []UserSQL
	if err := r.db.Find(&usersSQL).Error; err != nil {
		return nil
	}

	users := make([]*domain.User, len(usersSQL))
	for i, userSQL := range usersSQL {
		users[i] = ToDomain(userSQL)
	}

	return users
}

func (r *UserRepository) GetUserByID(ctx context.Context, id string) *domain.User {
	var userSQL UserSQL

	if err := r.db.Where("id = ?", id).First(&userSQL).Error; err != nil {
		return nil
	}

	return ToDomain(userSQL)
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) *domain.User {
	var userSQL UserSQL

	if err := r.db.Where("email = ?", email).First(&userSQL).Error; err != nil {
		return nil
	}

	return ToDomain(userSQL)
}

func (r *UserRepository) GetUserByUsername(ctx context.Context, username string) *domain.User {
	var userSQL UserSQL

	if err := r.db.Where("username = ?", username).First(&userSQL).Error; err != nil {
		return nil
	}

	return ToDomain(userSQL)
}
