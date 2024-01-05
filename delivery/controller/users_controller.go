package controller

import (
	"field_work/config"
	"field_work/delivery/middleware"
	"field_work/entity"
	"field_work/shared/common"
	"field_work/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	usersUC usecase.UsersUseCase
	rg      *gin.RouterGroup
	authMiddleware middleware.AuthMiddleware
}

func (u *UserController) createHandler(ctx *gin.Context) {
	var payload entity.Users
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	user, err := u.usersUC.RegisterNewUsers(payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	common.SendCreateResponse(ctx, user, "Created")
}

func (u *UserController) getByIdHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := u.usersUC.FindUsersByID(id)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusNotFound, "User with Id "+id+" not found")
		return
	}
	common.SendSingleResponse(ctx, user, "Ok")
}

func (u *UserController) getByUsernameHandler(ctx *gin.Context) {
	username := ctx.Query("user")
	user, err := u.usersUC.FindUsersByUsername(username)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusNotFound, "User with username "+username+" not found")
		return
	}
	common.SendSingleResponse(ctx, user, "Ok")
}

func (u *UserController) listHandler(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(ctx.DefaultQuery("size", "5"))
	user, paging, err := u.usersUC.ListAll(page, size)
	if err != nil{
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	var response []interface{}

	for _, v := range user{
		response = append(response, v)
	}
	common.SendPagedResponse(ctx, response, paging, "Ok")
}

func (u *UserController) putHandler(ctx *gin.Context) {
	var payload entity.Users
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, "Failed to bind data")
		return
	}
	user, err := u.usersUC.UpdateUsers(payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}
	common.SendSingleResponse(ctx, user, "Updated Successfully")
}

func (u *UserController) deleteHandler(ctx *gin.Context){
	id := ctx.Param("id")
	err := u.usersUC.DeleteUsers(id)
	if err != nil{
		common.SendErrorResponse(ctx, http.StatusNotFound, "Employee with ID "+id+" not found. Delete data failed")
		return
	}
	common.SendSingleResponse(ctx, err, "Accepted")
}


func (u *UserController) Route(){
	u.rg.GET(config.UserList, u.authMiddleware.RequireToken("admin"), u.listHandler)
	u.rg.GET(config.UserById, u.authMiddleware.RequireToken("admin"), u.getByIdHandler)
	u.rg.GET(config.UserByUsername, u.authMiddleware.RequireToken("admin"), u.getByUsernameHandler)
	u.rg.POST(config.UserCreate, u.authMiddleware.RequireToken("admin", "customer"), u.createHandler)
	u.rg.PUT(config.UserUpdate, u.authMiddleware.RequireToken("admin", "customer"), u.putHandler)
	u.rg.DELETE(config.UserDelete, u.authMiddleware.RequireToken("admin", "customer"), u.deleteHandler)
}
func NewUsersController(usersUC usecase.UsersUseCase, rg *gin.RouterGroup ,authMiddleware middleware.AuthMiddleware) *UserController {
	return &UserController{
		usersUC: usersUC,
		rg:      rg,
		authMiddleware: authMiddleware,
	}
}
