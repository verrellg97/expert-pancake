package errors

import "fmt"

var ServerErrorType = ServerError{}

type ServerError struct {
	Code    int64
	Message string
}

func NewServerError(code int64, message string) error {
	return ServerError{
		Code:    code,
		Message: message,
	}
}

func (ths ServerError) Error() string {
	return fmt.Sprintf("[%v] %v", ths.Code, ths.Message)
}
