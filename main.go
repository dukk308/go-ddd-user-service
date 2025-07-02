package main

import (
	"net/http"
	middleware "user-service/internal/middlewares"
	"user-service/internal/modules/user"
	"user-service/internal/shared/components"

	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.Default()
	e.Use(middleware.JsonRecoverMiddleware())
	db := components.NewDB("postgres://postgres:123123@localhost:4633/user_service?sslmode=disable")

	// Modules Registration
	user.UserModule(e, db)

	e.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	e.Run(":3000")
}
