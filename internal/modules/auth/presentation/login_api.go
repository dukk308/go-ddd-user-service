package presentation

import (
	"github.com/gin-gonic/gin"
)

type LoginApi interface {
	Handle(ctx *gin.Engine)
}

type loginApi struct {
	loginHandler LoginHandler
}

func NewLoginApi(loginHandler LoginHandler) *loginApi {
	return &loginApi{loginHandler: loginHandler}
}

func (l *loginApi) Handle(ctx *gin.Engine) {

}
