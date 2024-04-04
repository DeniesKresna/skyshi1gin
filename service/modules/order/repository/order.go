package repository

import (
	"errors"

	"github.com/DeniesKresna/skyshi1gin/service/extensions/terror"
	"github.com/DeniesKresna/skyshi1gin/types/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (r *OrderRepository) GetDB(ctx *gin.Context) (tx interface{}) {
	return r.db
}

func (r *OrderRepository) OrderList(ctx *gin.Context) (order []models.Order, terr terror.ErrInterface) {
	err := r.db.Find(&order).Error
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			terr = terror.ErrNotFoundData(err.Error())
			return
		}
		terr = terror.New(err)
	}
	return
}

func (r *OrderRepository) OrderGetByID(ctx *gin.Context, id int64) (order models.Order, terr terror.ErrInterface) {
	err := r.db.Take(&order, "id = ?", id).Error
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			terr = terror.ErrNotFoundData(err.Error())
			return
		}
		terr = terror.New(err)
	}
	return
}

func (r *OrderRepository) OrderCreate(ctx *gin.Context, req models.Order) (order models.Order, terr terror.ErrInterface) {
	err := r.db.Create(&req).Error
	if err != nil {
		terr = terror.New(err)
	}
	order = req
	return
}
