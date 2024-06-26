package repository

import (
	"errors"

	"github.com/DeniesKresna/skyshi1gin/service/extensions/helper"
	"github.com/DeniesKresna/skyshi1gin/service/extensions/terror"
	"github.com/DeniesKresna/skyshi1gin/types/constants"
	"github.com/DeniesKresna/skyshi1gin/types/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (r *WarehouseRepository) GetDB(ctx *gin.Context) (tx interface{}) {
	return r.db
}

func (r *WarehouseRepository) WarehouseGetByID(ctx *gin.Context, id int64) (warehouse models.Warehouse, terr terror.ErrInterface) {
	err := r.db.First(&warehouse, "id = ?", id).Error
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			terr = terror.ErrNotFoundData(err.Error())
			return
		}
		terr = terror.New(err)
	}
	return
}

func (r *WarehouseRepository) WarehouseList(ctx *gin.Context) (warehouses []models.Warehouse, terr terror.ErrInterface) {
	err := r.db.Find(&warehouses).Error
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			terr = terror.ErrNotFoundData(err.Error())
			return
		}
		terr = terror.New(err)
	}
	return
}

func (r *WarehouseRepository) WarehouseCreate(ctx *gin.Context, req models.Warehouse) (warehouse models.Warehouse, terr terror.ErrInterface) {
	err := r.db.Create(&req).Error
	if err != nil {
		terr = terror.New(err)
	}
	warehouse = req
	return
}

func (r *WarehouseRepository) WarehouseDelete(ctx *gin.Context, id int64) (terr terror.ErrInterface) {
	err := r.db.Delete(&models.Warehouse{}, "id = ?", id).Error
	if err != nil {
		terr = terror.New(err)
	}
	return
}

func (r *WarehouseRepository) WarehouseProductUpdateStock(ctx *gin.Context, req models.WarehouseProductStockRequest) (warehouseProduct models.WarehouseProduct, terr terror.ErrInterface) {
	tx := helper.TxGet(ctx)
	if tx == nil {
		tx = r.db
	}

	data := map[string]interface{}{
		"warehouse_id": req.WarehouseDestinyID,
		"product_id":   req.ProductID,
		"amount":       req.Amount,
	}
	err := tx.Where(models.WarehouseProduct{WarehouseID: req.WarehouseDestinyID, ProductID: req.ProductID}).
		Assign(data).FirstOrCreate(&warehouseProduct).Error

	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			terr = terror.ErrNotFoundData(err.Error())
			return
		}
		terr = terror.New(err)
	}
	return
}

func (r *WarehouseRepository) WarehouseProductLock(ctx *gin.Context, req models.WarehouseProduct) (warehouseProduct models.WarehouseProduct, terr terror.ErrInterface) {
	tx := helper.TxGet(ctx)
	if tx == nil {
		tx = r.db
	}

	err := tx.Clauses(clause.Locking{
		Strength: constants.TX_SHARE,
		Table:    clause.Table{Name: clause.CurrentTable},
	}).Where(req).First(&warehouseProduct).Error

	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return
		}
		terr = terror.New(err)
	}
	return
}

func (r *WarehouseRepository) WarehouseProductTotal(ctx *gin.Context, productID int64) (total int64, terr terror.ErrInterface) {
	err := r.db.Raw(`
		select sum(wp.amount) from warehouse_product wp
		join warehouses w on w.id = wp.warehouse_id
		where wp.product_id = ? and wp.deleted_at is null and w.deleted_at is null
		and w.active >= 1
	`, productID).Scan(&total).Error
	if err != nil {
		terr = terror.New(err)
	}
	return
}

func (r *WarehouseRepository) WarehouseGetProductList(ctx *gin.Context) (productList []models.Item, terr terror.ErrInterface) {
	err := r.db.Raw(`
		select product_id, sum(wp.amount) as 'amount' from warehouse_product wp
		join warehouses w on w.id = wp.warehouse_id
		where wp.deleted_at is null and w.deleted_at is null
		and w.active >= 1 group by product_id
	`).Scan(&productList).Error
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return
		}
		terr = terror.New(err)
	}
	return
}

func (r *WarehouseRepository) WarehouseUpdateActive(ctx *gin.Context, active bool, warehouseID int64) (terr terror.ErrInterface) {
	activeNumb := 0
	if active {
		activeNumb = 1
	}
	err := r.db.Model(&models.Warehouse{}).Where("id = ?", warehouseID).Update("active", activeNumb).Error
	if err != nil {
		terr = terror.New(err)
	}
	return
}
