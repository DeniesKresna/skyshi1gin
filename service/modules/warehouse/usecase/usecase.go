package usecase

import (
	"github.com/DeniesKresna/skyshi1gin/service/extensions/terror"
	"github.com/DeniesKresna/skyshi1gin/service/modules/warehouse/repository"
	"github.com/DeniesKresna/skyshi1gin/types/models"
	"github.com/gin-gonic/gin"
)

type WarehouseUsecase struct {
	warehouseRepo repository.IRepository
}

func WarehouseCreateUsecase(repo repository.IRepository) IUsecase {
	warehouseRepo := WarehouseUsecase{
		warehouseRepo: repo,
	}
	return warehouseRepo
}

type IUsecase interface {
	WarehouseList(ctx *gin.Context) (warehouse []models.Warehouse, terr terror.ErrInterface)
	WarehouseGetByID(ctx *gin.Context, id int64) (warehouse models.Warehouse, terr terror.ErrInterface)
	WarehouseCreate(ctx *gin.Context, req models.Warehouse) (warehouse models.Warehouse, terr terror.ErrInterface)
	WarehouseProductUpdateStock(ctx *gin.Context, req []models.WarehouseProductStockRequest) (warehouseProducts []models.WarehouseProduct, terr terror.ErrInterface)
	WarehouseProductTransfer(ctx *gin.Context, req models.WarehouseTransferRequest) (warehouseProducts []models.WarehouseProduct, terr terror.ErrInterface)
	WarehouseProductTotal(ctx *gin.Context, productID int64) (total int64, terr terror.ErrInterface)
}
