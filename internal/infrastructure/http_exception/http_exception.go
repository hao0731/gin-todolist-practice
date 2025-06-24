package http_exception

import "net/http"

type HttpException struct {
	StatusCode int                 `json:"statusCode"`
	Message    string              `json:"message"`
	Detail     map[string][]string `json:"detail"`
}

func New(code int, message string, detail map[string][]string) HttpException {
	return HttpException{
		StatusCode: code,
		Message:    message,
		Detail:     detail,
	}
}

func NewBadRequestException(message string, detail map[string][]string) HttpException {
	return New(http.StatusBadRequest, message, detail)
}

func NewInternalServerErrorException(message string, detail map[string][]string) HttpException {
	return New(http.StatusInternalServerError, message, detail)
}
