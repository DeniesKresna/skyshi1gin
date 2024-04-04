package usecase

import (
	"fmt"

	"github.com/DeniesKresna/skyshi1gin/service/extensions/helper"
	"github.com/DeniesKresna/skyshi1gin/service/extensions/terror"
	"github.com/DeniesKresna/skyshi1gin/types/models"
	"github.com/gin-gonic/gin"
)

func (u OrderUsecase) OrderList(ctx *gin.Context) (order []models.Order, terr terror.ErrInterface) {
	order, terr = u.orderRepo.OrderList(ctx)

	return
}

func (u OrderUsecase) OrderGetByID(ctx *gin.Context, id int64) (order models.Order, terr terror.ErrInterface) {
	order, terr = u.orderRepo.OrderGetByID(ctx, id)

	return
}

func (u OrderUsecase) OrderItem(ctx *gin.Context, req []models.Item) (order models.Order, terr terror.ErrInterface) {
	helper.TxCreate(ctx, u.orderRepo.GetDB)
	defer func() {
		if terr != nil {
			helper.TxRollBack(ctx)
		} else {
			helper.TxCommit(ctx)
		}
	}()

	for _, v := range req {
		_, terr = u.orderRepo.WarehouseProductLock(ctx, models.WarehouseProduct{ProductID: v.ProductID})
		if terr != nil {
			return
		}

		var pd models.Product
		pd, terr = u.orderRepo.ProductGetByID(ctx, v.ProductID)
		if terr != nil {
			return
		}

		var total int64
		total, terr = u.orderRepo.WarehouseProductTotal(ctx, v.ProductID)
		if terr != nil {
			return
		}

		if total < v.Amount {
			terr = terror.ErrInvalidRule(fmt.Sprintf("Product %s is out of stock", pd.Name))
		}
	}

	return
}
