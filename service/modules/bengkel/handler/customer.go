package handler

import (
	"github.com/DeniesKresna/bengkelgin/service/extensions/terror"
	"github.com/DeniesKresna/bengkelgin/types/models"
	"github.com/DeniesKresna/gohelper/utint"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func (h BengkelHandler) CustomerGetByID(ctx *gin.Context) {
	var (
		id   int64
		terr terror.ErrInterface
	)

	id = utint.Convert64FromString(ctx.Param("id"), 0)

	customer, terr := h.BengkelUsecase.CustomerGetByID(ctx, id)
	if terr != nil {
		ResponseJson(ctx, terr)
		return
	}
	ResponseJson(ctx, customer)
	return
}

func (h BengkelHandler) CustomerCreate(ctx *gin.Context) {
	var (
		customer models.Customer
		terr     terror.ErrInterface
	)

	if err := ctx.ShouldBindJSON(&customer); err != nil {
		terr = terror.ErrParameter(err.Error())
		ResponseJson(ctx, terr)
		return
	}

	customer, terr = h.BengkelUsecase.CustomerCreate(ctx, customer)
	if terr != nil {
		ResponseJson(ctx, terr)
		return
	}
	ResponseJson(ctx, customer)
	return
}

func (h BengkelHandler) CustomerUpdate(ctx *gin.Context) {
	var (
		customer models.Customer
		terr     terror.ErrInterface
	)

	if err := ctx.ShouldBindJSON(&customer); err != nil {
		terr = terror.ErrParameter(err.Error())
		ResponseJson(ctx, terr)
		return
	}

	customer, terr = h.BengkelUsecase.CustomerUpdate(ctx, customer)
	if terr != nil {
		ResponseJson(ctx, terr)
		return
	}
	ResponseJson(ctx, customer)
	return
}

func (h BengkelHandler) CustomerGetByPhone(ctx *gin.Context) {
	var (
		customer models.Customer
		terr     terror.ErrInterface
	)

	if err := ctx.ShouldBindBodyWith(&customer, binding.JSON); err != nil {
		terr = terror.ErrParameter(err.Error())
		ResponseJson(ctx, terr)
		return
	}

	customer, terr = h.BengkelUsecase.CustomerGetByPhone(ctx, customer.Phone)
	if terr != nil {
		ResponseJson(ctx, terr)
		return
	}
	ResponseJson(ctx, customer)
	return
}

func (h BengkelHandler) CustomerSearch(ctx *gin.Context) {
	var (
		search models.DbSearchObject
		terr   terror.ErrInterface
	)

	if err := ctx.ShouldBindJSON(&search); err != nil {
		terr = terror.ErrParameter(err.Error())
		ResponseJson(ctx, terr)
		return
	}

	res, terr := h.BengkelUsecase.CustomerSearch(ctx, search)
	if terr != nil {
		ResponseJson(ctx, terr)
		return
	}
	ResponseJson(ctx, res)
	return
}

func (h BengkelHandler) CustomerListGetByName(ctx *gin.Context) {
	var (
		nameReq models.NameRequest
		terr    terror.ErrInterface
	)

	if err := ctx.ShouldBindJSON(&nameReq); err != nil {
		terr = terror.ErrParameter(err.Error())
		ResponseJson(ctx, terr)
		return
	}

	customers, terr := h.BengkelUsecase.CustomerListGetByName(ctx, nameReq.Name)
	if terr != nil {
		ResponseJson(ctx, terr)
		return
	}
	ResponseJson(ctx, customers)
	return
}
