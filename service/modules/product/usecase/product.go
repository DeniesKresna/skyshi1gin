package usecase

import (
	"strings"

	"github.com/DeniesKresna/skyshi1gin/service/extensions/terror"
	"github.com/DeniesKresna/skyshi1gin/types/models"
	"github.com/gin-gonic/gin"
)

func (u ProductUsecase) ProductList(ctx *gin.Context) (product []models.Product, terr terror.ErrInterface) {
	product, terr = u.productRepo.ProductList(ctx)

	return
}

func (u ProductUsecase) ProductGetByID(ctx *gin.Context, id int64) (product models.Product, terr terror.ErrInterface) {
	product, terr = u.productRepo.ProductGetByID(ctx, id)

	return
}

func (u ProductUsecase) ProductCreate(ctx *gin.Context, req models.Product) (product models.Product, terr terror.ErrInterface) {
	req.Code = strings.ToUpper(req.Code)
	product, terr = u.productRepo.ProductCreate(ctx, req)

	return
}
