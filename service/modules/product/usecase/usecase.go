package usecase

import (
	"github.com/DeniesKresna/skyshi1gin/service/extensions/terror"
	"github.com/DeniesKresna/skyshi1gin/service/modules/product/repository"
	"github.com/DeniesKresna/skyshi1gin/types/models"
	"github.com/gin-gonic/gin"
)

type ProductUsecase struct {
	productRepo repository.IRepository
}

func ProductCreateUsecase(repo repository.IRepository) IUsecase {
	productRepo := ProductUsecase{
		productRepo: repo,
	}
	return productRepo
}

type IUsecase interface {
	ProductList(ctx *gin.Context) (product []models.Product, terr terror.ErrInterface)
	ProductGetByID(ctx *gin.Context, id int64) (product models.Product, terr terror.ErrInterface)
	ProductCreate(ctx *gin.Context, req models.Product) (product models.Product, terr terror.ErrInterface)
}
