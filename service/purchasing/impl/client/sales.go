package client

const (
	SalesRootPath             = "http://sales-service:4080"
	UpsertSalesOrderPath      = "/sales/order/upsert"
	UpdateSalesOrderItemsPath = "/sales/order/items/update"
)

type SalesService interface {
	UpsertSalesOrder(req UpsertSalesOrderRequest) error
	UpdateSalesOrderItems(req UpdateSalesOrderItemsRequest) error
}
