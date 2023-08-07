package model

const (
	CreateNewCompanyError = 400100
	UpdateCompanyError    = 400101
	DeleteCompanyError    = 400102
	GetUserCompaniesError = 400104

	CompanyUniqueNameError        = 400105
	CompanyUniqueNameErrorMessage = "Name is already taken."

	CompanyUpdateTypeError        = 400106
	CompanyUpdateTypeErrorMessage = "Type can't be changed."

	CreateNewCompanyBranchError = 400200
	UpdateCompanyBranchError    = 400201
	DeleteCompanyBranchError    = 400202
	GetUserCompanyBranchesError = 400204
	GetCompanyBranchesError     = 400205

	AddMemberRequestError         = 400300
	GetReceiveMemberRequestsError = 400301
	UpdateMemberRequestError      = 400302

	GetPublicCompaniesError = 400400
)
