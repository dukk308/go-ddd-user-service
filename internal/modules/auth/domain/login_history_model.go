package domain

import "user-service/internal/shared/common"

type LoginHistory struct {
	common.BaseModel
	UserID    string `json:"user_id"`
	IPAddress string `json:"ip_address"`
}
