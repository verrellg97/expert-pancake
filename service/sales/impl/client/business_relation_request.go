package client

type GetContactBooksRequest struct {
	Id             string `json:"id"`
	CompanyId      string `json:"company_id" validate:"required"`
	ContactGroupId string `json:"contact_group_id"`
	Applicant      string `json:"applicant"`
}
