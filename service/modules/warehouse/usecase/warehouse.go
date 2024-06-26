package usecase

import (
	"fmt"

	"github.com/DeniesKresna/skyshi1gin/service/extensions/helper"
	"github.com/DeniesKresna/skyshi1gin/service/extensions/terror"
	"github.com/DeniesKresna/skyshi1gin/types/models"
	"github.com/gin-gonic/gin"
)

func (u WarehouseUsecase) WarehouseList(ctx *gin.Context) (warehouse []models.Warehouse, terr terror.ErrInterface) {
	warehouse, terr = u.warehouseRepo.WarehouseList(ctx)

	return
}

func (u WarehouseUsecase) WarehouseGetByID(ctx *gin.Context, id int64) (warehouse models.Warehouse, terr terror.ErrInterface) {
	warehouse, terr = u.warehouseRepo.WarehouseGetByID(ctx, id)

	return
}

func (u WarehouseUsecase) WarehouseCreate(ctx *gin.Context, req models.Warehouse) (warehouse models.Warehouse, terr terror.ErrInterface) {
	req.Active = 1
	warehouse, terr = u.warehouseRepo.WarehouseCreate(ctx, req)

	return
}

func (u WarehouseUsecase) WarehouseProductUpdateStock(ctx *gin.Context, req []models.WarehouseProductStockRequest) (warehouseProducts []models.WarehouseProduct, terr terror.ErrInterface) {
	helper.TxCreate(ctx, u.warehouseRepo.GetDB)
	defer func() {
		if terr != nil {
			helper.TxRollBack(ctx)
		} else {
			helper.TxCommit(ctx)
		}
	}()

	for _, v := range req {
		var (
			wh models.Warehouse
			pd models.Product
		)
		wh, terr = u.WarehouseGetByID(ctx, v.WarehouseDestinyID)
		if terr != nil {
			return
		}

		if wh.Active < 1 {
			terr = terror.ErrInvalidRule(fmt.Sprintf("Warehouse %s is inactive", wh.Name))
			return
		}

		pd, terr = u.warehouseRepo.ProductGetByID(ctx, v.ProductID)
		if terr != nil {
			return
		}

		var wpLock models.WarehouseProduct
		wpLock, terr = u.warehouseRepo.WarehouseProductLock(ctx, models.WarehouseProduct{
			WarehouseID: v.WarehouseDestinyID,
			ProductID:   v.ProductID,
		})
		if terr != nil {
			return
		}
		{
			amount := wpLock.Amount
			amount += v.Amount

			if amount < 0 {
				terr = terror.ErrInvalidRule(fmt.Sprintf("Product %s in warehouse %s out of stock", pd.Name, wh.Name))
				return
			}

			v.Amount = amount

			var updatedWP models.WarehouseProduct
			updatedWP, terr = u.warehouseRepo.WarehouseProductUpdateStock(ctx, v)
			if terr != nil {
				return
			}
			warehouseProducts = append(warehouseProducts, updatedWP)
		}
	}

	return
}

func (u WarehouseUsecase) WarehouseProductTransfer(ctx *gin.Context, req models.WarehouseTransferRequest) (warehouseProducts []models.WarehouseProduct, terr terror.ErrInterface) {
	if req.WarehouseDestinyID == req.WarehouseSenderID {
		terr = terror.ErrParameter("Warehouse annot transfer to itself")
		return
	}

	if req.Amount < 1 {
		terr = terror.ErrParameter("Amount to be transfered should be more than 1")
		return
	}

	pd, terr := u.warehouseRepo.ProductGetByID(ctx, req.ProductID)
	if terr != nil {
		return
	}

	whSource, terr := u.WarehouseGetByID(ctx, req.WarehouseSenderID)
	if terr != nil {
		return
	}

	_, terr = u.WarehouseGetByID(ctx, req.WarehouseDestinyID)
	if terr != nil {
		return
	}

	helper.TxCreate(ctx, u.warehouseRepo.GetDB)
	defer func() {
		if terr != nil {
			helper.TxRollBack(ctx)
		} else {
			helper.TxCommit(ctx)
		}
	}()

	wpLockSource, terr := u.warehouseRepo.WarehouseProductLock(ctx, models.WarehouseProduct{
		WarehouseID: req.WarehouseSenderID,
		ProductID:   req.ProductID,
	})
	if terr != nil {
		return
	}

	if wpLockSource.Amount < req.Amount {
		terr = terror.ErrInvalidRule(fmt.Sprintf("Product %s in warehouse %s out of stock", pd.Name, whSource.Name))
		return
	}

	_, terr = u.warehouseRepo.WarehouseProductLock(ctx, models.WarehouseProduct{
		WarehouseID: req.WarehouseDestinyID,
		ProductID:   req.ProductID,
	})
	if terr != nil {
		return
	}

	res, terr := u.warehouseRepo.WarehouseProductUpdateStock(ctx, models.WarehouseProductStockRequest{
		WarehouseDestinyID: req.WarehouseSenderID,
		ProductID:          req.ProductID,
		Amount:             wpLockSource.Amount - req.Amount,
	})
	if terr != nil {
		return
	}

	warehouseProducts = append(warehouseProducts, res)

	res, terr = u.warehouseRepo.WarehouseProductUpdateStock(ctx, models.WarehouseProductStockRequest{
		WarehouseDestinyID: req.WarehouseDestinyID,
		ProductID:          req.ProductID,
		Amount:             req.Amount,
	})
	if terr != nil {
		return
	}

	warehouseProducts = append(warehouseProducts, res)

	return
}

