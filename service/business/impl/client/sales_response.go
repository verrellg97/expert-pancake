package client

type GetCheckPOSResponse struct {
	Result struct {
		Status bool `json:"status" validate:"required"`
	} `json:"result"`
}
