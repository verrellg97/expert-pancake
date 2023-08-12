package client

type GetContactBooksRequest struct {
	Id             string `json:"id"`
	CompanyId      string `json:"company_id" validate:"required"`
	ContactGroupId string `json:"contact_group_id"`
	Applicant      string `json:"applicant"`
}

type GetKonekinContactBookRequest struct {
	PrimaryCompanyId   string `json:"primary_company_id" validate:"required"`
	SecondaryCompanyId string `json:"secondary_company_id" validate:"required"`
}
