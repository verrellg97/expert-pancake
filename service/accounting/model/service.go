package model

import (
	"net/http"
)

type AccountingService interface {
	UpsertCompanyFiscalYear(w http.ResponseWriter, r *http.Request) error
	GetAccountingBanks(w http.ResponseWriter, r *http.Request) error
	GetCompanyChartOfAccounts(w http.ResponseWriter, r *http.Request) error
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
	IsDeleted         int32  `json:"is_deleted" validate:"required"`
}

type UpsertCompanyFiscalYearRequestResponse struct {
	CompanyId  string `json:"company_id" validate:"required"`
	StartMonth int    `json:"start_month" validate:"required"`
	StartYear  int    `json:"start_year" validate:"required"`
	EndMonth   int    `json:"end_month" validate:"required"`
	EndYear    int    `json:"end_year" validate:"required"`
}

type GetAccountingBanksRequest struct {
	Type string `json:"type" validate:"required"`
}

type GetAccountingBanksResponse struct {
	Banks []string `json:"banks"`
}

type CompanyChartOfAccountsRequest struct {
	CompanyId string `json:"company_id" validate:"required"`
	Keyword   string `json:"keyword"`
}
