package usecase

import (
	"github.com/DeniesKresna/skyshi1gin/service/extensions/terror"
	"github.com/DeniesKresna/skyshi1gin/service/modules/order/repository"
	"github.com/DeniesKresna/skyshi1gin/types/models"
	"github.com/gin-gonic/gin"
)

type OrderUsecase struct {
	orderRepo repository.IRepository
}

func OrderCreateUsecase(repo repository.IRepository) IUsecase {
	orderRepo := OrderUsecase{
		orderRepo: repo,
	}
	return orderRepo
}

type IUsecase interface {
	OrderList(ctx *gin.Context) (order []models.Order, terr terror.ErrInterface)
	OrderGetByID(ctx *gin.Context, id int64) (order models.Order, terr terror.ErrInterface)
	OrderItem(ctx *gin.Context, req []models.Item) (order models.Order, terr terror.ErrInterface)
}
