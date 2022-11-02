package main

import (
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	"github.com/expert-pancake/service/account/model"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"net/http"
)

const (
	RegisterPath                    = "/account/register"
	LoginPath                       = "/account/login"
	CheckPhoneNumberPath            = "/account/phone-number/check"
	PostOtpPath                     = "/account/otp/post"
	GetDefaultSecurityQuestionsPath = "/account/security-questions"
	UpsertUserAddressPath           = "/account/address/upsert"
	UpdateUserPath                  = "/account/user/update"
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
	mux.Method("GET", GetDefaultSecurityQuestionsPath, httpHandler.New(accountService.GetDefaultSecurityQuestions))

	mux.Method("POST", RegisterPath, httpHandler.New(accountService.Register))
	mux.Method("POST", LoginPath, httpHandler.New(accountService.Login))
	mux.Method("POST", CheckPhoneNumberPath, httpHandler.New(accountService.CheckPhoneNumber))
	mux.Method("POST", PostOtpPath, httpHandler.New(accountService.PostOtp))
	// mux.Method("POST", UpsertUserAddressPath, httpHandler.New(accountService.UpsertUserAddress))
	mux.Method("POST", UpdateUserPath, httpHandler.New(accountService.UpdateUser))

	// mux.Get("/hello-world", httpHandler.New(accountService.HelloWorld))

	return mux
}
