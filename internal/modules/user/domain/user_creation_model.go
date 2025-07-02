package domain

import "errors"

type UserCreation struct {
	Username string  `json:"username" binding:"required"`
	Password string  `json:"password" binding:"required"`
	Email    string  `json:"email" binding:"required,email"`
	Phone    *string `json:"phone,omitempty"`
}

func (u *UserCreation) Validate() error {
	if u.Username == "" {
		return errors.New("username is required")
	}

	if u.Email == "" {
		return errors.New("email is required")
	}

	return nil
}
