package client

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

type GetCompanyChartOfAccountsResponse struct {
	Result []ChartOfAccount `json:"result" validate:"required"`
}
