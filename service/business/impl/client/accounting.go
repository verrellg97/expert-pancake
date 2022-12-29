package client

const (
	RootPath                            = "http://accounting-service:4020"
	AddDefaultCompanyChartOfAccountPath = "/accounting/company/setting/chart-of-account"
)

type AccountingService interface {
	AddDefaultCompanyChartOfAccount(req AddDefaultCompanyChartOfAccountRequest) error
}
