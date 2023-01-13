package model

import (
	"net/http"
)

type BusinessRelationService interface {
	AddContactGroup(w http.ResponseWriter, r *http.Request) error
	UpdateContactGroup(w http.ResponseWriter, r *http.Request) error
	GetContactGroups(w http.ResponseWriter, r *http.Request) error

	AddDefaultContactBook(w http.ResponseWriter, r *http.Request) error
	GetMyContactBook(w http.ResponseWriter, r *http.Request) error
	AddContactBook(w http.ResponseWriter, r *http.Request) error
	UpdateContactBook(w http.ResponseWriter, r *http.Request) error
	GetContactBooks(w http.ResponseWriter, r *http.Request) error

	UpdateCustomer(w http.ResponseWriter, r *http.Request) error
	GetCustomers(w http.ResponseWriter, r *http.Request) error

	UpdateSupplier(w http.ResponseWriter, r *http.Request) error
	GetSuppliers(w http.ResponseWriter, r *http.Request) error
}

type ContactGroup struct {
	GroupId     string `json:"group_id" validate:"required"`
	CompanyId   string `json:"company_id" validate:"required"`
	ImageUrl    string `json:"image_url" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type ContactGroupWithMember struct {
	GroupId     string `json:"group_id" validate:"required"`
	CompanyId   string `json:"company_id" validate:"required"`
	ImageUrl    string `json:"image_url" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Member      string `json:"member" validate:"required"`
}

type AddContactGroupRequest struct {
	CompanyId   string   `json:"company_id" validate:"required"`
	ImageUrl    string   `json:"image_url"`
	Name        string   `json:"name" validate:"required"`
	Description string   `json:"description"`
	Members     []string `json:"members"`
}

type AddContactGroupResponse struct {
	ContactGroupWithMember
}

