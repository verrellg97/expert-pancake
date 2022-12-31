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
	AddDefaultCompanyChartOfAccountPath  = "/accounting/company/setting/chart-of-account"
	GetAccountingBanksPath               = "/accounting/banks"
	GetAccountingChartOfAccountTypesPath = "/accounting/chart-of-account/types"
	GetChartOfAccountGroupsPath          = "/accounting/company/chart-of-account-groups"
	AddChartOfAccountGroupPath           = "/accounting/company/chart-of-account-group/add"
	UpdateChartOfAccountGroupPath        = "/accounting/company/chart-of-account-group/update"
	GetCompanyChartOfAccountsPath        = "/accounting/company/chart-of-accounts"
	AddCompanyChartOfAccountPath         = "/accounting/company/chart-of-account/add"
	UpdateCompanyChartOfAccountPath      = "/accounting/company/chart-of-account/update"
	CheckCompanySettingStatePath         = "/accounting/company/setting/state"
	GetJournalBooksPath                  = "/accounting/company/journal-books"
	AddJournalBookPath                   = "/accounting/company/journal-book/add"
	GetMemorialJournalsPath              = "/accounting/company/memorial-journals"
	AddMemorialJournalPath               = "/accounting/company/memorial-journal/add"
	AddCashTransactionPath               = "/accounting/transaction/cash/add"
	GetCashTransactionsPath              = "/accounting/transaction/cash/list"
	GetCashTransactionsGroupByDatePath   = "/accounting/transaction/cash/list/group-by-date"
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

	mux.Method("POST", AddDefaultCompanyChartOfAccountPath, httpHandler.New(accountingService.AddDefaultCompanyChartOfAccount))
	mux.Method("GET", GetAccountingBanksPath, httpHandler.New(accountingService.GetAccountingBanks))
	mux.Method("GET", GetAccountingChartOfAccountTypesPath, httpHandler.New(accountingService.GetAccountingChartOfAccountTypes))
	mux.Method("POST", GetChartOfAccountGroupsPath, httpHandler.New(accountingService.GetChartOfAccountGroups))
	mux.Method("POST", AddChartOfAccountGroupPath, httpHandler.New(accountingService.AddChartOfAccountGroup))
	mux.Method("POST", UpdateChartOfAccountGroupPath, httpHandler.New(accountingService.UpdateChartOfAccountGroup))
	mux.Method("POST", GetCompanyChartOfAccountsPath, httpHandler.New(accountingService.GetCompanyChartOfAccounts))
	mux.Method("POST", AddCompanyChartOfAccountPath, httpHandler.New(accountingService.AddCompanyChartOfAccount))
	mux.Method("POST", UpdateCompanyChartOfAccountPath, httpHandler.New(accountingService.UpdateCompanyChartOfAccount))
	mux.Method("POST", CheckCompanySettingStatePath, httpHandler.New(accountingService.CheckCompanySettingState))
	mux.Method("POST", GetJournalBooksPath, httpHandler.New(accountingService.GetJournalBooks))
	mux.Method("POST", AddJournalBookPath, httpHandler.New(accountingService.AddJournalBook))
	mux.Method("POST", GetMemorialJournalsPath, httpHandler.New(accountingService.GetMemorialJournals))
	mux.Method("POST", AddMemorialJournalPath, httpHandler.New(accountingService.AddMemorialJournal))
	mux.Method("POST", AddCashTransactionPath, httpHandler.New(accountingService.AddCashTransaction))
	mux.Method("POST", GetCashTransactionsPath, httpHandler.New(accountingService.GetCashTransactions))
	mux.Method("POST", GetCashTransactionsGroupByDatePath, httpHandler.New(accountingService.GetCashTransactionsGroupByDate))

	return mux
}
