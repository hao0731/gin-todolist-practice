package validator

import (
	"example.com/todolist/internal/infrastructure/http_exception"
	z "github.com/Oudwins/zog"
	zhttp "github.com/Oudwins/zog/zhttp"
	"github.com/gin-gonic/gin"
)

func ZogValidatePayload[T any](schema *z.StructSchema) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var payload T
		if errs := schema.Parse(zhttp.Request(ctx.Request), &payload); errs != nil {
			sanitized := z.Issues.SanitizeMap(errs)
			exception := http_exception.NewBadRequestException("Invalid request payload", sanitized)
			ctx.AbortWithStatusJSON(exception.StatusCode, exception)
			return
		}

		ctx.Set(VALIDATED_PAYLOAD_KEY, &payload)
		ctx.Next()
	}
}
