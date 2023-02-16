package model

import (
	"net/http"
)

type InventoryService interface {
	AddBrand(w http.ResponseWriter, r *http.Request) error
	UpdateBrand(w http.ResponseWriter, r *http.Request) error
	GetBrands(w http.ResponseWriter, r *http.Request) error

	AddGroup(w http.ResponseWriter, r *http.Request) error
	UpdateGroup(w http.ResponseWriter, r *http.Request) error
	GetGroups(w http.ResponseWriter, r *http.Request) error

	AddUnit(w http.ResponseWriter, r *http.Request) error
	UpdateUnit(w http.ResponseWriter, r *http.Request) error
	GetUnits(w http.ResponseWriter, r *http.Request) error

	AddItem(w http.ResponseWriter, r *http.Request) error
	UpdateItem(w http.ResponseWriter, r *http.Request) error
	GetItems(w http.ResponseWriter, r *http.Request) error

	UpsertItemInfo(w http.ResponseWriter, r *http.Request) error
	GetItemInfo(w http.ResponseWriter, r *http.Request) error

	UpsertItemVariant(w http.ResponseWriter, r *http.Request) error
	GetItemVariants(w http.ResponseWriter, r *http.Request) error

	UpsertItemUnit(w http.ResponseWriter, r *http.Request) error
	GetItemUnits(w http.ResponseWriter, r *http.Request) error

	AddInternalStockTransfer(w http.ResponseWriter, r *http.Request) error
	GetInternalStockTransfers(w http.ResponseWriter, r *http.Request) error

	AddUpdateStock(w http.ResponseWriter, r *http.Request) error
	GetUpdateStocks(w http.ResponseWriter, r *http.Request) error

	UpsertItemReorder(w http.ResponseWriter, r *http.Request) error
	GetItemReorders(w http.ResponseWriter, r *http.Request) error

	UpsertUnitCategory(w http.ResponseWriter, r *http.Request) error
	GetUnitCategories(w http.ResponseWriter, r *http.Request) error

	GetVariantWarehouseRacks(w http.ResponseWriter, r *http.Request) error
	GetVariantWarehouseRackBatches(w http.ResponseWriter, r *http.Request) error
	GetVariantWarehouseRackBatchExpiredDates(w http.ResponseWriter, r *http.Request) error

	GetTransferHistory(w http.ResponseWriter, r *http.Request) error
}

type Brand struct {
	CompanyId string `json:"company_id" validate:"required"`
	BrandId   string `json:"brand_id" validate:"required"`
	Name      string `json:"name" validate:"required"`
}

type AddBrandRequest struct {
	CompanyId string `json:"company_id" validate:"required"`
	Name      string `json:"name" validate:"required"`
}

type AddBrandResponse struct {
	Brand
}

