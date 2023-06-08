package client

type GetCheckPurchaseOrdersRequest struct {
	CompanyId string `json:"company_id" validate:"required"`
}
