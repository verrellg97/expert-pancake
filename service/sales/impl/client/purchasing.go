package client

const (
	PurchasingRootPath     = "http://purchasing-service:4070"
	UpsertReceiptOrderPath = "/purchasing/receipt-order/upsert"
)

type PurchasingService interface {
	UpsertReceiptOrder(req UpsertReceiptOrderRequest) error
}
