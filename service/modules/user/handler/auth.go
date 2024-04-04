package handler

import (
	"github.com/DeniesKresna/skyshi1gin/service/extensions/terror"
	"github.com/DeniesKresna/skyshi1gin/types/models"
	"github.com/gin-gonic/gin"
)

func (h UserHandler) AuthLogin(ctx *gin.Context) {
	var (
		loginReq models.LoginRequest
		terr     terror.ErrInterface
	)

	if err := ctx.ShouldBindJSON(&loginReq); err != nil {
		terr = terror.ErrParameter(err.Error())
		ResponseJson(ctx, terr)
		return
	}

	authResp, terr := h.userUsecase.AuthLogin(ctx, loginReq.Identifier, loginReq.Password)
	if terr != nil {
		ResponseJson(ctx, terr)
		return
	}
	ResponseJson(ctx, authResp)
}
