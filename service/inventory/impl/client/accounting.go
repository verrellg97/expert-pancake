package client

const (
	AccountingRootPath            = "http://accounting-service:4020"
	GetCompanyChartOfAccountsPath = "/accounting/company/chart-of-accounts"
)

type AccountingService interface {
	GetCompanyChartOfAccounts(req GetCompanyChartOfAccountsRequest) error
}
