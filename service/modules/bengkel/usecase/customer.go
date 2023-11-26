package usecase

import (
	"github.com/DeniesKresna/bengkelgin/service/extensions/helper"
	"github.com/DeniesKresna/bengkelgin/service/extensions/terror"
	"github.com/DeniesKresna/bengkelgin/types/constants"
	"github.com/DeniesKresna/bengkelgin/types/models"
	"github.com/gin-gonic/gin"
)

func (u BengkelUsecase) CustomerGetByID(ctx *gin.Context, id int64) (customer models.Customer, terr terror.ErrInterface) {
	return u.bengkelRepo.CustomerGetByID(ctx, id)
}

func (u BengkelUsecase) CustomerGetByPhone(ctx *gin.Context, phone string) (customer models.Customer, terr terror.ErrInterface) {
	phone = helper.StandardizePhoneNumber(phone)
	return u.bengkelRepo.CustomerGetByPhone(ctx, phone)
}

func (u BengkelUsecase) CustomerCreate(ctx *gin.Context, customer models.Customer) (customerRes models.Customer, terr terror.ErrInterface) {
	customer.Phone = helper.StandardizePhoneNumber(customer.Phone)

	// check conflict
	{
		conflict := true
		_, terr = u.CarGetByPlat(ctx, customer.Phone)
		if terr != nil {
			if terr.GetType() == terror.ERROR_TYPE_DATA_NOT_FOUND {
				conflict = false
				terr = nil
			}
		}
		if conflict {
			terr = terror.ErrInvalidRule("Customer with the phone number has been exist")
			return
		}
	}

	terr = u.bengkelRepo.CustomerCreate(ctx, &customer)
	if terr != nil {
		return
	}

	customerRes = customer
	return
}

func (u BengkelUsecase) CustomerUpdate(ctx *gin.Context, customer models.Customer) (customerRes models.Customer, terr terror.ErrInterface) {
	customer.Phone = helper.StandardizePhoneNumber(customer.Phone)

	terr = u.bengkelRepo.CustomerUpdate(ctx, &customer)
	if terr != nil {
		return
	}

	customerRes = customer
	return
}

func (u BengkelUsecase) CustomerSearch(ctx *gin.Context, filter models.DbSearchObject) (res models.DbSearchObject, terr terror.ErrInterface) {
	filter.Mode = constants.DB_MODE_PAGE

	customer := models.Customer{}
	err := helper.MapAnyToStruct(filter.PayloadData, &customer)
	if err != nil {
		terr = terror.New(err)
		return
	}

	err = helper.WrapPercentOnStructString(&customer)
	if err != nil {
		terr = terror.New(err)
		return
	}

	customersRes, totalData, terr := u.bengkelRepo.CustomerSearch(ctx, customer, filter)
	if terr != nil {
		return
	}

	filter.ResponseData = customersRes
	filter.TotalData = totalData
	res = filter

	return
}

func (u BengkelUsecase) CustomerListGetByName(ctx *gin.Context, name string) (customers []models.Customer, terr terror.ErrInterface) {
	customerName := helper.WrapString(name, "%")
	return u.bengkelRepo.CustomerListGetByName(ctx, customerName)
}
