package client

const (
	BusinessRelationRootPath = "http://127.0.0.1:4030"
	GetContactBooksPath      = "/business-relation/contact-books"
)

type BusinessRelationService interface {
	GetContactBooks(req GetContactBooksRequest) error
}
