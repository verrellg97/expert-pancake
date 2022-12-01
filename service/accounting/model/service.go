package model

import (
	"net/http"
)

type AccountingService interface {
	UpsertCompanyFiscalYear(w http.ResponseWriter, r *http.Request) error
	GetAccountingBanks(w http.ResponseWriter, r *http.Request) error
	GetAccountingChartOfAccountTypes(w http.ResponseWriter, r *http.Request) error
	GetCompanyChartOfAccounts(w http.ResponseWriter, r *http.Request) error
	AddCompanyChartOfAccount(w http.ResponseWriter, r *http.Request) error
	UpdateCompanyChartOfAccount(w http.ResponseWriter, r *http.Request) error
	CheckCompanySettingState(w http.ResponseWriter, r *http.Request) error
	GetAccountingTransactionTypes(w http.ResponseWriter, r *http.Request) error
	AddCashTransaction(w http.ResponseWriter, r *http.Request) error
	GetCashTransactions(w http.ResponseWriter, r *http.Request) error
	GetCashTransactionsGroupByDate(w http.ResponseWriter, r *http.Request) error
}

type Bank struct {
	BankName string `json:"bank_name" validate:"required"`
	BankCode string `json:"bank_code" validate:"required"`
}

type FiscalYear struct {
	CompanyId   string `json:"company_id" validate:"required"`
	StartPeriod string `json:"start_period" validate:"required"`
	EndPeriod   string `json:"end_period" validate:"required"`
}

type ChartOfAccount struct {
	ChartOfAccountId  string `json:"chart_of_account_id" validate:"required"`
	CompanyId         string `json:"company_id" validate:"required"`
	BranchId          string `json:"branch_id" validate:"required"`
	AccountCode       string `json:"account_code" validate:"required"`
	AccountName       string `json:"account_name" validate:"required"`
	AccountGroup      string `json:"account_group" validate:"required"`
	BankName          string `json:"bank_name" validate:"required"`
	BankAccountNumber string `json:"bank_account_number" validate:"required"`
	BankCode          string `json:"bank_code" validate:"required"`
	OpeningBalance    string `json:"opening_balance" validate:"required"`
	IsDeleted         bool   `json:"is_deleted" validate:"required"`
}

type ChartOfAccountIdName struct {
	ChartOfAccountId string `json:"chart_of_account_id" validate:"required"`
	AccountName      string `json:"account_name" validate:"required"`
}

type CashTransaction struct {
	CompanyId            string               `json:"company_id" validate:"required"`
	BranchId             string               `json:"branch_id" validate:"required"`
	TransactionId        string               `json:"transaction_id" validate:"required"`
	TransactionDate      string               `json:"transaction_date" validate:"required"`
	TransactionType      string               `json:"transaction_type" validate:"required"`
	Type                 string               `json:"type" validate:"required"`
	MainChartOfAccount   ChartOfAccountIdName `json:"main_chart_of_account" validate:"required"`
	ContraChartOfAccount ChartOfAccountIdName `json:"contra_chart_of_account"`
	Amount               string               `json:"amount" validate:"required"`
	Description          string               `json:"description" validate:"required"`
}

type CashInCashOut struct {
	CashIn  string `json:"cash_in" validate:"required"`
	CashOut string `json:"cash_out" validate:"required"`
}

type UpsertCompanyFiscalYearRequestResponse struct {
	CompanyId   string `json:"company_id" validate:"required"`
	StartPeriod string `json:"start_period" validate:"required"`
	EndPeriod   string `json:"end_period" validate:"required"`
}

type UpsertCompanyChartOfAccountResponse struct {
	ChartOfAccount
}

type GetAccountingBanksRequest struct {
	Type string `json:"type" validate:"required" schema:"type"`
}

type GetAccountingBanksResponse struct {
	Banks []Bank `json:"banks"`
}

type GetAccountingChartOfAccountTypesResponse struct {
	ChartOfAccountTypes []string `json:"types"`
}

type GetCompanyChartOfAccountsRequest struct {
	CompanyId       string `json:"company_id" validate:"required"`
	Keyword         string `json:"keyword"`
	GroupFilter     string `json:"group_filter"`
	IsDeletedFilter string `json:"is_deleted_filter"`
}

type AddCompanyChartOfAccountRequest struct {
	CompanyId         string `json:"company_id" validate:"required"`
	BranchId          string `json:"branch_id" validate:"required"`
	AccountCode       string `json:"account_code" validate:"required"`
	AccountName       string `json:"account_name" validate:"required"`
	AccountGroup      string `json:"account_group" validate:"required"`
	BankName          string `json:"bank_name"`
	BankAccountNumber string `json:"bank_account_number"`
	BankCode          string `json:"bank_code"`
	OpeningBalance    string `json:"opening_balance" validate:"required"`
}

type UpdateCompanyChartOfAccountRequest struct {
	ChartOfAccountId  string `json:"chart_of_account_id" validate:"required"`
	CompanyId         string `json:"company_id" validate:"required"`
	BranchId          string `json:"branch_id" validate:"required"`
	AccountCode       string `json:"account_code" validate:"required"`
	AccountName       string `json:"account_name" validate:"required"`
	AccountGroup      string `json:"account_group" validate:"required"`
	BankName          string `json:"bank_name"`
	BankAccountNumber string `json:"bank_account_number"`
	BankCode          string `json:"bank_code"`
	OpeningBalance    string `json:"opening_balance" validate:"required"`
	IsDeleted         bool   `json:"is_deleted"`
}

type CheckCompanySettingStateRequest struct {
	CompanyId string `json:"company_id" validate:"required"`
}

type CheckCompanySettingStateResponse struct {
	FiscalYear  *FiscalYear     `json:"fiscal_year"`
	BankAccount *ChartOfAccount `json:"bank_account"`
	CashAccount *ChartOfAccount `json:"cash_account"`
}

type GetAccountingTransactionTypesRequest struct {
	Type string `json:"type" validate:"required" schema:"type"`
}

type GetAccountingTransactionTypesResponse struct {
	TransactionTypes []string `json:"types"`
}

type AddCashTransactionRequest struct {
	CompanyId              string `json:"company_id" validate:"required"`
	BranchId               string `json:"branch_id" validate:"required"`
	TransactionDate        string `json:"transaction_date" validate:"required"`
	TransactionType        string `json:"transaction_type" validate:"required"`
	Type                   string `json:"type" validate:"required"`
	MainChartOfAccountId   string `json:"main_chart_of_account_id" validate:"required"`
	ContraChartOfAccountId string `json:"contra_chart_of_account_id"`
	Amount                 string `json:"amount" validate:"required"`
	Description            string `json:"description" validate:"required"`
}

type AddCashTransactionResponse struct {
	CashTransaction
}

type GetCashTransactionsRequest struct {
	CompanyId string `json:"company_id" validate:"required"`
	BranchId  string `json:"branch_id"  validate:"required"`
	Type      string `json:"type"`
}

type GetCashTransactionsGroupByDateRequest struct {
	CompanyId string `json:"company_id" validate:"required"`
	BranchId  string `json:"branch_id"  validate:"required"`
}

type GetCompanyCashTransactionsGroupByDateResponse struct {
	TransactionDate string        `json:"transaction_date" validate:"required"`
	Amount          CashInCashOut `json:"amount" validate:"required"`
}
