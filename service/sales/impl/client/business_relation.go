package client

const (
	BusinessRelationRootPath  = "http://business-relation-service:4030"
	GetContactBooksPath       = "/business-relation/contact-books"
	GetKonekinContactBookPath = "/business-relation/konekin/contact-book"
)

type BusinessRelationService interface {
	GetContactBooks(req GetContactBooksRequest) error
	GetKonekinContactBook(req GetKonekinContactBookRequest) error
}
