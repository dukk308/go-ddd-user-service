package middleware

import (
	"net/http"
	"user-service/internal/shared/common"

	"github.com/gin-gonic/gin"
)

func JsonRecoverMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				if appErr, ok := err.(*common.AppError); ok {
					ctx.AbortWithStatusJSON(appErr.StatusCode, appErr)
					panic(err)
				}

				ctx.AbortWithStatusJSON(http.StatusInternalServerError, common.NewAppError(nil, "INTERNAL_SERVER_ERROR", "Internal server error", http.StatusInternalServerError, ""))
				panic(err)
			}
		}()

		ctx.Next()
	}
}
