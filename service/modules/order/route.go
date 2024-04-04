package user

import (
	"github.com/DeniesKresna/skyshi1gin/config"
	"github.com/DeniesKresna/skyshi1gin/service/middlewares"
	"github.com/DeniesKresna/skyshi1gin/service/modules/order/handler"
	"github.com/DeniesKresna/skyshi1gin/service/modules/order/repository"
	"github.com/DeniesKresna/skyshi1gin/service/modules/order/usecase"
	productCrossHandler "github.com/DeniesKresna/skyshi1gin/service/modules/product/handler/cross"
	userUsecase "github.com/DeniesKresna/skyshi1gin/service/modules/user/usecase"
	warehouseCrossHandler "github.com/DeniesKresna/skyshi1gin/service/modules/warehouse/handler/cross"
	"github.com/DeniesKresna/skyshi1gin/types/constants"
	"github.com/gin-gonic/gin"
)

func InitRoutes(v1 *gin.RouterGroup, userCase userUsecase.IUsecase, cfg *config.Config) {
	// set up cross module for warehouse
	warehouseCross := warehouseCrossHandler.WarehouseCreateCross(cfg)
	productCross := productCrossHandler.ProductCreateCross(cfg)

	repo := repository.OrderCreateRepository(cfg.DB, warehouseCross, productCross)
	ucase := usecase.OrderCreateUsecase(repo)
	handler := handler.OrderCreateHandler(ucase)

	moduleRoute := v1.Group("/order")

	adminRoute := moduleRoute.Use(roleCheck(userCase, constants.ROLES_ADMIN))
	{

	}

	authRoute := moduleRoute.Use(roleCheck(userCase, constants.ROLES_ADMIN, constants.ROLES_USER))
	{
		authRoute.GET("/list", handler.OrderList)
		authRoute.GET("/detail/:id", handler.OrderGetByID)
		adminRoute.POST("/purchase", handler.OrderItem)
	}
}

func roleCheck(userCase userUsecase.IUsecase, roles ...constants.Roles) gin.HandlerFunc {
	return middlewares.CheckRole(userCase, roles)
}
