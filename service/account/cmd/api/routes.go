package main

import (
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	"github.com/expert-pancake/service/account/model"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"net/http"
)

func (c *component) Routes(accountService model.AccountService) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Logger)

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	mux.Use(middleware.Heartbeat("/ping"))

	httpHandler.New(accountService.HelloWorld)

	mux.Method("GET", "/hello-world", httpHandler.New(accountService.HelloWorld))
	mux.Method("GET", "/hello-error", httpHandler.New(accountService.HelloError))
	mux.Method("POST", "/register", httpHandler.New(accountService.Register))
	mux.Method("POST", "/login", httpHandler.New(accountService.Login))

	// mux.Get("/hello-world", httpHandler.New(accountService.HelloWorld))

	return mux
}
