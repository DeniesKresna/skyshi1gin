package usecase

import (
	"errors"
	"fmt"

	"github.com/DeniesKresna/bengkelgin/service/extensions/helper"
	"github.com/DeniesKresna/bengkelgin/service/extensions/terror"
	"github.com/DeniesKresna/bengkelgin/types/constants"
	"github.com/DeniesKresna/bengkelgin/types/models"
	"github.com/gin-gonic/gin"
)

func (u BengkelUsecase) ExecutionGetByID(ctx *gin.Context, id int64) (user models.ExecutionSearchResponse, terr terror.ErrInterface) {
	return u.bengkelRepo.ExecutionGetByID(ctx, id)
}

func (u BengkelUsecase) ExecutionCreate(ctx *gin.Context, execution models.Execution) (executionRes models.Execution, terr terror.ErrInterface) {
	operator, terr := u.userUsecase.AuthGetFromContext(ctx)
	if terr != nil {
		return
	}

	execution.UpdatedBy = operator.Name

	terr = u.bengkelRepo.ExecutionCreate(ctx, &execution)
	if terr != nil {
		return
	}

	executionRes = execution
	return
}

func (u BengkelUsecase) ExecutionUpdate(ctx *gin.Context, execution models.Execution) (executionRes models.Execution, terr terror.ErrInterface) {
	operator, terr := u.userUsecase.AuthGetFromContext(ctx)
	if terr != nil {
		return
	}

	execution.UpdatedBy = operator.Name

	terr = u.bengkelRepo.ExecutionUpdate(ctx, &execution)
	if terr != nil {
		return
	}

	executionRes = execution
	return
}

func (u BengkelUsecase) ExecutionSearch(ctx *gin.Context, filter models.DbSearchObject) (res models.DbSearchObject, terr terror.ErrInterface) {
	defer func() {
		if r := recover(); r != nil {
			terr = terror.New(errors.New(fmt.Sprint(r)))
		}
	}()

	filter.Mode = constants.DB_MODE_PAGE
	execution := models.ExecutionFilter{
		ExecuterName:  filter.PayloadData["executer_name"].(string),
		CarPlat:       filter.PayloadData["car_plat"].(string),
		CustomerName:  filter.PayloadData["customer_name"].(string),
		CustomerPhone: filter.PayloadData["customer_phone"].(string),
		CustomerEmail: filter.PayloadData["customer_email"].(string),
		Price:         helper.InterfaceSliceToSliceInt64(filter.PayloadData["price"]),
		Paid:          helper.InterfaceSliceToSliceInt64(filter.PayloadData["paid"]),
		PaidOff:       helper.InterfacePointerBoolToPointerBool(filter.PayloadData["paid_off"]),
		PaidFrom:      helper.InterfacePointerTimeToPointerTime(filter.PayloadData["paid_from"]),
		PaidTo:        helper.InterfacePointerTimeToPointerTime(filter.PayloadData["paid_to"]),
		ExecutedFrom:  helper.InterfacePointerTimeToPointerTime(filter.PayloadData["executed_from"]),
		ExecutedTo:    helper.InterfacePointerTimeToPointerTime(filter.PayloadData["executed_to"]),
		FinishFrom:    helper.InterfacePointerTimeToPointerTime(filter.PayloadData["finish_from"]),
		FinishTo:      helper.InterfacePointerTimeToPointerTime(filter.PayloadData["finish_to"]),
	}

	err := helper.WrapPercentOnStructString(&execution)
	if err != nil {
		terr = terror.New(err)
		return
	}

	executionsRes, totalData, terr := u.bengkelRepo.ExecutionSearch(ctx, execution, filter)
	if terr != nil {
		return
	}

	filter.ResponseData = executionsRes
	filter.TotalData = totalData
	res = filter

	return
}

func (u BengkelUsecase) ExecutionDownload(ctx *gin.Context, filter models.DbSearchObject) (res string, terr terror.ErrInterface) {
	defer func() {
		if r := recover(); r != nil {
			terr = terror.New(errors.New(fmt.Sprint(r)))
		}
	}()

	filter.Mode = constants.DB_MODE_DOWNLOAD
	execution := models.ExecutionFilter{
		ExecuterName:  filter.PayloadData["executer_name"].(string),
		CarPlat:       filter.PayloadData["car_plat"].(string),
		CustomerName:  filter.PayloadData["customer_name"].(string),
		CustomerPhone: filter.PayloadData["customer_phone"].(string),
		CustomerEmail: filter.PayloadData["customer_email"].(string),
		Price:         helper.InterfaceSliceToSliceInt64(filter.PayloadData["price"]),
		Paid:          helper.InterfaceSliceToSliceInt64(filter.PayloadData["paid"]),
		PaidOff:       helper.InterfacePointerBoolToPointerBool(filter.PayloadData["paid_off"]),
		PaidFrom:      helper.InterfacePointerTimeToPointerTime(filter.PayloadData["paid_from"]),
		PaidTo:        helper.InterfacePointerTimeToPointerTime(filter.PayloadData["paid_to"]),
		ExecutedFrom:  helper.InterfacePointerTimeToPointerTime(filter.PayloadData["executed_from"]),
		ExecutedTo:    helper.InterfacePointerTimeToPointerTime(filter.PayloadData["executed_to"]),
		FinishFrom:    helper.InterfacePointerTimeToPointerTime(filter.PayloadData["finish_from"]),
		FinishTo:      helper.InterfacePointerTimeToPointerTime(filter.PayloadData["finish_to"]),
	}

	err := helper.WrapPercentOnStructString(&execution)
	if err != nil {
		terr = terror.New(err)
		return
	}

	executionsRes, _, terr := u.bengkelRepo.ExecutionSearch(ctx, execution, filter)
	if terr != nil {
		return
	}
	fmt.Printf("executionsRes: %+v\n", executionsRes)

	// generate excel
	{
		randomString := helper.CreateRandomString(10)
		filePath := fmt.Sprintf("storage/docs/executions_%s.xlsx", randomString)
		err = helper.GenerateExcel(ctx, executionsRes, filePath, "Sheet1")
		if err != nil {
			terr = terror.New(err)
			return
		}

		res, err = helper.GetDocBase64(ctx, filePath)
		if err != nil {
			terr = terror.New(err)
			return
		}
	}

	return
}
