package components

import (
	"fmt"

	"user-service/internal/modules/user/infrastructure/persistence/relational"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	userSQL := relational.UserSQL{}
	userSQL.Migrate(db)

	fmt.Println("DB connected")

	return db
}
