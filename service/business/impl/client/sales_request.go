package client

type GetCheckPOSRequest struct {
	CompanyId string `json:"company_id" validate:"required"`
}
