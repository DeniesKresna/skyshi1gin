package cross

import (
	"github.com/DeniesKresna/skyshi1gin/config"
	"github.com/DeniesKresna/skyshi1gin/service/extensions/terror"
	productCrossHandler "github.com/DeniesKresna/skyshi1gin/service/modules/product/handler/cross"
	"github.com/DeniesKresna/skyshi1gin/service/modules/warehouse/repository"
	"github.com/DeniesKresna/skyshi1gin/service/modules/warehouse/usecase"
	"github.com/DeniesKresna/skyshi1gin/types/models"
	"github.com/gin-gonic/gin"
)

type WarehouseCross struct {
	warehouseUsecase usecase.IUsecase
}

func WarehouseCreateCross(cfg *config.Config) WarehouseCross {
	productCross := productCrossHandler.ProductCreateCross(cfg)
	repo := repository.WarehouseCreateRepository(cfg.DB, productCross)
	warehouseUsecase := usecase.WarehouseCreateUsecase(repo)
	return WarehouseCross{
		warehouseUsecase: warehouseUsecase,
	}
}

func (h WarehouseCross) WarehouseGetByID(ctx *gin.Context, id int64) (res models.Warehouse, terr terror.ErrInterface) {
	return h.warehouseUsecase.WarehouseGetByID(ctx, id)
}

func (h WarehouseCross) WarehouseProductTotal(ctx *gin.Context, productID int64) (total int64, terr terror.ErrInterface) {
	return h.warehouseUsecase.WarehouseProductTotal(ctx, productID)
}

func (h WarehouseCross) WarehouseProductLock(ctx *gin.Context, req models.WarehouseProduct) (warehouseProduct models.WarehouseProduct, terr terror.ErrInterface) {
	return h.warehouseUsecase.WarehouseProductLock(ctx, req)
}

func (h WarehouseCross) WarehouseCustomerBuy(ctx *gin.Context, order models.Order) (warehouseProducts []models.WarehouseProduct, terr terror.ErrInterface) {
	return h.warehouseUsecase.WarehouseCustomerBuy(ctx, order)
}
