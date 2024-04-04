package user

import (
	"github.com/DeniesKresna/skyshi1gin/config"
	"github.com/DeniesKresna/skyshi1gin/service/middlewares"
	productCrossHandler "github.com/DeniesKresna/skyshi1gin/service/modules/product/handler/cross"
	userUsecase "github.com/DeniesKresna/skyshi1gin/service/modules/user/usecase"
	"github.com/DeniesKresna/skyshi1gin/service/modules/warehouse/handler"
	"github.com/DeniesKresna/skyshi1gin/service/modules/warehouse/repository"
	"github.com/DeniesKresna/skyshi1gin/service/modules/warehouse/usecase"
	"github.com/DeniesKresna/skyshi1gin/types/constants"
	"github.com/gin-gonic/gin"
)

func InitRoutes(v1 *gin.RouterGroup, userCase userUsecase.IUsecase, cfg *config.Config) {
	// set up cross module for product
	productCross := productCrossHandler.ProductCreateCross(cfg)

	repo := repository.WarehouseCreateRepository(cfg.DB, productCross)
	ucase := usecase.WarehouseCreateUsecase(repo)
	handler := handler.WarehouseCreateHandler(ucase)

	moduleRoute := v1.Group("/warehouse")

	adminRoute := moduleRoute.Use(roleCheck(userCase, constants.ROLES_ADMIN))
	{
		adminRoute.POST("/create", handler.WarehouseCreate)
		adminRoute.POST("/product/update-stock", handler.WarehouseProductUpdateStock)
		adminRoute.POST("/product/transfer-stock", handler.WarehouseProductTransfer)
	}

	authRoute := moduleRoute.Use(roleCheck(userCase, constants.ROLES_ADMIN, constants.ROLES_USER))
	{
		authRoute.GET("/list", handler.WarehouseList)
		authRoute.GET("/detail/:id", handler.WarehouseGetByID)
	}
}

func roleCheck(userCase userUsecase.IUsecase, roles ...constants.Roles) gin.HandlerFunc {
	return middlewares.CheckRole(userCase, roles)
}