type UpdateContactGroupRequest struct {
	GroupId     string `json:"group_id" validate:"required"`
	ImageUrl    string `json:"image_url"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}

type UpdateContactGroupResponse struct {
	ContactGroup
}

type GetContactGroupsRequest struct {
	CompanyId string `json:"company_id" validate:"required"`
}

type MyContactBook struct {
	ContactBookId    string                   `json:"contact_book_id" validate:"required"`
	KonekinId        string                   `json:"konekin_id" validate:"required"`
	PrimaryCompanyId string                   `json:"primary_company_id" validate:"required"`
	Name             string                   `json:"name" validate:"required"`
	Email            string                   `json:"email" validate:"required"`
	Phone            string                   `json:"phone" validate:"required"`
	Mobile           string                   `json:"mobile" validate:"required"`
	Web              string                   `json:"web" validate:"required"`
	AdditionalInfo   ContactBookAdditionaInfo `json:"additional_info" validate:"required"`
	MailingAddress   ContactBookAddress       `json:"mailing_address" validate:"required"`
	ShippingAddress  ContactBookAddress       `json:"shipping_address" validate:"required"`
}

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

type CustomerInfo struct {
	ContactBookId    string `json:"contact_book_id" validate:"required"`
	ContactGroupName string `json:"contact_group_name" validate:"required"`
	Name             string `json:"name" validate:"required"`
	Email            string `json:"email" validate:"required"`
	Phone            string `json:"phone" validate:"required"`
	Mobile           string `json:"mobile" validate:"required"`
	Web              string `json:"web" validate:"required"`
	IsTax            bool   `json:"is_tax" validate:"required"`
	TaxId            string `json:"tax_id" validate:"required"`
	Pic              string `json:"pic" validate:"required"`
	CreditLimit      string `json:"credit_limit" validate:"required"`
	PaymentTerm      string `json:"payment_term" validate:"required"`
}

type SupplierInfo struct {
	ContactBookId    string `json:"contact_book_id" validate:"required"`
	ContactGroupName string `json:"contact_group_name" validate:"required"`
	Name             string `json:"name" validate:"required"`
	Email            string `json:"email" validate:"required"`
	Phone            string `json:"phone" validate:"required"`
	Mobile           string `json:"mobile" validate:"required"`
	Web              string `json:"web" validate:"required"`
	IsTax            bool   `json:"is_tax" validate:"required"`
	TaxId            string `json:"tax_id" validate:"required"`
	Pic              string `json:"pic" validate:"required"`
	CreditLimit      string `json:"credit_limit" validate:"required"`
	PaymentTerm      string `json:"payment_term" validate:"required"`
}

type AddContactBookRequest struct {
	PrimaryCompanyId string                   `json:"primary_company_id" validate:"required"`
	ContactGroupId   string                   `json:"contact_group_id"`
	Name             string                   `json:"name" validate:"required"`
	Email            string                   `json:"email"`
	Phone            string                   `json:"phone"`
	Mobile           string                   `json:"mobile"`
	Web              string                   `json:"web"`
	AdditionalInfo   ContactBookAdditionaInfo `json:"additional_info"`
	MailingAddress   ContactBookAddress       `json:"mailing_address"`
	ShippingAddress  ContactBookAddress       `json:"shipping_address"`
	IsAllBranches    bool                     `json:"is_all_branches"`
	Branches         []string                 `json:"branches"`
	IsCustomer       bool                     `json:"is_customer"`
	IsSupplier       bool                     `json:"is_supplier"`
}

type AddContactBookResponse struct {
	ContactBook
}

type UpdateContactBookRequest struct {
	ContactBookId   string                   `json:"contact_book_id" validate:"required"`
	ContactGroupId  string                   `json:"contact_group_id"`
	Name            string                   `json:"name" validate:"required"`
	Email           string                   `json:"email"`
	Phone           string                   `json:"phone"`
	Mobile          string                   `json:"mobile"`
	Web             string                   `json:"web"`
	AdditionalInfo  ContactBookAdditionaInfo `json:"additional_info"`
	MailingAddress  ContactBookAddress       `json:"mailing_address"`
	ShippingAddress ContactBookAddress       `json:"shipping_address"`
	IsAllBranches   bool                     `json:"is_all_branches"`
	Branches        []string                 `json:"branches"`
	IsCustomer      bool                     `json:"is_customer"`
	IsSupplier      bool                     `json:"is_supplier"`
}

type UpdateContactBookResponse struct {
	ContactBook
}

type GetContactBooksRequest struct {
	CompanyId      string `json:"company_id" validate:"required"`
	ContactGroupId string `json:"contact_group_id"`
}

type UpdateCustomerRequest struct {
	ContactBookId string `json:"contact_book_id" validate:"required"`
	IsTax         bool   `json:"is_tax"`
	TaxId         string `json:"tax_id"`
	Pic           string `json:"pic"`
	CreditLimit   string `json:"credit_limit"`
	PaymentTerm   string `json:"payment_term"`
}

type UpdateCustomerResponse struct {
	CustomerInfo
}

type GetCustomersRequest struct {
	CompanyId string `json:"company_id" validate:"required"`
}

type UpdateSupplierRequest struct {
	ContactBookId string `json:"contact_book_id" validate:"required"`
	IsTax         bool   `json:"is_tax"`
	TaxId         string `json:"tax_id"`
	Pic           string `json:"pic"`
	CreditLimit   string `json:"credit_limit"`
	PaymentTerm   string `json:"payment_term"`
}

type UpdateSupplierResponse struct {
	SupplierInfo
}

type GetSuppliersRequest struct {
	CompanyId string `json:"company_id" validate:"required"`
}

type AddDefaultContactBookRequest struct {
	CompanyId   string `json:"company_id" validate:"required"`
	CompanyName string `json:"company_name" validate:"required"`
}

type AddDefaultContactBookResponse struct {
	Message string `json:"message"`
}

type GetMyContactBookRequest struct {
	CompanyId string `json:"company_id" validate:"required"`
}
