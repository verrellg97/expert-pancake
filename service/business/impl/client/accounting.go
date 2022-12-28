package client

const (
	RootPath                            = "http://103.174.67.78:4020"
	AddDefaultCompanyChartOfAccountPath = "/accounting/company/setting/chart-of-account"
)

type AccountingService interface {
	AddDefaultCompanyChartOfAccount(req AddDefaultCompanyChartOfAccountRequest) error
}
