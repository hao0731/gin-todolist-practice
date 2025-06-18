package error_handler

type HttpException struct {
	StatusCode int               `json:"status_code"`
	Message    string            `json:"message"`
	Detail     map[string]string `json:"detail,omitempty"`
}

func (exception HttpException) Error() string {
	return exception.Message
}

func NewHttpException(code int, message string, detail map[string]string) HttpException {
	return HttpException{
		StatusCode: code,
		Message:    message,
		Detail:     detail,
	}
}
