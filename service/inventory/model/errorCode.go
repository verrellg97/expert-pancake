package model

const (
	AddNewBrandError = 700001
	UpdateBrandError = 700002
	DeleteBrandError = 700003
	GetBrandsError   = 700004

	AddNewGroupError = 700005
	UpdateGroupError = 700006
	DeleteGroupError = 700007
	GetGroupsError   = 700008

	AddNewUnitError = 700009
	UpdateUnitError = 700010
	GetUnitsError   = 700011

	AddNewItemError = 700012
	UpdateItemError = 700013
	GetItemsError   = 700014

	UpsertItemVariantError = 700015
	GetItemVariantsError   = 700016

	UpsertItemUnitError = 700017
	GetItemUnitsError   = 700018
	DeleteItemUnitError = 700619

	AddNewInternalStockTransferError = 700019
	GetInternalStockTransfersError   = 700020

	UpsertItemReorderError = 700021
	GetItemReordersError   = 700022

	UpsertUnitCategoryError = 700023
	GetUnitCategoriesError  = 700024

	UpsertItemInfoError = 700025
	GetItemInfoError    = 700026

	AddNewUpdateStockError = 700027
	GetUpdateStocksError   = 700028

	GetVariantWarehouseRacksError                 = 700029
	GetVariantWarehouseRackBatchesError           = 700030
	GetVariantWarehouseRackBatchExpiredDatesError = 700031
	GetVariantWarehouseRackStockError             = 700032
	GetVariantWarehouseStocksError                = 700033

	GetTransferHistoryError = 700034
	GetStockHistoryError    = 700035

	GetItemReorderNotificationsError = 700036

	GetSupplierCatalogsError    = 700037
	GetMappingItemsError        = 700038
	GetMappingItemVariantsError = 700039
	GetMappingItemUnitsError    = 700040

	UpsertItemVariantMapError = 700041
	GetItemVariantMapsError   = 700042

	UpsertPricelistError = 700043
	GetPricelistsError   = 700044

	UpsertPricelistItemsError = 700045
	GetPricelistItemsError    = 700046

	GetPurchaseItemsError            = 700047
	GetPurchaseItemVariantsError     = 700048
	GetPurchaseItemVariantUnitsError = 700049

	GetPOSItemsError                      = 700050
	GetVariantWarehouseRacksByBranchError = 700051

	GetCheckStockHistoryError = 700052

	InsertStokMovementError = 700053
	DeleteStokMovement      = 700054

	GetUnderMinimumOrderError = 700055
	GetOutgoingStockError    = 700056
)
