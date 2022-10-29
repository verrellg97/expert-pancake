package httpHandler

import (
	"encoding/json"
	"fmt"
	"github.com/calvinkmts/expert-pancake/engine/errors"
	"log"
	"net/http"
)

type HttpHandlerFunc func(w http.ResponseWriter, r *http.Request) error

type HttpHandler struct {
	handlerFunc HttpHandlerFunc
}

type dataError struct {
	Field     string `json:"field"`
	ErrorType string `json:"error_type"`
}

type customErrorResponse struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Errors  []dataError `json:"errors"`
}

type generalResponse struct {
	Result interface{}          `json:"result,omitempty"`
	Error  *customErrorResponse `json:"error,omitempty"`
}

func New(handlerFunc HttpHandlerFunc) HttpHandler {
	return HttpHandler{handlerFunc: handlerFunc}
}

func (ths HttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := ths.handlerFunc(w, r)
	if err != nil {
		errEncode := ths.encodeJsonErrorResponse(w, err)
		if errEncode != nil {
			log.Println("ServeHTTP", errEncode)
		}
		log.Println("ServeHTTP", err)
		return
	}
}

func (ths HttpHandler) encodeJsonErrorResponse(w http.ResponseWriter, err error) error {
	var response generalResponse
	statusCode := http.StatusInternalServerError

	switch errType := err.(type) {
	case *errors.ClientError:
		statusCode = http.StatusBadRequest
		var clientDataError []dataError
		for _, data := range errType.Data {
			clientDataError = append(clientDataError, dataError{
				Field:     data.Field,
				ErrorType: data.ErrorType,
			})
		}

		response = generalResponse{
			Error: &customErrorResponse{
				Code:    errType.Code,
				Message: errType.Message,
				Errors:  clientDataError,
			},
		}
	case errors.ServerError:
		statusCode = http.StatusInternalServerError
		response = generalResponse{
			Error: &customErrorResponse{
				Code:    errType.Code,
				Message: "Internal Server Error",
				Errors:  []dataError{},
			},
		}
	default:
		statusCode = http.StatusInternalServerError
		response = generalResponse{
			Error: &customErrorResponse{
				Code:    errors.UnknownErrorTypeErrorCode,
				Message: "Internal Server Error",
				Errors:  []dataError{},
			},
		}
	}

	responseByte, err := json.Marshal(response)
	if err != nil {
		return fmt.Errorf("[EncodeJsonErrorResponse] error when marshal response with message: %v", err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	_, err = w.Write(responseByte)
	if err != nil {
		return fmt.Errorf("[EncodeJsonErrorResponse] error when write response with message: %v", err.Error())
	}

	return nil
}
