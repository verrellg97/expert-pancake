package client

const (
	BusinessRelationRootPath = "http://business-relation-service:4030"
	GetContactBooksPath      = "/business-relation/contact-books"
)

type BusinessRelationService interface {
	GetContactBooks(req GetContactBooksRequest) error
}
