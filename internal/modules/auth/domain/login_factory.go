package domain

import "user-service/internal/shared/common"

type LoginFactory struct {
	common.BaseModel
	UserID    string `json:"user_id"`
	IPAddress string `json:"ip_address"`
}
