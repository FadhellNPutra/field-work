package controller

import (
	"field_work/config"
	"field_work/entity"
	"field_work/entity/dto"
	"field_work/shared/common"
	"field_work/usecase"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
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
	common.SendCreatedResponse(ctx, rsv, "", "Login successfully")
}

func (a *AuthController) registerHandler(ctx *gin.Context) {
	role := ctx.Request.RequestURI

	var payload entity.Users
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	user, err := a.authUC.Register(payload, role)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	var userDTO dto.UserDTO
	if err := copier.Copy(&userDTO, &user); err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendCreatedResponse(ctx, userDTO, time.Now().Format("Monday 02, January 2006 15:04:05"), "Ok")
}

func (a *AuthController) Route() {
	a.rg.POST(config.AdminRegister, a.registerHandler)
	a.rg.POST(config.AdminLogin, a.loginHandler)
	a.rg.POST(config.CustomerRegister, a.registerHandler)
	a.rg.POST(config.CustomerLogin, a.loginHandler)
}

func NewAuthController(authUC usecase.AuthUseCase, rg *gin.RouterGroup) *AuthController {
	return &AuthController{authUC: authUC, rg: rg}
}
