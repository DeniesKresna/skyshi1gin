package repository

import (
	"github.com/DeniesKresna/skyshi1gin/service/extensions/terror"
	productCross "github.com/DeniesKresna/skyshi1gin/service/modules/product/handler/cross"
	"github.com/DeniesKresna/skyshi1gin/types/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type WarehouseRepository struct {
	db           *gorm.DB
	productCross productCross.ProductCross
}

func WarehouseCreateRepository(db *gorm.DB, productCross productCross.ProductCross) IRepository {
	warehouseRepository := WarehouseRepository{
		db:           db,
		productCross: productCross,
	}
	return &warehouseRepository
}

type IRepository interface {
	GetDB(ctx *gin.Context) (tx interface{})
	WarehouseGetByID(ctx *gin.Context, id int64) (warehouse models.Warehouse, terr terror.ErrInterface)
	WarehouseList(ctx *gin.Context) (warehouse []models.Warehouse, terr terror.ErrInterface)
	WarehouseCreate(ctx *gin.Context, req models.Warehouse) (warehouse models.Warehouse, terr terror.ErrInterface)
	WarehouseDelete(ctx *gin.Context, id int64) (terr terror.ErrInterface)
	WarehouseProductUpdateStock(ctx *gin.Context, req models.WarehouseProductStockRequest) (warehouseProduct models.WarehouseProduct, terr terror.ErrInterface)
	WarehouseProductLock(ctx *gin.Context, req models.WarehouseProduct) (warehouseProduct models.WarehouseProduct, terr terror.ErrInterface)
	WarehouseProductTotal(ctx *gin.Context, productID int64) (total int64, terr terror.ErrInterface)
	WarehouseGetProductList(ctx *gin.Context) (productList []models.Item, terr terror.ErrInterface)
	WarehouseUpdateActive(ctx *gin.Context, active bool, warehouseID int64) (terr terror.ErrInterface)

	// product cross module
	ProductGetByID(ctx *gin.Context, id int64) (product models.Product, terr terror.ErrInterface)
	ProductList(ctx *gin.Context) (product []models.Product, terr terror.ErrInterface)
}
