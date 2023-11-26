package handler

import (
	"fmt"
	"net/http"

	"github.com/DeniesKresna/bengkelgin/service/extensions/helper"
	"github.com/DeniesKresna/bengkelgin/service/extensions/terror"
	"github.com/DeniesKresna/bengkelgin/service/modules/user/usecase"
	"github.com/DeniesKresna/bengkelgin/types/constants"
	"github.com/DeniesKresna/bengkelgin/types/models"
	"github.com/DeniesKresna/gohelper/utstring"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserUsecase usecase.UserUsecase
}

func UserCreateHandler(userUsecase usecase.UserUsecase) UserHandler {
	return UserHandler{
		UserUsecase: userUsecase,
	}
}

func ResponseJson(ctx *gin.Context, data interface{}) {
	httpStatusCode := http.StatusOK
	resp, ok := data.(*terror.ErrorModel)
	if ok {
		responseData := models.Response{
			ResponseDesc: resp.Message,
			ResponseCode: resp.Code,
		}

		appEnv := utstring.GetEnv(constants.ENV_APP_ENV, "local")
		if appEnv != "production" {
			responseData.ResponseTrace = resp.Trace
		}

		ctx.JSON(httpStatusCode, responseData)
		return
	}

	responseData := models.Response{
		ResponseDesc: "Success",
		ResponseCode: "00",
	}

	if helper.IsStruct(data) || helper.IsMap(data) || helper.IsSlice(data) {
		responseData.ResponseData = data
	} else {
		responseData.ResponseDesc = fmt.Sprintf("%v", data)
	}

	ctx.JSON(httpStatusCode, responseData)
	return
}
