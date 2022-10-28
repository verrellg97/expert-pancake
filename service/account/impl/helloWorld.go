package impl

import (
	"encoding/json"
	"log"
	"net/http"
)

type JsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

func (a accountService) HelloWorld(w http.ResponseWriter, r *http.Request) {
	resp, err := json.Marshal(JsonResponse{
		Error:   false,
		Message: "hello world",
	})

	if err != nil {
		log.Panic(err)
	}

	w.Write(resp)
}
