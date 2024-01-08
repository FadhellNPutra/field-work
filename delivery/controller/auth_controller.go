package controller

import (
	"field_work/config"
	"field_work/entity/dto"
	"field_work/shared/common"
	"field_work/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authUC usecase.AuthUseCase
	rg     *gin.RouterGroup
}

func (a *AuthController) loginHandler(ctx *gin.Context) {
	var payload dto.AuthRequestDto
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	rsv, err := a.authUC.Login(payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	common.SendCreatedResponse(ctx, rsv, "", "Ok")
}

func (a *AuthController) Route() {
	admin := a.rg.Group(config.AdminGroup)
	admin.POST(config.Login, a.loginHandler)

	customers := a.rg.Group(config.CustomerGroup)
	customers.POST(config.Login, a.loginHandler)
}

func NewAuthController(authUC usecase.AuthUseCase, rg *gin.RouterGroup) *AuthController {
	return &AuthController{authUC: authUC, rg: rg}
}
