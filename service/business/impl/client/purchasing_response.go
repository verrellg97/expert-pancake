package client

type GetCheckPurchaseOrdersResponse struct {
	Result struct {
		Status bool `json:"status" validate:"required"`
	} `json:"result"`
}
