package model

const (
	UpsertPOSError   = 800101
	DeletePOSError   = 800102
	GetPOSError      = 800103
	GetPOSItemsError = 800104

	GetPOSUserSettingError    = 800201
	UpdatePOSUserSettingError = 800202

	UpdatePOSCOASettingError = 800301
	GetPOSCOASettingError    = 800302

	UpdatePOSCustomerSettingError = 800401
	GetPOSCustomerSettingError    = 800402

	UpsertPOSPaymentMethodError = 800501
	DeletePOSPaymentMethodError = 800502
	GetPOSPaymentMethodError    = 800503

	GetCheckPOSError = 800601

	UpsertSalesOrderError       = 800701
	UpdateSalesOrderItemsError  = 800702
	GetSalesOrdersError         = 800703
	GetSalesOrderItemsError     = 800704
	UpdateSalesOrderStatusError = 800705

	UpsertDeliveryOrderError        = 800801
	GetDeliveryOrdersError          = 800802
	GetSalesOrderDeliveryItemsError = 800803
	UpdateDeliveryOrderItemsError   = 800804
	GetDeliveryOrderItemsError      = 800805

	UpsertSalesInvoiceError   = 800901
	GetSalesInvoicesError     = 800902
	GetSalesInvoiceItemsError = 800903

	GetSalesSummaryReportError = 801001
	GetMostSoldItemsError      = 801002
	GetMonthlyGrossSalesError  = 801003
)
