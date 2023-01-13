package client

const (
	BusinessRelationRootPath  = "http://business-relation-service:4030"
	AddDefaultContactBookPath = "/business-relation/contact-book/default-data"
)

type BusinessRelationService interface {
	AddDefaultContactBook(req AddDefaultCompanyChartOfAccountRequest) error
}
