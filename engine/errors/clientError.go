package errors

import "fmt"

var ClientErrorType = ClientError{}

type ClientErrorData struct {
	Field     string
	ErrorType string
}

type ClientError struct {
	Code    int64
	Message string
	Data    []ClientErrorData
}

func NewClientError() *ClientError {
	return &ClientError{
		Code:    32602,
		Message: "Invalid params",
	}
}

func (ths ClientError) Error() string {
	return fmt.Sprintf("[%v] %v", ths.Code, ths.Message)
}

func (ths *ClientError) With(err ClientErrorData) *ClientError {
	ths.Data = append(ths.Data, err)
	return ths
}

func (ths *ClientError) WithDataMap(errorMap map[string]string) *ClientError {
	for key, value := range errorMap {
		ths.Data = append(ths.Data, ClientErrorData{key, value})
	}
	return ths
}
