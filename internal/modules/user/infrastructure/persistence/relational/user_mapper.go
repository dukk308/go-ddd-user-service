package relational

import (
	"user-service/internal/modules/user/domain"
)

func ToDomain(userSQL UserSQL) domain.User {
	user := domain.User{}

	user.ID = userSQL.ID
	user.Email = userSQL.Email
	user.Phone = userSQL.Phone
	user.Username = userSQL.Username
	user.Password = userSQL.Password
	user.CreatedAt = userSQL.CreatedAt
	user.UpdatedAt = userSQL.UpdatedAt

	return user
}

func ToEntity(user domain.User) UserSQL {
	userSQL := UserSQL{}

	userSQL.Email = user.Email
	userSQL.Phone = user.Phone
	userSQL.Username = user.Username
	userSQL.Password = user.Password

	return userSQL
}
