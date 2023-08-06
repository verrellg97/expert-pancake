package client

const (
	AccountRootPath        = "http://account-service:4000"
	GetUserInformationPath = "/account/user/info"
)

type AccountService interface {
	GetUserInformation(req GetUserInformationRequest) error
}
