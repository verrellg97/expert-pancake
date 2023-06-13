package client

const (
	BusinessRootPath       = "http://business-service:4010"
	GetCompanyBranchesPath = "/business/company/branches"
)

type BusinessService interface {
	GetCompanyBranches(req GetCompanyBranchesRequest) error
}
