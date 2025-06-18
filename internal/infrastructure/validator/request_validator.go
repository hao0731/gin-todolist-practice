package validator

import (
	"errors"

	"github.com/gin-gonic/gin"

	v "github.com/go-playground/validator/v10"
)

type RequestValidator[T any] struct {
	ctx *gin.Context
}

func NewRequestValidator[T any](ctx *gin.Context) *RequestValidator[T] {
	return &RequestValidator[T]{
		ctx: ctx,
	}
}

func (validator *RequestValidator[T]) Validate(payload *T) map[string]string {
	if err := validator.ctx.ShouldBindJSON(payload); err != nil {
		var validationErrors v.ValidationErrors
		if errors.As(err, &validationErrors) {
			out := make(map[string]string)
			for _, fieldError := range validationErrors {
				out[fieldError.Field()] = fieldError.Error()
			}
			return out
		}
		return map[string]string{"error": err.Error()}
	}
	return nil
}
