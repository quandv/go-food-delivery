package middleware

import (
	common "go-food-delivery/common"
	appctx "go-food-delivery/component/app-context"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
* Recover all app crash errors
 */
func Recover(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				err, ok := r.(*common.AppError)
				if ok {
					ctx.AbortWithStatusJSON(err.StatusCode, err)
					panic(err)
					return
				}

				appErr := r.(error) // ép recover về kiểu error
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, common.InternalServerError(appErr))
				panic(appErr)
				return
			}
		}()

		ctx.Next()
	}
}
