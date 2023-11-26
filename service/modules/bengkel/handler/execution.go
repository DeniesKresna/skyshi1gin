package handler

import (
	"errors"

	"github.com/DeniesKresna/bengkelgin/service/extensions/terror"
	"github.com/DeniesKresna/bengkelgin/types/models"
	"github.com/DeniesKresna/gohelper/utint"
	"github.com/gin-gonic/gin"
)

func (h BengkelHandler) ExecutionGetByID(ctx *gin.Context) {
	var (
		id   int64
		terr terror.ErrInterface
	)

	id = utint.Convert64FromString(ctx.Param("id"), 0)

	execution, terr := h.BengkelUsecase.ExecutionGetByID(ctx, id)
	if terr != nil {
		ResponseJson(ctx, terr)
		return
	}
	ResponseJson(ctx, execution)
	return
}

func (h BengkelHandler) ExecutionCreate(ctx *gin.Context) {
	var (
		execution models.Execution
		terr      terror.ErrInterface
	)

	if err := ctx.ShouldBindJSON(&execution); err != nil {
		terr = terror.ErrParameter(err.Error())
		ResponseJson(ctx, terr)
		return
	}

	execution, terr = h.BengkelUsecase.ExecutionCreate(ctx, execution)
	if terr != nil {
		ResponseJson(ctx, terr)
		return
	}
	ResponseJson(ctx, execution)
	return
}

func (h BengkelHandler) ExecutionUpdate(ctx *gin.Context) {
	var (
		execution models.Execution
		terr      terror.ErrInterface
	)

	if err := ctx.ShouldBindJSON(&execution); err != nil {
		terr = terror.ErrParameter(err.Error())
		ResponseJson(ctx, terr)
		return
	}

	execution, terr = h.BengkelUsecase.ExecutionUpdate(ctx, execution)
	if terr != nil {
		ResponseJson(ctx, terr)
		return
	}
	ResponseJson(ctx, execution)
	return
}

func (h BengkelHandler) ExecutionSearch(ctx *gin.Context) {
	var (
		search models.DbSearchObject
		terr   terror.ErrInterface
	)

	if err := ctx.ShouldBindJSON(&search); err != nil {
		terr = terror.ErrParameter(err.Error())
		ResponseJson(ctx, terr)
		return
	}

	res, terr := h.BengkelUsecase.ExecutionSearch(ctx, search)
	if terr != nil {
		ResponseJson(ctx, terr)
		return
	}
	ResponseJson(ctx, res)
	return
}

func (h BengkelHandler) ExecutionDownload(ctx *gin.Context) {
	var (
		search models.DbSearchObject
		terr   terror.ErrInterface
	)

	if err := ctx.ShouldBindJSON(&search); err != nil {
		terr = terror.ErrParameter(err.Error())
		ResponseJson(ctx, terr)
		return
	}

	res, terr := h.BengkelUsecase.ExecutionDownload(ctx, search)
	if terr != nil {
		ResponseJson(ctx, terr)
		return
	}

	if res == "" {
		terr = terror.New(errors.New("Failed download excel"))
		ResponseJson(ctx, terr)
		return
	}

	// Set the response headers
	ResponseJson(ctx, res)
	return
}
