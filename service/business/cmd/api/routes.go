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
	DeleteCompanyPath          = "/business/company/delete"
	RegisterCompanyBranchPath  = "/business/company/branch/register"
	GetUserCompanyBranchesPath = "/business/user/company/branch"
	UpdateCompanyBranchPath    = "/business/company/branch/update"
	DeleteCompanyBranchPath    = "/business/company/branch/delete"
	GetCompanyBranchesPath     = "/business/company/branches"

	AddMemberRequestPath         = "/business/company/member-request/add"
	GetReceiveMemberRequestsPath = "/business/company/member-request/receive"
	UpdateMemberRequestPath      = "/business/company/member-request/update"

	GetCompaniesPath = "/business/companies"
)

func (c *component) Routes(businessService model.BusinessService) http.Handler {
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

	mux.Method("POST", RegisterCompanyPath, httpHandler.New(businessService.RegisterCompany))
	mux.Method("GET", GetCompanyTypesPath, httpHandler.New(businessService.GetCompanyTypes))
	mux.Method("POST", UpdateCompanyPath, httpHandler.New(businessService.UpdateCompany))
	mux.Method("POST", DeleteCompanyPath, httpHandler.New(businessService.DeleteCompany))
	mux.Method("POST", GetUserCompaniesPath, httpHandler.New(businessService.GetUserCompanies))
	mux.Method("POST", RegisterCompanyBranchPath, httpHandler.New(businessService.RegisterCompanyBranch))
	mux.Method("POST", GetUserCompanyBranchesPath, httpHandler.New(businessService.GetUserCompanyBranches))
	mux.Method("POST", UpdateCompanyBranchPath, httpHandler.New(businessService.UpdateCompanyBranch))
	mux.Method("POST", DeleteCompanyBranchPath, httpHandler.New(businessService.DeleteCompanyBranch))
	mux.Method("POST", GetCompanyBranchesPath, httpHandler.New(businessService.GetCompanyBranches))

	mux.Method("POST", AddMemberRequestPath, httpHandler.New(businessService.AddMemberRequest))
	mux.Method("POST", GetReceiveMemberRequestsPath, httpHandler.New(businessService.GetReceiveMemberRequests))
	mux.Method("POST", UpdateMemberRequestPath, httpHandler.New(businessService.UpdateMemberRequest))

	mux.Method("POST", GetCompaniesPath, httpHandler.New(businessService.GetCompanies))

	return mux
}
