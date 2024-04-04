package repository

import (
	"github.com/DeniesKresna/skyshi1gin/service/extensions/terror"
	"github.com/DeniesKresna/skyshi1gin/types/models"
	"github.com/gin-gonic/gin"
)

func (r *WarehouseRepository) ProductGetByID(ctx *gin.Context, id int64) (product models.Product, terr terror.ErrInterface) {
	return r.productCross.ProductGetByID(ctx, id)
}
