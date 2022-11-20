package impl

import (
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
)

type JsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

type HelloWorldRequest struct {
}

func (a businessService) HelloWorld(w http.ResponseWriter, r *http.Request) error {

	var req HelloWorldRequest

	httpHandler.ParseHTTPRequest(r, &req)

	httpHandler.WriteResponse(w, JsonResponse{
		Error:   false,
		Message: "hello world",
	})

	return nil
}