type UpdateBrandRequest struct {
	Id   string `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

type UpdateBrandResponse struct {
	Brand
}

type GetBrandsRequest struct {
	CompanyId string `json:"company_id" validate:"required"`
	Keyword   string `json:"keyword"`
}

type Group struct {
	CompanyId string `json:"company_id" validate:"required"`
	GroupId   string `json:"group_id" validate:"required"`
	Name      string `json:"name" validate:"required"`
}

type AddGroupRequest struct {
	CompanyId string `json:"company_id" validate:"required"`
	Name      string `json:"name" validate:"required"`
}

type AddGroupResponse struct {
	Group
}

type UpdateGroupRequest struct {
	Id   string `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

type UpdateGroupResponse struct {
	Group
}

type GetGroupsRequest struct {
	CompanyId string `json:"company_id" validate:"required"`
	Keyword   string `json:"keyword"`
}

type Unit struct {
	CompanyId      string `json:"company_id" validate:"required"`
	UnitCategoryId string `json:"unit_category_id" validate:"required"`
	UnitId         string `json:"unit_id" validate:"required"`
	Name           string `json:"name" validate:"required"`
}

type AddUnitRequest struct {
	CompanyId      string `json:"company_id" validate:"required"`
	UnitCategoryId string `json:"unit_category_id" validate:"required"`
	Name           string `json:"name" validate:"required"`
}

type AddUnitResponse struct {
	Unit
}

type UpdateUnitRequest struct {
	Id             string `json:"id" validate:"required"`
	UnitCategoryId string `json:"unit_category_id" validate:"required"`
	Name           string `json:"name" validate:"required"`
}

type UpdateUnitResponse struct {
	Unit
}

type GetUnitsRequest struct {
	CompanyId      string `json:"company_id" validate:"required"`
	UnitCategoryId string `json:"unit_category_id"`
	Keyword        string `json:"keyword"`
}

type Item struct {
	CompanyId   string `json:"company_id" validate:"required"`
	ItemId      string `json:"item_id" validate:"required"`
	VariantId   string `json:"variant_id" validate:"required"`
	ImageUrl    string `json:"image_url" validate:"required"`
	Code        string `json:"code" validate:"required"`
	Barcode     string `json:"barcode" validate:"required"`
	Name        string `json:"name" validate:"required"`
	VariantName string `json:"variant_name" validate:"required"`
	BrandId     string `json:"brand_id" validate:"required"`
	BrandName   string `json:"brand_name" validate:"required"`
	GroupId     string `json:"group_id" validate:"required"`
	GroupName   string `json:"group_name" validate:"required"`
	Tag         string `json:"tag" validate:"required"`
	Description string `json:"description" validate:"required"`
	IsDefault   bool   `json:"is_default" validate:"required"`
	Price       string `json:"price" validate:"required"`
}

type AddItemRequest struct {
	CompanyId   string `json:"company_id" validate:"required"`
	ImageUrl    string `json:"image_url"`
	Barcode     string `json:"barcode"`
	Name        string `json:"name" validate:"required"`
	BrandId     string `json:"brand_id" validate:"required"`
	GroupId     string `json:"group_id" validate:"required"`
	Tag         string `json:"tag"`
	Price       string `json:"price" validate:"required"`
	Description string `json:"description"`
}

type AddItemResponse struct {
	Item
}

type UpdateItemRequest struct {
	ItemId      string `json:"item_id" validate:"required"`
	ImageUrl    string `json:"image_url"`
	Barcode     string `json:"barcode"`
	Name        string `json:"name" validate:"required"`
	BrandId     string `json:"brand_id" validate:"required"`
	GroupId     string `json:"group_id" validate:"required"`
	Tag         string `json:"tag"`
	Price       string `json:"price" validate:"required"`
	Description string `json:"description"`
}

type UpdateItemResponse struct {
	Item
}

type GetItemsRequest struct {
	CompanyId string `json:"company_id" validate:"required"`
	Keyword   string `json:"keyword"`
}

type GetItemsResponse struct {
	Items []Item `json:"items" validate:"required"`
}

type ItemInfo struct {
	ItemId                     string `json:"item_id" validate:"required"`
	IsPurchase                 bool   `json:"is_purchase" validate:"required"`
	IsSale                     bool   `json:"is_sale" validate:"required"`
	IsRawMaterial              bool   `json:"is_raw_material" validate:"required"`
	IsAsset                    bool   `json:"is_asset" validate:"required"`
	PurchaseChartOfAccountId   string `json:"purchase_chart_of_account_id" validate:"required"`
	PurchaseChartOfAccountName string `json:"purchase_chart_of_account_name" validate:"required"`
	SaleChartOfAccountId       string `json:"sale_chart_of_account_id" validate:"required"`
	SaleChartOfAccountName     string `json:"sale_chart_of_account_name" validate:"required"`
	PurchaseItemUnitId         string `json:"purchase_item_unit_id" validate:"required"`
	PurchaseItemUnitName       string `json:"purchase_item_unit_name" validate:"required"`
}

type UpsertItemInfoRequest struct {
	ItemId                   string `json:"item_id" validate:"required"`
	IsPurchase               bool   `json:"is_purchase"`
	IsSale                   bool   `json:"is_sale"`
	IsRawMaterial            bool   `json:"is_raw_material"`
	IsAsset                  bool   `json:"is_asset"`
	PurchaseChartOfAccountId string `json:"purchase_chart_of_account_id" validate:"required"`
	SaleChartOfAccountId     string `json:"sale_chart_of_account_id" validate:"required"`
	PurchaseItemUnitId       string `json:"purchase_item_unit_id" validate:"required"`
}

type UpsertItemInfoResponse struct {
	ItemInfo
}

type GetItemInfoRequest struct {
	ItemId string `json:"item_id" validate:"required"`
}

type GetItemInfoResponse struct {
	ItemInfo
}

type UpsertItemVariantRequest struct {
	ItemVariantId string `json:"item_variant_id"`
	ItemId        string `json:"item_id" validate:"required"`
	ImageUrl      string `json:"image_url"`
	Barcode       string `json:"barcode"`
	Name          string `json:"name" validate:"required"`
	Price         string `json:"price" validate:"required"`
}

type UpsertItemVariantResponse struct {
	Item
}

type GetItemVariantsRequest struct {
	ItemId  string `json:"item_id" validate:"required"`
	Keyword string `json:"keyword"`
}

type GetItemVariantsResponse struct {
	ItemVariants []Item `json:"item_variants" validate:"required"`
}

type ItemUnit struct {
	ItemUnitId string `json:"item_unit_id" validate:"required"`
	ItemId     string `json:"item_id" validate:"required"`
	UnitId     string `json:"unit_id" validate:"required"`
	UnitName   string `json:"unit_name" validate:"required"`
	Value      string `json:"value" validate:"required"`
	IsDefault  bool   `json:"is_default" validate:"required"`
}

type UpsertItemUnitRequest struct {
	ItemUnitId string `json:"item_unit_id"`
	ItemId     string `json:"item_id" validate:"required"`
	UnitId     string `json:"unit_id" validate:"required"`
	Value      string `json:"value" validate:"required"`
	IsDefault  bool   `json:"is_default"`
}

type UpsertItemUnitResponse struct {
	ItemUnit
}

type GetItemUnitsRequest struct {
	ItemId  string `json:"item_id" validate:"required"`
	Keyword string `json:"keyword"`
}

type GetItemUnitsResponse struct {
	ItemUnits []ItemUnit `json:"item_units" validate:"required"`
}

type InternalStockTransferItem struct {
	DetailId          string  `json:"detail_id" validate:"required"`
	WarehouseRackId   string  `json:"warehouse_rack_id" validate:"required"`
	WarehouseRackName string  `json:"warehouse_rack_name" validate:"required"`
	ItemName          string  `json:"item_name" validate:"required"`
	VariantId         string  `json:"variant_id" validate:"required"`
	VariantName       string  `json:"variant_name" validate:"required"`
	ItemUnitId        string  `json:"item_unit_id" validate:"required"`
	ItemUnitName      string  `json:"item_unit_name" validate:"required"`
	ItemUnitValue     string  `json:"item_unit_value" validate:"required"`
	Amount            string  `json:"amount" validate:"required"`
	Batch             *string `json:"batch" validate:"required"`
	ExpiredDate       *string `json:"expired_date" validate:"required"`
}

type InternalStockTransfer struct {
	TransactionId            string                      `json:"transaction_id" validate:"required"`
	SourceWarehouseId        string                      `json:"source_warehouse_id" validate:"required"`
	SourceWarehouseName      string                      `json:"source_warehouse_name" validate:"required"`
	DestinationWarehouseId   string                      `json:"destination_warehouse_id" validate:"required"`
	DestinationWarehouseName string                      `json:"destination_warehouse_name" validate:"required"`
	FormNumber               string                      `json:"form_number" validate:"required"`
	TransactionDate          string                      `json:"transaction_date" validate:"required"`
	Items                    []InternalStockTransferItem `json:"items" validate:"required"`
}

type InternalStockTransferItemRequest struct {
	WarehouseRackId string `json:"warehouse_rack_id"`
	VariantId       string `json:"variant_id" validate:"required"`
	ItemUnitId      string `json:"item_unit_id" validate:"required"`
	ItemUnitValue   string `json:"item_unit_value" validate:"required"`
	Amount          string `json:"amount" validate:"required"`
	Batch           string `json:"batch"`
	ExpiredDate     string `json:"expired_date"`
}

type AddInternalStockTransferRequest struct {
	SourceWarehouseId      string                             `json:"source_warehouse_id" validate:"required"`
	DestinationWarehouseId string                             `json:"destination_warehouse_id" validate:"required"`
	FormNumber             string                             `json:"form_number" validate:"required"`
	TransactionDate        string                             `json:"transaction_date" validate:"required"`
	Items                  []InternalStockTransferItemRequest `json:"items" validate:"required"`
}

type AddInternalStockTransferResponse struct {
	InternalStockTransfer
}

type GetInternalStockTransfersRequest struct {
	BranchId  string `json:"branch_id" validate:"required"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type GetInternalStockTransfersResponse struct {
	InternalStockTransfers []InternalStockTransfer `json:"internal_stock_transfers" validate:"required"`
}

type UpdateStock struct {
	TransactionId     string  `json:"transaction_id" validate:"required"`
	FormNumber        string  `json:"form_number" validate:"required"`
	TransactionDate   string  `json:"transaction_date" validate:"required"`
	WarehouseId       string  `json:"warehouse_id" validate:"required"`
	WarehouseName     string  `json:"warehouse_name" validate:"required"`
	WarehouseRackId   string  `json:"warehouse_rack_id" validate:"required"`
	WarehouseRackName string  `json:"warehouse_rack_name" validate:"required"`
	ItemId            string  `json:"item_id" validate:"required"`
	ItemName          string  `json:"item_name" validate:"required"`
	VariantId         string  `json:"variant_id" validate:"required"`
	VariantName       string  `json:"variant_name" validate:"required"`
	ItemUnitId        string  `json:"item_unit_id" validate:"required"`
	ItemUnitName      string  `json:"item_unit_name" validate:"required"`
	ItemUnitValue     string  `json:"item_unit_value" validate:"required"`
	BeginningStock    string  `json:"beginning_stock" validate:"required"`
	EndingStock       string  `json:"ending_stock" validate:"required"`
	Batch             *string `json:"batch" validate:"required"`
	ExpiredDate       *string `json:"expired_date" validate:"required"`
}

type AddUpdateStockRequest struct {
	TransactionDate string `json:"transaction_date" validate:"required"`
	WarehouseId     string `json:"warehouse_id" validate:"required"`
	WarehouseRackId string `json:"warehouse_rack_id" validate:"required"`
	VariantId       string `json:"variant_id" validate:"required"`
	ItemUnitId      string `json:"item_unit_id" validate:"required"`
	ItemUnitValue   string `json:"item_unit_value" validate:"required"`
	BeginningStock  string `json:"beginning_stock" validate:"required"`
	EndingStock     string `json:"ending_stock" validate:"required"`
	Batch           string `json:"batch"`
	ExpiredDate     string `json:"expired_date"`
}

type AddUpdateStockResponse struct {
	UpdateStock
}

type GetUpdateStocksRequest struct {
	BranchId  string `json:"branch_id" validate:"required"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type GetUpdateStocksResponse struct {
	UpdateStocks []UpdateStock `json:"update_stocks" validate:"required"`
}

type ItemReorder struct {
	Id            string `json:"id" validate:"required"`
	ItemId        string `json:"item_id" validate:"required"`
	ItemName      string `json:"item_name" validate:"required"`
	VariantId     string `json:"variant_id" validate:"required"`
	VariantName   string `json:"variant_name" validate:"required"`
	ItemUnitId    string `json:"item_unit_id" validate:"required"`
	ItemUnitName  string `json:"item_unit_name" validate:"required"`
	WarehouseId   string `json:"warehouse_id" validate:"required"`
	WarehouseName string `json:"warehouse_name" validate:"required"`
	MinimumStock  string `json:"minimum_stock" validate:"required"`
}
type UpsertItemReorderRequest struct {
	Id           string `json:"id"`
	VariantId    string `json:"variant_id" validate:"required"`
	ItemUnitId   string `json:"item_unit_id" validate:"required"`
	WarehouseId  string `json:"warehouse_id" validate:"required"`
	MinimumStock string `json:"minimum_stock" validate:"required"`
}

type UpsertItemReorderResponse struct {
	ItemReorder ItemReorder `json:"item_reorder" validate:"required"`
}

type GetItemReordersRequest struct {
	ItemId      string `json:"item_id"`
	WarehouseId string `json:"warehouse_id"`
}

type GetItemReordersResponse struct {
	ItemReorders []ItemReorder `json:"item_reorders" validate:"required"`
}

type UnitCategory struct {
	Id        string `json:"id" validate:"required"`
	CompanyId string `json:"company_id" validate:"required"`
	Name      string `json:"name" validate:"required"`
}

type UpsertUnitCategoryRequest struct {
	Id        string `json:"id"`
	CompanyId string `json:"company_id" validate:"required"`
	Name      string `json:"name" validate:"required"`
}

type UpsertUnitCategoryResponse struct {
	UnitCategory UnitCategory `json:"unit_category" validate:"required"`
}

type GetUnitCategoriesRequest struct {
	CompanyId string `json:"company_id" validate:"required"`
	Keyword   string `json:"keyword"`
}

type GetUnitCategoriesResponse struct {
	UnitCategories []UnitCategory `json:"unit_categories" validate:"required"`
}

type WarehouseRack struct {
	RackId      string `json:"rack_id" validate:"required"`
	WarehouseId string `json:"warehouse_id" validate:"required"`
	Name        string `json:"name" validate:"required"`
}

type GetVariantWarehouseRacksRequest struct {
	WarehouseId string `json:"warehouse_id" validate:"required"`
	VariantId   string `json:"variant_id" validate:"required"`
}

type GetVariantWarehouseRacksResponse struct {
	WarehouseRacks []WarehouseRack `json:"warehouse_racks" validate:"required"`
}

type GetVariantWarehouseRackBatchesRequest struct {
	WarehouseRackId string `json:"warehouse_rack_id" validate:"required"`
	VariantId       string `json:"variant_id" validate:"required"`
}

type GetVariantWarehouseRackBatchesResponse struct {
	Batches []*string `json:"batches" validate:"required"`
}

type GetVariantWarehouseRackBatchExpiredDatesRequest struct {
	WarehouseRackId string `json:"warehouse_rack_id" validate:"required"`
	VariantId       string `json:"variant_id" validate:"required"`
	Batch           string `json:"batch"`
}

type GetVariantWarehouseRackBatchExpiredDatesResponse struct {
	ExpiredDates []*string `json:"expired_dates" validate:"required"`
}

type TransferHistory struct {
	FormNumber               string `json:"form_number" validate:"required"`
	TransactionDate          string `json:"transaction_date" validate:"required"`
	SourceWarehouseId        string `json:"source_warehouse_id" validate:"required"`
	SourceWarehouseName      string `json:"source_warehouse_name" validate:"required"`
	DestinationWarehouseId   string `json:"destination_warehouse_id" validate:"required"`
	DestinationWarehouseName string `json:"destination_warehouse_name" validate:"required"`
	ItemId                   string `json:"item_id" validate:"required"`
	ItemName                 string `json:"item_name" validate:"required"`
	ItemImageUrl             string `json:"item_image_url" validate:"required"`
	VariantId                string `json:"variant_id" validate:"required"`
	VariantName              string `json:"variant_name" validate:"required"`
	Amount                   string `json:"amount" validate:"required"`
}

type GetTransferHistoryRequest struct {
	WarehouseId string `json:"warehouse_id"`
	ItemId      string `json:"item_id"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
}

type GetTransferHistoryResponse struct {
	TransferHistories []TransferHistory `json:"transfer_histories" validate:"required"`
}
