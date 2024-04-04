package handler

import (
	"github.com/DeniesKresna/gohelper/utint"
	"github.com/DeniesKresna/skyshi1gin/service/extensions/terror"
	"github.com/DeniesKresna/skyshi1gin/types/models"
	"github.com/gin-gonic/gin"
)

func (h WarehouseHandler) WarehouseGetByID(ctx *gin.Context) {
	var (
		id   int64
		terr terror.ErrInterface
	)

	id = utint.Convert64FromString(ctx.Param("id"), 0)

	warehouse, terr := h.warehouseUsecase.WarehouseGetByID(ctx, id)
	if terr != nil {
		ResponseJson(ctx, terr)
		return
	}
	ResponseJson(ctx, warehouse)
}

func (h WarehouseHandler) WarehouseList(ctx *gin.Context) {
	var (
		terr terror.ErrInterface
	)

	warehouse, terr := h.warehouseUsecase.WarehouseList(ctx)
	if terr != nil {
		ResponseJson(ctx, terr)
		return
	}
	ResponseJson(ctx, warehouse)
}

func (h WarehouseHandler) WarehouseCreate(ctx *gin.Context) {
	var (
		terr terror.ErrInterface
		req  models.Warehouse
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		terr = terror.ErrParameter(err.Error())
		ResponseJson(ctx, terr)
		return
	}

	warehouse, terr := h.warehouseUsecase.WarehouseCreate(ctx, req)
	if terr != nil {
		ResponseJson(ctx, terr)
		return
	}
	ResponseJson(ctx, warehouse)
}

func (h WarehouseHandler) WarehouseProductUpdateStock(ctx *gin.Context) {
	var (
		terr terror.ErrInterface
		req  []models.WarehouseProductStockRequest
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		terr = terror.ErrParameter(err.Error())
		ResponseJson(ctx, terr)
		return
	}

	warehouseProduct, terr := h.warehouseUsecase.WarehouseProductUpdateStock(ctx, req)
	if terr != nil {
		ResponseJson(ctx, terr)
		return
	}
	ResponseJson(ctx, warehouseProduct)
}

func (h WarehouseHandler) WarehouseProductTransfer(ctx *gin.Context) {
	var (
		terr terror.ErrInterface
		req  models.WarehouseTransferRequest
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		terr = terror.ErrParameter(err.Error())
		ResponseJson(ctx, terr)
		return
	}

	warehouseProduct, terr := h.warehouseUsecase.WarehouseProductTransfer(ctx, req)
	if terr != nil {
		ResponseJson(ctx, terr)
		return
	}
	ResponseJson(ctx, warehouseProduct)
}
