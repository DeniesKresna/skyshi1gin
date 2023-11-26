package repository

import (
	"github.com/DeniesKresna/bengkelgin/service/extensions/terror"
	"github.com/DeniesKresna/bengkelgin/types/constants"
	"github.com/DeniesKresna/bengkelgin/types/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (r BengkelRepository) CustomerSearch(ctx *gin.Context, cust models.Customer, searchPayload models.DbSearchObject) (custs []models.Customer, totalData int64, terr terror.ErrInterface) {
	queryDB := r.db

	// filter
	{
		if cust.Name != "" {
			queryDB = queryDB.Where("name like ?", cust.Name)
		}
		if cust.Email != "" {
			queryDB = queryDB.Where("email like ?", cust.Email)
		}
		if cust.Phone != "" {
			queryDB = queryDB.Where("phone like ?", cust.Phone)
		}
	}

	queryDB = queryDB.Session(&gorm.Session{})

	if searchPayload.Mode == constants.DB_MODE_DATA || searchPayload.Mode == constants.DB_MODE_PAGE {
		offset := (searchPayload.Page - 1) * searchPayload.Limit
		for _, v := range searchPayload.Order {
			queryDB = queryDB.Order(v)
		}
		err := queryDB.Limit(int(searchPayload.Limit)).Offset(int(offset)).Find(&custs).Error
		if err != nil {
			terr = terror.New(err)
			return
		}
	}

	if searchPayload.Mode == constants.DB_MODE_COUNT || searchPayload.Mode == constants.DB_MODE_PAGE {
		err := queryDB.Model(&models.Customer{}).Limit(-1).Offset(-1).Count(&totalData).Error
		if err != nil {
			terr = terror.New(err)
			return
		}
	}

	return
}

func (r BengkelRepository) CustomerCreate(ctx *gin.Context, cust *models.Customer) (terr terror.ErrInterface) {
	err := r.db.Create(cust).Error
	if err != nil {
		terr = terror.New(err)
	}
	return
}

func (r BengkelRepository) CustomerUpdate(ctx *gin.Context, cust *models.Customer) (terr terror.ErrInterface) {
	err := r.db.Save(cust).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			terr = terror.ErrNotFoundData(err.Error())
			return
		}
		terr = terror.New(err)
	}
	return
}

func (r BengkelRepository) CustomerGetByPhone(ctx *gin.Context, phone string) (cust models.Customer, terr terror.ErrInterface) {
	err := r.db.First(&cust, "phone = ?", phone).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			terr = terror.ErrNotFoundData(err.Error())
			return
		}
		terr = terror.New(err)
	}
	return
}

func (r BengkelRepository) CustomerGetByID(ctx *gin.Context, id int64) (customer models.Customer, terr terror.ErrInterface) {
	err := r.db.First(&customer, "id = ?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			terr = terror.ErrNotFoundData(err.Error())
			return
		}
		terr = terror.New(err)
	}
	return
}

func (r BengkelRepository) CustomerListGetByName(ctx *gin.Context, name string) (customer []models.Customer, terr terror.ErrInterface) {
	err := r.db.Find(&customer, "name like ?", name).Error
	if err != nil {
		terr = terror.New(err)
	}
	return
}
