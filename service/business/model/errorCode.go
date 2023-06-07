package model

const (
	CreateNewCompanyError = 400100
	UpdateCompanyError    = 400101
	DeleteCompanyError    = 400102
	GetUserCompaniesError = 400104

	CompanyUniqueNameError        = 400105
	CompanyUniqueNameErrorMessage = "Name is already taken."

	CreateNewCompanyBranchError = 400200
	UpdateCompanyBranchError    = 400201
	DeleteCompanyBranchError    = 400202
	GetUserCompanyBranchesError = 400204
)
