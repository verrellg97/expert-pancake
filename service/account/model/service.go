package model

import "net/http"

type AccountService interface {
	HelloWorld(w http.ResponseWriter, r *http.Request)
}
