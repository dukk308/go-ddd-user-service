package presentation

import (
	"net/http"
	"user-service/internal/modules/user/domain/queries"
	"user-service/internal/shared/common"

	"github.com/gin-gonic/gin"
)

type GetUserByIdApi interface {
	Handle(ctx *gin.Context)
}

type getUserByIdApi struct {
	GetUserByIdCommand queries.QueryUserById
}

func NewGetUserByIdApi(getUserByIdCommand queries.QueryUserById) *getUserByIdApi {
	return &getUserByIdApi{
		GetUserByIdCommand: getUserByIdCommand,
	}
}

func (api *getUserByIdApi) Handle(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := api.GetUserByIdCommand.Execute(ctx, id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, common.NewApiResponseError(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, common.NewApiResponse(user))
}
