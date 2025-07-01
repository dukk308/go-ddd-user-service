package common

import "time"

type BaseSQL struct {
	ID        string    `gorm:"column:id;primaryKey;type:uuid;default:gen_random_uuid()"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
