package validator

import "github.com/gin-gonic/gin"

func New[T any](ctx *gin.Context) *RequestValidator[T] {
	return &RequestValidator[T]{
		ctx: ctx,
	}
}
