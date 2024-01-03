package controller

import (
	"field_work/entity"
	"field_work/shared/common"
	"field_work/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	usersUC usecase.UsersUseCase
	rg      *gin.RouterGroup
	// authMiddleware middleware.AuthMiddleware
}

func(u *UserController) createHandler(ctx *gin.Context){
	var payload entity.Users
	if err := ctx.ShouldBindJSON(&payload); err != nil{
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	user, err := u.usersUC.RegisterNewUsers(payload)
	if err != nil{
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	common.SendCreateResponse(ctx, user, "Created")
}

func (u *UserController) getByIdHandler(ctx *gin.Context){
	id := ctx.Param("id")
	user, err := u.usersUC.FindUsersByID(id)
	if err != nil{
		common.SendErrorResponse(ctx, http.StatusNotFound, "User with Id "+id+" not found")
		return
	}
	common.SendSingleResponse(ctx, user, "Ok")
}

func (u *UserController) getByUsernameHandler(ctx *gin.Context){
	username := ctx.Param("user")
	user, err := u.usersUC.FindUsersByUsername(username)
	if err != nil{
		common.SendErrorResponse(ctx, http.StatusNotFound, "User with username "+username+" not found")
		return
	}
	common.SendSingleResponse(ctx, user, "Ok")
}



func NewUsersController(usersUC usecase.UsersUseCase, rg *gin.RouterGroup /*,authMiddleware middleware.AuthMiddleware*/) *UserController {
	return &UserController{
		usersUC: usersUC,
		rg:      rg,
		// authMiddleware: authMiddleware
	}
}
