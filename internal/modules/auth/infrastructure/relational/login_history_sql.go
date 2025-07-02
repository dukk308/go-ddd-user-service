package relational

import "user-service/internal/shared/common"

type LoginHistorySQL struct {
	common.BaseSQL
	UserID    string `gorm:"column:user_id"`
	IPAddress string `gorm:"column:ip_address"`
}
