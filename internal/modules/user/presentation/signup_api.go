package presentation

import (
	"net/http"
	"user-service/internal/modules/user/domain"
	"user-service/internal/modules/user/domain/commands"
	"user-service/internal/shared/common"

	"github.com/gin-gonic/gin"
)

type SignupApi interface {
	Handle(ctx *gin.Context)
}

type signupApi struct {
	CreateUserCommand commands.CreateUserHandler
}

func NewSignupApi(createUserCommand commands.CreateUserHandler) SignupApi {
	return &signupApi{
		CreateUserCommand: createUserCommand,
	}
}

func (api *signupApi) Handle(ctx *gin.Context) {
	dto := &domain.UserCreation{}

	if err := ctx.ShouldBindJSON(dto); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	user := domain.UserCreation{
		Username: dto.Username,
		Password: dto.Password,
		Email:    dto.Email,
		Phone:    dto.Phone,
	}

	if err := api.CreateUserCommand.Execute(ctx, &user); err != nil {
		ctx.JSON(err.StatusCode, err)
		return
	}

	ctx.JSON(http.StatusOK, common.NewApiResponse(nil))
}
