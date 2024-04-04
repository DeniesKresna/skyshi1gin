package repository

import (
	"github.com/DeniesKresna/skyshi1gin/service/extensions/terror"
	productCross "github.com/DeniesKresna/skyshi1gin/service/modules/product/handler/cross"
	userCross "github.com/DeniesKresna/skyshi1gin/service/modules/user/handler/cross"
	warehouseCross "github.com/DeniesKresna/skyshi1gin/service/modules/warehouse/handler/cross"
	"github.com/DeniesKresna/skyshi1gin/types/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrderRepository struct {
	db             *gorm.DB
	warehouseCross warehouseCross.WarehouseCross
	productCross   productCross.ProductCross
	userCross      userCross.UserCross
}

func OrderCreateRepository(db *gorm.DB, warehouseCross warehouseCross.WarehouseCross, productCross productCross.ProductCross, userCross userCross.UserCross) IRepository {
	orderRepository := OrderRepository{
		db:             db,
		warehouseCross: warehouseCross,
		productCross:   productCross,
		userCross:      userCross,
	}
	return &orderRepository
}

type IRepository interface {
	GetDB(ctx *gin.Context) (tx interface{})
	OrderList(ctx *gin.Context) (order []models.Order, terr terror.ErrInterface)
	OrderGetByID(ctx *gin.Context, id int64) (order models.Order, terr terror.ErrInterface)
	OrderCreate(ctx *gin.Context, req models.Order) (order models.Order, terr terror.ErrInterface)
	PaymentGetByCode(ctx *gin.Context, code string) (payment models.Payment, terr terror.ErrInterface)
	PaymentCreate(ctx *gin.Context, req models.Payment) (payment models.Payment, terr terror.ErrInterface)

	// warehouse cross
	WarehouseGetByID(ctx *gin.Context, id int64) (warehouse models.Warehouse, terr terror.ErrInterface)
	WarehouseProductTotal(ctx *gin.Context, productID int64) (total int64, terr terror.ErrInterface)
	WarehouseProductLock(ctx *gin.Context, req models.WarehouseProduct) (warehouseProduct models.WarehouseProduct, terr terror.ErrInterface)

	// product cross
	ProductGetByID(ctx *gin.Context, id int64) (product models.Product, terr terror.ErrInterface)

	// user cross
	AuthGetFromContext(ctx *gin.Context) (userRole models.UserRole, terr terror.ErrInterface)
}
