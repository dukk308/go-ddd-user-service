package domain

import (
	"user-service/internal/shared/common"
)

type User struct {
	common.BaseModel
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Password string  `json:"password,omitempty"`
	Phone    *string `json:"phone,omitempty"`
	Role     string  `json:"role,omitempty"`
}

func (u *User) From(ur *UserCreation) *User {
	return &User{
		Username: ur.Username,
		Email:    ur.Email,
		Password: ur.Password,
		Phone:    ur.Phone,
		Role:     "user",
	}
}
