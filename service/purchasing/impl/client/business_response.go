package client

type CompanyBranch struct {
	AccountId   string `json:"account_id" validate:"required"`
	CompanyId   string `json:"company_id" validate:"required"`
	BranchId    string `json:"branch_id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Address     string `json:"address" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	IsCentral   bool   `json:"is_central" validate:"required"`
}

type GetCompanyBranchesResponse struct {
	Result []CompanyBranch `json:"result" validate:"required"`
}
