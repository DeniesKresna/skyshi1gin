package handler

import (
	"github.com/DeniesKresna/gohelper/utint"
	"github.com/DeniesKresna/skyshi1gin/service/extensions/terror"
	"github.com/DeniesKresna/skyshi1gin/types/models"
	"github.com/gin-gonic/gin"
)

func (h ProductHandler) ProductGetByID(ctx *gin.Context) {
	var (
		id   int64
		terr terror.ErrInterface
	)

	id = utint.Convert64FromString(ctx.Param("id"), 0)

	product, terr := h.productUsecase.ProductGetByID(ctx, id)
	if terr != nil {
		ResponseJson(ctx, terr)
		return
	}
	ResponseJson(ctx, product)
}

func (h ProductHandler) ProductList(ctx *gin.Context) {
	var (
		terr terror.ErrInterface
	)

	product, terr := h.productUsecase.ProductList(ctx)
	if terr != nil {
		ResponseJson(ctx, terr)
		return
	}
	ResponseJson(ctx, product)
}

func (h ProductHandler) ProductCreate(ctx *gin.Context) {
	var (
		terr terror.ErrInterface
		req  models.Product
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		terr = terror.ErrParameter(err.Error())
		ResponseJson(ctx, terr)
		return
	}

	product, terr := h.productUsecase.ProductCreate(ctx, req)
	if terr != nil {
		ResponseJson(ctx, terr)
		return
	}
	ResponseJson(ctx, product)
}
