package repository

import (
	"github.com/DeniesKresna/skyshi1gin/service/extensions/terror"
	"github.com/DeniesKresna/skyshi1gin/types/models"
	"github.com/gin-gonic/gin"
)

func (r *OrderRepository) WarehouseGetByID(ctx *gin.Context, id int64) (warehouse models.Warehouse, terr terror.ErrInterface) {
	return r.warehouseCross.WarehouseGetByID(ctx, id)
}

func (h OrderRepository) WarehouseProductTotal(ctx *gin.Context, productID int64) (total int64, terr terror.ErrInterface) {
	return h.warehouseCross.WarehouseProductTotal(ctx, productID)
}

func (h OrderRepository) WarehouseProductLock(ctx *gin.Context, req models.WarehouseProduct) (warehouseProduct models.WarehouseProduct, terr terror.ErrInterface) {
	return h.warehouseCross.WarehouseProductLock(ctx, req)
}
