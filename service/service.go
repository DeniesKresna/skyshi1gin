package service

import (
	"github.com/DeniesKresna/skyshi1gin/config"
	"github.com/DeniesKresna/skyshi1gin/service/middlewares"

	userModule "github.com/DeniesKresna/skyshi1gin/service/modules/user"
	userrepo "github.com/DeniesKresna/skyshi1gin/service/modules/user/repository"
	usercase "github.com/DeniesKresna/skyshi1gin/service/modules/user/usecase"

	orderModule "github.com/DeniesKresna/skyshi1gin/service/modules/order"
	productModule "github.com/DeniesKresna/skyshi1gin/service/modules/product"
	warehouseModule "github.com/DeniesKresna/skyshi1gin/service/modules/warehouse"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func setRoutes(cfg *config.Config) (r *gin.Engine, err error) {
	r = gin.New()

	userRepo := userrepo.UserCreateRepository(cfg.DB)
	userCase := usercase.UserCreateUsecase(userRepo)

	r.Use(corsConfig())
	r.Use(middlewares.ActivityLogger())

	r.StaticFile("/", "./storage/dist/index.html")
	r.StaticFile("/caricon.png", "./storage/dist/caricon.png")
	r.Static("/static", "./storage/dist/static")
	r.Static("/assets", "./storage/dist/assets")
	r.NoRoute(func(c *gin.Context) {
		c.File("./storage/dist/index.html")
	})

	api := r.Group("/api")
	v1 := api.Group("/v1")
	{
		userModule.InitRoutes(v1, userCase, cfg)
		productModule.InitRoutes(v1, userCase, cfg)
		warehouseModule.InitRoutes(v1, userCase, cfg)
		orderModule.InitRoutes(v1, userCase, cfg)
	}

	return
}

func corsConfig() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "DELETE", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // Include "Content-Type" in the list of allowed headers
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	})
}

func Start(cfg *config.Config) (err error) {
	eng, err := setRoutes(cfg)

	eng.Run(cfg.App.Host + ":" + cfg.App.Port)
	return
}
