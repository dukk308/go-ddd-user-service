package biz_errors

import (
	"net/http"
	"user-service/internal/shared/common"
)

func UsernameIsExistedError(err error) *common.AppError {
	return common.NewAppError(err, "USERNAME_IS_EXISTED", "Username is existed", http.StatusBadRequest, "")
}
