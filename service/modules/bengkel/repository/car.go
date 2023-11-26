package repository

import (
	"github.com/DeniesKresna/bengkelgin/service/extensions/terror"
	"github.com/DeniesKresna/bengkelgin/types/constants"
	"github.com/DeniesKresna/bengkelgin/types/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (r BengkelRepository) CarSearch(ctx *gin.Context, car models.Car, searchPayload models.DbSearchObject) (cars []models.Car, totalData int64, terr terror.ErrInterface) {
	queryDB := r.db

	// filter
	{
		if car.Plat != "" {
			queryDB = queryDB.Where("plat like ?", car.Plat)
		}
		if car.CarType != "" {
			queryDB = queryDB.Where("car_type like ?", car.CarType)
		}
	}

	queryDB = queryDB.Session(&gorm.Session{})

	if searchPayload.Mode == constants.DB_MODE_DATA || searchPayload.Mode == constants.DB_MODE_PAGE {
		for _, v := range searchPayload.Order {
			queryDB = queryDB.Order(v)
		}
		offset := (searchPayload.Page - 1) * searchPayload.Limit
		err := queryDB.Limit(int(searchPayload.Limit)).Offset(int(offset)).Find(&cars).Error
		if err != nil {
			terr = terror.New(err)
			return
		}
	}

	if searchPayload.Mode == constants.DB_MODE_COUNT || searchPayload.Mode == constants.DB_MODE_PAGE {
		err := queryDB.Model(&models.Car{}).Limit(-1).Offset(-1).Count(&totalData).Error
		if err != nil {
			terr = terror.New(err)
			return
		}
	}

	return
}

func (r BengkelRepository) CarCreate(ctx *gin.Context, car *models.Car) (terr terror.ErrInterface) {
	err := r.db.Create(car).Error
	if err != nil {
		terr = terror.New(err)
	}
	return
}

func (r BengkelRepository) CarUpdate(ctx *gin.Context, car *models.Car) (terr terror.ErrInterface) {
	err := r.db.Save(car).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			terr = terror.ErrNotFoundData(err.Error())
			return
		}
		terr = terror.New(err)
	}
	return
}

func (r BengkelRepository) CarGetByID(ctx *gin.Context, id int64) (car models.Car, terr terror.ErrInterface) {
	err := r.db.First(&car, "id = ?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			terr = terror.ErrNotFoundData(err.Error())
			return
		}
		terr = terror.New(err)
	}
	return
}

func (r BengkelRepository) CarGetByPlat(ctx *gin.Context, plat string) (car models.Car, terr terror.ErrInterface) {
	err := r.db.First(&car, "plat like ?", plat).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			terr = terror.ErrNotFoundData(err.Error())
			return
		}
		terr = terror.New(err)
	}
	return
}

func (r BengkelRepository) CarListGetByPlat(ctx *gin.Context, plat string) (cars []models.Car, terr terror.ErrInterface) {
	err := r.db.Find(&cars, "plat like ?", plat).Error
	if err != nil {
		terr = terror.New(err)
	}
	return
}
