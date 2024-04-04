package repository

import (
	"github.com/DeniesKresna/skyshi1gin/service/extensions/terror"
	"github.com/DeniesKresna/skyshi1gin/types/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func ProductCreateRepository(db *gorm.DB) IRepository {
	productRepository := ProductRepository{
		db: db,
	}
	return productRepository
}

type IRepository interface {
	ProductList(ctx *gin.Context) (product []models.Product, terr terror.ErrInterface)
	ProductGetByID(ctx *gin.Context, id int64) (product models.Product, terr terror.ErrInterface)
	ProductCreate(ctx *gin.Context, req models.Product) (product models.Product, terr terror.ErrInterface)
}