func (u WarehouseUsecase) WarehouseProductTotal(ctx *gin.Context, productID int64) (total int64, terr terror.ErrInterface) {
	return u.warehouseRepo.WarehouseProductTotal(ctx, productID)
}

func (u WarehouseUsecase) WarehouseProductLock(ctx *gin.Context, req models.WarehouseProduct) (warehouseProduct models.WarehouseProduct, terr terror.ErrInterface) {
	return u.warehouseRepo.WarehouseProductLock(ctx, req)
}

func (u WarehouseUsecase) WarehouseGetAllProductList(ctx *gin.Context) (warehouseProducts []models.AllWarehouseProduct, terr terror.ErrInterface) {
	products, terr := u.warehouseRepo.ProductList(ctx)
	if terr != nil {
		return
	}
	var productsObj = make(map[int64]models.Product)
	for _, v := range products {
		productsObj[int64(v.ID)] = v
	}

	allWarehouseItems, terr := u.warehouseRepo.WarehouseGetProductList(ctx)
	if terr != nil {
		return
	}

	var productAmount = make(map[int64]int64)
	for _, v := range allWarehouseItems {
		productAmount[v.ProductID] = v.Amount
	}

	for _, v := range products {
		amount, ok := productAmount[int64(v.ID)]
		if !ok {
			amount = 0
		}
		wp := models.AllWarehouseProduct{
			Product: productsObj[int64(v.ID)],
			Amount:  amount,
		}
		warehouseProducts = append(warehouseProducts, wp)
	}

	return
}

func (u WarehouseUsecase) WarehouseUpdateActive(ctx *gin.Context, warehouseID int64) (warehouse models.Warehouse, terr terror.ErrInterface) {
	_, terr = u.warehouseRepo.WarehouseGetByID(ctx, warehouseID)
	if terr != nil {
		return
	}

	terr = u.warehouseRepo.WarehouseUpdateActive(ctx, true, warehouseID)
	if terr != nil {
		return
	}

	warehouse, terr = u.warehouseRepo.WarehouseGetByID(ctx, warehouseID)
	if terr != nil {
		return
	}
	return
}

func (u WarehouseUsecase) WarehouseUpdateInactive(ctx *gin.Context, warehouseID int64) (warehouse models.Warehouse, terr terror.ErrInterface) {
	_, terr = u.warehouseRepo.WarehouseGetByID(ctx, warehouseID)
	if terr != nil {
		return
	}

	terr = u.warehouseRepo.WarehouseUpdateActive(ctx, false, warehouseID)
	if terr != nil {
		return
	}

	warehouse, terr = u.warehouseRepo.WarehouseGetByID(ctx, warehouseID)
	if terr != nil {
		return
	}
	return
}

func (u WarehouseUsecase) WarehouseCustomerBuy(ctx *gin.Context, order models.Order) (warehouseProducts []models.WarehouseProduct, terr terror.ErrInterface) {
	whs, terr := u.warehouseRepo.WarehouseList(ctx)
	if terr != nil {
		return
	}

	var pd models.Product
	pd, terr = u.warehouseRepo.ProductGetByID(ctx, order.PaymentID)
	if terr != nil {
		return
	}

	var productDeductMap = make(map[int64]int64)
	remaining := order.Amount

	for _, wh := range whs {
		var wp models.WarehouseProduct
		wp, terr = u.warehouseRepo.WarehouseProductLock(ctx, models.WarehouseProduct{ProductID: order.ProductID, WarehouseID: int64(wh.ID)})
		if terr != nil {
			return
		}

		for {
			if remaining <= 0 {
				break
			}
			if wp.Amount >= remaining {
				if amt, ok := productDeductMap[int64(wh.ID)]; !ok {
					productDeductMap[int64(wh.ID)] = remaining
				} else {
					productDeductMap[int64(wh.ID)] = remaining + amt
				}
				remaining -= remaining
			} else {
				if amt, ok := productDeductMap[int64(wh.ID)]; !ok {
					productDeductMap[int64(wh.ID)] = wp.Amount
				} else {
					productDeductMap[int64(wh.ID)] = wp.Amount + amt
				}
				remaining -= wp.Amount
			}
		}
	}

	if remaining < 0 {
		terr = terror.ErrInvalidRule(fmt.Sprintf("Product %s is out of stock in shop", pd.Name))
		return
	}

	var wpsrs []models.WarehouseProductStockRequest
	for k, v := range productDeductMap {
		wpsr := models.WarehouseProductStockRequest{
			WarehouseDestinyID: k,
			ProductID:          int64(pd.ID),
			Amount:             -1 * v,
		}
		wpsrs = append(wpsrs, wpsr)
	}

	return u.WarehouseProductUpdateStock(ctx, wpsrs)
}
