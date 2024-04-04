package user

import (
	"github.com/DeniesKresna/skyshi1gin/config"
	"github.com/DeniesKresna/skyshi1gin/service/middlewares"
	"github.com/DeniesKresna/skyshi1gin/service/modules/user/handler"
	"github.com/DeniesKresna/skyshi1gin/service/modules/user/usecase"
	"github.com/DeniesKresna/skyshi1gin/types/constants"
	"github.com/gin-gonic/gin"
)

func InitRoutes(v1 *gin.RouterGroup, userCase usecase.IUsecase, cfg *config.Config) {
	userHandler := handler.UserCreateHandler(userCase)

	moduleRoute := v1.Group("/user")

	moduleRoute.POST("/login", userHandler.AuthLogin)

	adminRoute := moduleRoute.Use(roleCheck(userCase, constants.ROLES_ADMIN))
	{
		adminRoute.PUT("/edit", userHandler.UserUpdate)
		adminRoute.GET("/detail/:id", userHandler.UserGetByID)
		adminRoute.POST("/search", userHandler.UserSearch)
	}

	authRoute := moduleRoute.Use(roleCheck(userCase, constants.ROLES_ADMIN, constants.ROLES_USER))
	{
		authRoute.POST("/create", userHandler.UserCreate)
		authRoute.POST("/get-by-email", userHandler.UserGetByEmail)
		authRoute.POST("/list-user", userHandler.UserGetAllUser)
	}
}

func roleCheck(userCase usecase.IUsecase, roles ...constants.Roles) gin.HandlerFunc {
	return middlewares.CheckRole(userCase, roles)
}
