package relational

import (
	"user-service/internal/shared/common"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserSQL struct {
	common.BaseSQL
	Username string `gorm:"column:username"`
	Email    string `gorm:"column:email"`
	Password string `gorm:"column:password"`
	Phone    string `gorm:"column:phone"`
	Role     string `gorm:"column:role"`
}

func (UserSQL) TableName() string {
	return "users"
}

func (u *UserSQL) BeforeCreate(tx *gorm.DB) (err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	return
}

func (u *UserSQL) Migrate(tx *gorm.DB) (err error) {
	if err := tx.AutoMigrate(&UserSQL{}); err != nil {
		return err
	}

	return nil
}
