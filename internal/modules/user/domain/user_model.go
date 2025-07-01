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
