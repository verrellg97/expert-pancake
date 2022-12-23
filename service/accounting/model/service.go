package model

import (
	"net/http"
)

type AccountingService interface {
	UpsertCompanyFiscalYear(w http.ResponseWriter, r *http.Request) error
	GetAccountingBanks(w http.ResponseWriter, r *http.Request) error
	GetAccountingChartOfAccountTypes(w http.ResponseWriter, r *http.Request) error
	GetChartOfAccountGroups(w http.ResponseWriter, r *http.Request) error
	AddChartOfAccountGroup(w http.ResponseWriter, r *http.Request) error
	UpdateChartOfAccountGroup(w http.ResponseWriter, r *http.Request) error
	GetCompanyChartOfAccounts(w http.ResponseWriter, r *http.Request) error
	AddCompanyChartOfAccount(w http.ResponseWriter, r *http.Request) error
	UpdateCompanyChartOfAccount(w http.ResponseWriter, r *http.Request) error
	CheckCompanySettingState(w http.ResponseWriter, r *http.Request) error
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

type ChartOfAccountType struct {
	ReportType  string `json:"report_type" validate:"required"`
	AccountType string `json:"account_type" validate:"required"`
}

type ChartOfAccountGroup struct {
	ChartOfAccountGroupId string `json:"chart_of_account_group_id" validate:"required"`
	CompanyId             string `json:"company_id" validate:"required"`
	ReportType            string `json:"report_type" validate:"required"`
	AccountType           string `json:"account_type" validate:"required"`
	AccountGroupName      string `json:"account_group_name" validate:"required"`
}

type GetChartOfAccountGroupsRequest struct {
	CompanyId string `json:"company_id" validate:"required"`
}

type AddChartOfAccounGroupRequest struct {
	CompanyId        string `json:"company_id" validate:"required"`
	ReportType       string `json:"report_type" validate:"required"`
	AccountType      string `json:"account_type" validate:"required"`
	AccountGroupName string `json:"account_group_name" validate:"required"`
}

type UpdateChartOfAccounGroupRequest struct {
	ChartOfAccountGroupId string `json:"chart_of_account_group_id" validate:"required"`
	ReportType            string `json:"report_type" validate:"required"`
	AccountType           string `json:"account_type" validate:"required"`
	AccountGroupName      string `json:"account_group_name" validate:"required"`
}

type UpsertChartOfAccountGroupResponse struct {
	ChartOfAccountGroup
}

type ChartOfAccount struct {
	ChartOfAccountId      string   `json:"chart_of_account_id" validate:"required"`
	CompanyId             string   `json:"company_id" validate:"required"`
	CurrencyCode          string   `json:"currency_code" validate:"required"`
	ChartOfAccountGroupId string   `json:"chart_of_account_group_id" validate:"required"`
	ReportType            string   `json:"report_type" validate:"required"`
	AccountType           string   `json:"account_type" validate:"required"`
	AccountGroup          string   `json:"account_group" validate:"required"`
	AccountCode           string   `json:"account_code" validate:"required"`
	AccountName           string   `json:"account_name" validate:"required"`
	BankName              string   `json:"bank_name" validate:"required"`
	BankAccountNumber     string   `json:"bank_account_number" validate:"required"`
	BankCode              string   `json:"bank_code" validate:"required"`
	IsAllBranches         bool     `json:"is_all_branches" validate:"required"`
	Branches              []string `json:"branches" validate:"required"`
	IsDeleted             bool     `json:"is_deleted" validate:"required"`
}

type ChartOfAccountBranch struct {
	ChartOfAccountId string `json:"chart_of_account_id" validate:"required"`
	BranchId         string `json:"branch_id" validate:"required"`
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
	ChartOfAccountTypes []ChartOfAccountType `json:"types"`
}

type GetCompanyChartOfAccountsRequest struct {
	CompanyId       string  `json:"company_id" validate:"required"`
	Keyword         string  `json:"keyword"`
	TypeFilter      *string `json:"type_filter"`
	IsDeletedFilter *bool   `json:"is_deleted_filter"`
}

type AddCompanyChartOfAccountRequest struct {
	CompanyId             string   `json:"company_id" validate:"required"`
	CurrencyCode          string   `json:"currency_code" validate:"required"`
	ChartOfAccountGroupId string   `json:"chart_of_account_group_id" validate:"required"`
	AccountCode           string   `json:"account_code" validate:"required"`
	AccountName           string   `json:"account_name" validate:"required"`
	BankName              string   `json:"bank_name"`
	BankAccountNumber     string   `json:"bank_account_number"`
	BankCode              string   `json:"bank_code"`
	IsAllBranches         bool     `json:"is_all_branches"`
	Branches              []string `json:"branches"`
}

type UpdateCompanyChartOfAccountRequest struct {
	ChartOfAccountId      string   `json:"chart_of_account_id" validate:"required"`
	CompanyId             string   `json:"company_id" validate:"required"`
	CurrencyCode          string   `json:"currency_code" validate:"required"`
	ChartOfAccountGroupId string   `json:"chart_of_account_group_id" validate:"required"`
	AccountCode           string   `json:"account_code" validate:"required"`
	AccountName           string   `json:"account_name" validate:"required"`
	BankName              string   `json:"bank_name"`
	BankAccountNumber     string   `json:"bank_account_number"`
	BankCode              string   `json:"bank_code"`
	IsAllBranches         bool     `json:"is_all_branches"`
	Branches              []string `json:"branches"`
	IsDeleted             bool     `json:"is_deleted"`
}

type CheckCompanySettingStateRequest struct {
	CompanyId string `json:"company_id" validate:"required"`
}

type CheckCompanySettingStateResponse struct {
	FiscalYear  *FiscalYear     `json:"fiscal_year"`
	BankAccount *ChartOfAccount `json:"bank_account"`
	CashAccount *ChartOfAccount `json:"cash_account"`
}

type AddCashTransactionRequest struct {
	CompanyId              string `json:"company_id" validate:"required"`
	BranchId               string `json:"branch_id" validate:"required"`
	TransactionDate        string `json:"transaction_date" validate:"required"`
	Type                   string `json:"type" validate:"required"`
	MainChartOfAccountId   string `json:"main_chart_of_account_id" validate:"required"`
	ContraChartOfAccountId string `json:"contra_chart_of_account_id"`
	Amount                 string `json:"amount" validate:"required"`
	Description            string `json:"description"`
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

type GetCashTransactionsGroupByDateResponse struct {
	TransactionDate string        `json:"transaction_date" validate:"required"`
	Amount          CashInCashOut `json:"amount" validate:"required"`
}
