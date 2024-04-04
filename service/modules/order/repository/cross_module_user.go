package repository

import (
	"github.com/DeniesKresna/skyshi1gin/service/extensions/terror"
	"github.com/DeniesKresna/skyshi1gin/types/models"
	"github.com/gin-gonic/gin"
)

func (r *OrderRepository) AuthGetFromContext(ctx *gin.Context) (userRole models.UserRole, terr terror.ErrInterface) {
	return r.userCross.AuthGetFromContext(ctx)
}
