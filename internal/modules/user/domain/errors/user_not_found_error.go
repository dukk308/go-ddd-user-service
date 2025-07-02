package biz_errors

import (
	"net/http"
	"user-service/internal/shared/common"
)

func UserNotFoundError(err error) *common.AppError {
	return common.NewAppError(err, "USER_NOT_FOUND", "User not found", http.StatusNotFound, "")
}
