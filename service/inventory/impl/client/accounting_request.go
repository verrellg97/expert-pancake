package client

type GetCompanyChartOfAccountsRequest struct {
	CompanyId       string  `json:"company_id" validate:"required"`
	Id              string  `json:"id"`
	Keyword         string  `json:"keyword"`
	TypeFilter      *string `json:"type_filter"`
	IsDeletedFilter *bool   `json:"is_deleted_filter"`
}
