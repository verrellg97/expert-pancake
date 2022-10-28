package impl

import (
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	"net/http"
)

type JsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

type HelloWorldRequest struct {
}

func (a accountService) HelloWorld(w http.ResponseWriter, r *http.Request) error {

	var req HelloWorldRequest

	httpHandler.ParseHTTPRequest(r, &req)

	httpHandler.WriteResponse(w, JsonResponse{
		Error:   false,
		Message: "hello world",
	})

	return nil
}
