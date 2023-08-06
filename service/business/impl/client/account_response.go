package client

type UserInformation struct {
	AccountId   string `json:"account_id" validate:"required"`
	ImageUrl    string `json:"image_url" validate:"required"`
	FullName    string `json:"full_name" validate:"required"`
	Nickname    string `json:"nickname" validate:"required"`
	Email       string `json:"email" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
}

type GetUserInformationResponse struct {
	Result struct {
		UserInformation
	} `json:"result"`
}
