package usecase

import (
	"strings"

	"github.com/DeniesKresna/bengkelgin/service/extensions/helper"
	"github.com/DeniesKresna/bengkelgin/service/extensions/terror"
	"github.com/DeniesKresna/bengkelgin/types/constants"
	"github.com/DeniesKresna/bengkelgin/types/models"
	"github.com/gin-gonic/gin"
)

func (u BengkelUsecase) CarGetByID(ctx *gin.Context, id int64) (car models.Car, terr terror.ErrInterface) {
	return u.bengkelRepo.CarGetByID(ctx, id)
}

func (u BengkelUsecase) CarGetByPlat(ctx *gin.Context, plat string) (car models.Car, terr terror.ErrInterface) {
	return u.bengkelRepo.CarGetByPlat(ctx, plat)
}

func (u BengkelUsecase) CarListGetByPlat(ctx *gin.Context, plat string) (cars []models.Car, terr terror.ErrInterface) {
	carPlat := helper.WrapString(plat, "%")
	return u.bengkelRepo.CarListGetByPlat(ctx, carPlat)
}

func (u BengkelUsecase) CarCreate(ctx *gin.Context, car models.Car) (carRes models.Car, terr terror.ErrInterface) {
	car.Plat = strings.ToUpper(car.Plat)

	// check conflict
	{
		conflict := true
		_, terr = u.CarGetByPlat(ctx, car.Plat)
		if terr != nil {
			if terr.GetType() == terror.ERROR_TYPE_DATA_NOT_FOUND {
				conflict = false
				terr = nil
			}
		}
		if conflict {
			terr = terror.ErrInvalidRule("Car with the plat number has been exist")
			return
		}
	}

	terr = u.bengkelRepo.CarCreate(ctx, &car)
	if terr != nil {
		return
	}

	carRes = car
	return
}

func (u BengkelUsecase) CarUpdate(ctx *gin.Context, car models.Car) (carRes models.Car, terr terror.ErrInterface) {
	car.Plat = strings.ToUpper(car.Plat)

	terr = u.bengkelRepo.CarUpdate(ctx, &car)
	if terr != nil {
		return
	}

	carRes = car
	return
}

func (u BengkelUsecase) CarSearch(ctx *gin.Context, filter models.DbSearchObject) (res models.DbSearchObject, terr terror.ErrInterface) {
	filter.Mode = constants.DB_MODE_PAGE

	car := models.Car{}
	err := helper.MapAnyToStruct(filter.PayloadData, &car)
	if err != nil {
		terr = terror.New(err)
		return
	}

	err = helper.WrapPercentOnStructString(&car)
	if err != nil {
		terr = terror.New(err)
		return
	}

	carsRes, totalData, terr := u.bengkelRepo.CarSearch(ctx, car, filter)
	if terr != nil {
		return
	}

	filter.ResponseData = carsRes
	filter.TotalData = totalData
	res = filter

	return
}
