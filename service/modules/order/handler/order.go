package handler

import (
	"github.com/DeniesKresna/gohelper/utint"
	"github.com/DeniesKresna/skyshi1gin/service/extensions/terror"
	"github.com/DeniesKresna/skyshi1gin/types/models"
	"github.com/gin-gonic/gin"
)

func (h OrderHandler) OrderGetByID(ctx *gin.Context) {
	var (
		id   int64
		terr terror.ErrInterface
	)

	id = utint.Convert64FromString(ctx.Param("id"), 0)

	order, terr := h.orderUsecase.OrderGetByID(ctx, id)
	if terr != nil {
		ResponseJson(ctx, terr)
		return
	}
	ResponseJson(ctx, order)
}

func (h OrderHandler) OrderList(ctx *gin.Context) {
	var (
		terr terror.ErrInterface
	)

	order, terr := h.orderUsecase.OrderList(ctx)
	if terr != nil {
		ResponseJson(ctx, terr)
		return
	}
	ResponseJson(ctx, order)
}

func (h OrderHandler) OrderItem(ctx *gin.Context) {
	var (
		terr terror.ErrInterface
		req  []models.Item
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		terr = terror.ErrParameter(err.Error())
		ResponseJson(ctx, terr)
		return
	}

	order, terr := h.orderUsecase.OrderItem(ctx, req)
	if terr != nil {
		ResponseJson(ctx, terr)
		return
	}
	ResponseJson(ctx, order)
}
