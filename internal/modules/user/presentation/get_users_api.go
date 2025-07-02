package presentation

import (
	"net/http"
	"user-service/internal/modules/user/domain/queries"
	"user-service/internal/shared/common"

	"github.com/gin-gonic/gin"
)

type GetUsersApi interface {
	Handle(ctx *gin.Context)
}

type getUsersApi struct {
	GetUserCommand queries.QueryUsersCommand
}

func NewGetUsersApi(getUserCommand queries.QueryUsersCommand) *getUsersApi {
	return &getUsersApi{
		GetUserCommand: getUserCommand,
	}
}

func (api *getUsersApi) Handle(ctx *gin.Context) {
	users, err := api.GetUserCommand.Execute(ctx)

	if err != nil {
		ctx.JSON(err.StatusCode, err)
		return
	}

	ctx.JSON(http.StatusOK, common.NewApiResponse(users))
}
