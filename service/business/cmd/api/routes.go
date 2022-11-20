package main

import (
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	"github.com/expert-pancake/service/business/model"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

const (
	RegisterCompanyPath        = "/business/company/register"
	GetCompanyTypesPath        = "/business/company/types"
	GetUserCompaniesPath       = "/business/user/company"
	UpdateCompanyPath          = "/business/company/update"
	RegisterCompanyBranchPath  = "/business/company/branch/register"
	GetUserCompanyBranchesPath = "/business/user/company/branch"
)

func (c *component) Routes(accountService model.BusinessService) http.Handler {
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

	mux.Method("POST", RegisterCompanyPath, httpHandler.New(accountService.RegisterCompany))
	mux.Method("GET", GetCompanyTypesPath, httpHandler.New(accountService.GetCompanyTypes))
	mux.Method("POST", UpdateCompanyPath, httpHandler.New(accountService.UpdateCompany))
	mux.Method("POST", GetUserCompaniesPath, httpHandler.New(accountService.GetUserCompanies))
	mux.Method("POST", RegisterCompanyBranchPath, httpHandler.New(accountService.RegisterCompanyBranch))
	mux.Method("POST", GetUserCompanyBranchesPath, httpHandler.New(accountService.GetUserCompanyBranches))

	return mux
}
