package error_handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		if len(ctx.Errors) == 0 {
			return
		}

		err := ctx.Errors.Last().Err
		var httpException HttpException
		if !errors.As(err, &httpException) {
			httpException = NewHttpException(http.StatusInternalServerError, "Internal Server Error", map[string]string{"error": err.Error()})
		}
		ctx.JSON(httpException.StatusCode, httpException)
	}
}
