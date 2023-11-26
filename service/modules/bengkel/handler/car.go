package handler

import (
	"github.com/DeniesKresna/bengkelgin/service/extensions/terror"
	"github.com/DeniesKresna/bengkelgin/types/models"
	"github.com/DeniesKresna/gohelper/utint"
	"github.com/gin-gonic/gin"
)

func (h BengkelHandler) CarGetByID(ctx *gin.Context) {
	var (
		id   int64
		terr terror.ErrInterface
	)

	id = utint.Convert64FromString(ctx.Param("id"), 0)

	car, terr := h.BengkelUsecase.CarGetByID(ctx, id)
	if terr != nil {
		ResponseJson(ctx, terr)
		return
	}
	ResponseJson(ctx, car)
	return
}

func (h BengkelHandler) CarCreate(ctx *gin.Context) {
	var (
		car  models.Car
		terr terror.ErrInterface
	)

	if err := ctx.ShouldBindJSON(&car); err != nil {
		terr = terror.ErrParameter(err.Error())
		ResponseJson(ctx, terr)
		return
	}

	car, terr = h.BengkelUsecase.CarCreate(ctx, car)
	if terr != nil {
		ResponseJson(ctx, terr)
		return
	}
	ResponseJson(ctx, car)
	return
}

func (h BengkelHandler) CarUpdate(ctx *gin.Context) {
	var (
		car  models.Car
		terr terror.ErrInterface
	)

	if err := ctx.ShouldBindJSON(&car); err != nil {
		terr = terror.ErrParameter(err.Error())
		ResponseJson(ctx, terr)
		return
	}

	car, terr = h.BengkelUsecase.CarUpdate(ctx, car)
	if terr != nil {
		ResponseJson(ctx, terr)
		return
	}
	ResponseJson(ctx, car)
	return
}

func (h BengkelHandler) CarGetByPlat(ctx *gin.Context) {
	var (
		carReq models.CarByPlatRequest
		terr   terror.ErrInterface
	)

	if err := ctx.ShouldBindJSON(&carReq); err != nil {
		terr = terror.ErrParameter(err.Error())
		ResponseJson(ctx, terr)
		return
	}

	car, terr := h.BengkelUsecase.CarGetByPlat(ctx, carReq.Plat)
	if terr != nil {
		ResponseJson(ctx, terr)
		return
	}
	ResponseJson(ctx, car)
	return
}

func (h BengkelHandler) CarListGetByPlat(ctx *gin.Context) {
	var (
		carReq models.CarByPlatRequest
		terr   terror.ErrInterface
	)

	if err := ctx.ShouldBindJSON(&carReq); err != nil {
		terr = terror.ErrParameter(err.Error())
		ResponseJson(ctx, terr)
		return
	}

	cars, terr := h.BengkelUsecase.CarListGetByPlat(ctx, carReq.Plat)
	if terr != nil {
		ResponseJson(ctx, terr)
		return
	}
	ResponseJson(ctx, cars)
	return
}

func (h BengkelHandler) CarSearch(ctx *gin.Context) {
	var (
		search models.DbSearchObject
		terr   terror.ErrInterface
	)

	if err := ctx.ShouldBindJSON(&search); err != nil {
		terr = terror.ErrParameter(err.Error())
		ResponseJson(ctx, terr)
		return
	}

	res, terr := h.BengkelUsecase.CarSearch(ctx, search)
	if terr != nil {
		ResponseJson(ctx, terr)
		return
	}
	ResponseJson(ctx, res)
	return
}
