package repository

import (
	"errors"

	"github.com/DeniesKresna/skyshi1gin/service/extensions/terror"
	"github.com/DeniesKresna/skyshi1gin/types/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (r ProductRepository) ProductList(ctx *gin.Context) (product []models.Product, terr terror.ErrInterface) {
	err := r.db.Find(&product).Error
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			terr = terror.ErrNotFoundData(err.Error())
			return
		}
		terr = terror.New(err)
	}
	return
}

func (r ProductRepository) ProductGetByID(ctx *gin.Context, id int64) (product models.Product, terr terror.ErrInterface) {
	err := r.db.Take(&product, "id = ?", id).Error
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			terr = terror.ErrNotFoundData(err.Error())
			return
		}
		terr = terror.New(err)
	}
	return
}

func (r ProductRepository) ProductCreate(ctx *gin.Context, req models.Product) (product models.Product, terr terror.ErrInterface) {
	err := r.db.Create(&req).Error
	if err != nil {
		terr = terror.New(err)
	}
	product = req
	return
}
