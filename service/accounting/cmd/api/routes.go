package main

import (
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	"github.com/expert-pancake/service/accounting/model"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

const (
	GetAccountingBanksPath        = "/accounting/banks"
	GetCompanyChartOfAccountsPath = "/accounting/company/chart-of-accounts"
)

func (c *component) Routes(accountingService model.AccountingService) http.Handler {
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

	mux.Method("GET", GetAccountingBanksPath, httpHandler.New(accountingService.GetAccountingBanks))
	mux.Method("POST", GetCompanyChartOfAccountsPath, httpHandler.New(accountingService.GetCompanyChartOfAccounts))

	return mux
}
