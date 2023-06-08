package client

const (
	SalesRootPath   = "http://sales-service:4080"
	GetCheckPOSPath = "/sales/pos/check"
)

type SalesService interface {
	GetCheckPOS(req GetCheckPOSRequest) error
}
