package cross

import (
	"github.com/DeniesKresna/skyshi1gin/config"
	"github.com/DeniesKresna/skyshi1gin/service/extensions/terror"
	"github.com/DeniesKresna/skyshi1gin/service/modules/user/repository"
	"github.com/DeniesKresna/skyshi1gin/service/modules/user/usecase"
	"github.com/DeniesKresna/skyshi1gin/types/models"
	"github.com/gin-gonic/gin"
)

type UserCross struct {
	userUsecase usecase.IUsecase
}

func UserCreateCross(cfg *config.Config) UserCross {
	repo := repository.UserCreateRepository(cfg.DB)
	userUsecase := usecase.UserCreateUsecase(repo)
	return UserCross{
		userUsecase: userUsecase,
	}
}

func (h UserCross) AuthGetFromContext(ctx *gin.Context) (res models.UserRole, terr terror.ErrInterface) {
	return h.userUsecase.AuthGetFromContext(ctx)
}
