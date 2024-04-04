package cross

import (
	"github.com/DeniesKresna/skyshi1gin/config"
	"github.com/DeniesKresna/skyshi1gin/service/extensions/terror"
	"github.com/DeniesKresna/skyshi1gin/service/modules/product/repository"
	"github.com/DeniesKresna/skyshi1gin/service/modules/product/usecase"
	"github.com/DeniesKresna/skyshi1gin/types/models"
	"github.com/gin-gonic/gin"
)

type ProductCross struct {
	productUsecase usecase.IUsecase
}

func ProductCreateCross(cfg *config.Config) ProductCross {
	repo := repository.ProductCreateRepository(cfg.DB)
	productUsecase := usecase.ProductCreateUsecase(repo)
	return ProductCross{
		productUsecase: productUsecase,
	}
}

func (h ProductCross) ProductList(ctx *gin.Context) (res []models.Product, terr terror.ErrInterface) {
	return h.productUsecase.ProductList(ctx)
}

func (h ProductCross) ProductGetByID(ctx *gin.Context, id int64) (res models.Product, terr terror.ErrInterface) {
	return h.productUsecase.ProductGetByID(ctx, id)
}
