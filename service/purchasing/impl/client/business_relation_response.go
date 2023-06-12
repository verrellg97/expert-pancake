package client

type ContactBook struct {
	ContactBookId      string                   `json:"contact_book_id" validate:"required"`
	KonekinId          string                   `json:"konekin_id" validate:"required"`
	PrimaryCompanyId   string                   `json:"primary_company_id" validate:"required"`
	SecondaryCompanyId string                   `json:"secondary_company_id" validate:"required"`
	ContactGroupId     string                   `json:"contact_group_id" validate:"required"`
	ContactGroupName   string                   `json:"contact_group_name" validate:"required"`
	Name               string                   `json:"name" validate:"required"`
	Email              string                   `json:"email" validate:"required"`
	Phone              string                   `json:"phone" validate:"required"`
	Mobile             string                   `json:"mobile" validate:"required"`
	Web                string                   `json:"web" validate:"required"`
	AdditionalInfo     ContactBookAdditionaInfo `json:"additional_info" validate:"required"`
	MailingAddress     ContactBookAddress       `json:"mailing_address" validate:"required"`
	ShippingAddress    ContactBookAddress       `json:"shipping_address" validate:"required"`
	IsAllBranches      bool                     `json:"is_all_branches" validate:"required"`
	Branches           []string                 `json:"branches" validate:"required"`
	IsCustomer         bool                     `json:"is_customer" validate:"required"`
	CustomerCp         string                   `json:"customer_contact_person" validate:"required"`
	IsSupplier         bool                     `json:"is_supplier" validate:"required"`
	SupplierCp         string                   `json:"supplier_contact_person" validate:"required"`
}

type ContactBookAdditionaInfo struct {
	Nickname string `json:"nickname"`
	Tag      string `json:"tag"`
	Note     string `json:"note"`
}

type ContactBookAddress struct {
	Province    string `json:"province"`
	Regency     string `json:"regency"`
	District    string `json:"district"`
	PostalCode  string `json:"postal_code"`
	FullAddress string `json:"full_address"`
}

type GetContactBooksResponse struct {
	Result []ContactBook `json:"result" validate:"required"`
}

type KonekinContactBook struct {
	ContactBookId    string `json:"contact_book_id" validate:"required"`
	KonekinId        string `json:"konekin_id" validate:"required"`
	PrimaryCompanyId string `json:"primary_company_id" validate:"required"`
	Name             string `json:"name" validate:"required"`
	Email            string `json:"email" validate:"required"`
	Phone            string `json:"phone" validate:"required"`
	Mobile           string `json:"mobile" validate:"required"`
	Web              string `json:"web" validate:"required"`
}

type GetKonekinContactBookResponse struct {
	Result KonekinContactBook `json:"result" validate:"required"`
}
