package httpHandler

import (
	"encoding/json"
	"github.com/calvinkmts/expert-pancake/engine/errorCode"
	"github.com/calvinkmts/expert-pancake/engine/errors"
	"io"
	"net/http"
)

func ParseHTTPRequest(r *http.Request, req interface{}) error {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		err = errors.NewServerError(errorCode.RequestReadErrorCode, err.Error())
		return err
	}

	err = json.Unmarshal(body, &req)
	if err != nil {
		err = errors.NewServerError(errorCode.RequestUnmarshalErrorCode, err.Error())
		return err
	}

	return nil
}

func WriteResponse(w http.ResponseWriter, res interface{}) error {

	body, err := json.Marshal(generalResponse{
		Result: res,
	})
	if err != nil {
		err = errors.NewServerError(errorCode.ResponseMarshalErrorCode, err.Error())
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(body)

	if err != nil {
		err = errors.NewServerError(errorCode.ResponseWriteErrorCode, err.Error())
		return err
	}

	return nil
}
