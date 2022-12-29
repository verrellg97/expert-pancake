package client

type AddDefaultCompanyChartOfAccountRequest struct {
	CompanyId string `json:"company_id" validate:"required"`
}
