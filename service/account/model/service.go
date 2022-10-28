package model

import "net/http"

type AccountService interface {
	HelloWorld(w http.ResponseWriter, r *http.Request) error
	HelloError(w http.ResponseWriter, r *http.Request) error
}
