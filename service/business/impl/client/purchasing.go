package client

const (
	PurchasingRootPath         = "http://purchasing-service:4070"
	GetCheckPurchaseOrdersPath = "/purchasing/orders/check"
)

type PurchasingService interface {
	GetCheckPurchaseOrders(req GetCheckPOSRequest) error
}
