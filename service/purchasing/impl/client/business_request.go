package client

type GetCompanyBranchesRequest struct {
	CompanyId string `json:"company_id" validate:"required"`
}
