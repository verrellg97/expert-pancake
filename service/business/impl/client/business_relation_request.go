package client

type AddDefaultContactBookRequest struct {
	CompanyId   string `json:"company_id" validate:"required"`
	CompanyName string `json:"company_name" validate:"required"`
}
