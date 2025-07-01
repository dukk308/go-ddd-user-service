package main

import (
	"net/http"
	"user-service/internal/shared/libs"

	"user-service/internal/modules/user"

	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.Default()
	db := libs.NewDB("postgres://postgres:123123@localhost:4633/user_service?sslmode=disable")

	// Routes
	user.UserModule(e, db)
	e.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	e.Run(":3000")
}
