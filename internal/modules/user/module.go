package user

import (
	"net/http"
	"user-service/internal/modules/user/domain/commands"
	"user-service/internal/modules/user/domain/queries"
	"user-service/internal/modules/user/infrastructure/persistence/relational"
	"user-service/internal/modules/user/presentation"
	"user-service/internal/shared/common"
	loggerPgk "user-service/pkg/logger"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserModule(router *gin.Engine, db *gorm.DB) {
	repository := relational.NewUserRepository(db)
	logger, _, _ := loggerPgk.InitializeLogger("info")
	// auth
	{
		authRouter := router.Group("v1/auth")

		createUserCommand := commands.NewCreateUserHandler(repository, logger)
		signupApi := presentation.NewSignupApi(createUserCommand)

		authRouter.POST("/signup", signupApi.Handle)
		authRouter.POST("/signin", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, common.NewApiResponse(nil))
		})
	}

	// user
	{
		userRouter := router.Group("v1/users")

		queryUsersHandler := queries.NewQueryUsers(repository)
		getUsersApi := presentation.NewGetUsersApi(queryUsersHandler)
		queryUserByIdHandler := queries.NewQueryUserById(repository)
		getUserByIdApi := presentation.NewGetUserByIdApi(queryUserByIdHandler)

		userRouter.GET("/", getUsersApi.Handle)
		userRouter.GET("/:id", getUserByIdApi.Handle)
	}

}
