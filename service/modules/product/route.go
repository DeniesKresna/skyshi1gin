package user

import (
	"github.com/DeniesKresna/skyshi1gin/config"
	"github.com/DeniesKresna/skyshi1gin/service/middlewares"
	"github.com/DeniesKresna/skyshi1gin/service/modules/product/handler"
	"github.com/DeniesKresna/skyshi1gin/service/modules/product/repository"
	"github.com/DeniesKresna/skyshi1gin/service/modules/product/usecase"
	userUsecase "github.com/DeniesKresna/skyshi1gin/service/modules/user/usecase"
	"github.com/DeniesKresna/skyshi1gin/types/constants"
	"github.com/gin-gonic/gin"
)

func InitRoutes(v1 *gin.RouterGroup, userCase userUsecase.IUsecase, cfg *config.Config) {
	repo := repository.ProductCreateRepository(cfg.DB)
	ucase := usecase.ProductCreateUsecase(repo)
	handler := handler.ProductCreateHandler(ucase)

	moduleRoute := v1.Group("/product")

	adminRoute := moduleRoute.Use(roleCheck(userCase, constants.ROLES_ADMIN))
	{
		adminRoute.POST("/create", handler.ProductCreate)
	}

	authRoute := moduleRoute.Use(roleCheck(userCase, constants.ROLES_ADMIN, constants.ROLES_USER))
	{
		authRoute.GET("/list", handler.ProductList)
		authRoute.GET("/detail/:id", handler.ProductGetByID)
	}
}

func roleCheck(userCase userUsecase.IUsecase, roles ...constants.Roles) gin.HandlerFunc {
	return middlewares.CheckRole(userCase, roles)
}
