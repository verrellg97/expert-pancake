package client

type GetUserInformationRequest struct {
	AccountId string `json:"account_id" validate:"required"`
}
