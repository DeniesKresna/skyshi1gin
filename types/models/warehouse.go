package models

type Warehouse struct {
	DbStandard
	Name   string `json:"name" binding:"required"`
	Active int8   `json:"active"`
}

type WarehouseProduct struct {
	DbStandard
	WarehouseID int64 `json:"warehouse_id"`
	ProductID   int64 `json:"product_id"`
	Amount      int64 `json:"amount"`
}

type AllWarehouseProduct struct {
	Product Product `json:"product"`
	Amount  int64   `json:"amount"`
}

func (wp *WarehouseProduct) TableName() string {
	return "warehouse_product"
}

type WarehouseTransferRequest struct {
	WarehouseSenderID  int64 `json:"warehouse_sender_id"`
	WarehouseDestinyID int64 `json:"warehouse_destiny_id"`
	ProductID          int64 `json:"product_id"`
	Amount             int64 `json:"amount"`
}

type WarehouseProductStockRequest struct {
	WarehouseDestinyID int64 `json:"warehouse_destiny_id"`
	ProductID          int64 `json:"product_id"`
	Amount             int64 `json:"amount"`
}
