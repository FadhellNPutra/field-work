package controller

import (
	"database/sql"
	"errors"
	"field_work/config"
	"field_work/delivery/middleware"
	"field_work/entity"
	"field_work/entity/dto"
	"field_work/shared/common"
	"field_work/usecase"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type UserController struct {
	usersUC        usecase.UsersUseCase
	rg             *gin.RouterGroup
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

	createdAt, _ := time.Parse("2006-01-02T15:04:05+07:00", user.CreatedAt)

	var userDto dto.UserDTO
	if err := copier.Copy(&userDto, &user); err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendCreatedResponse(ctx, userDto, createdAt.Format("Monday 02, January 2006 15:04:05"), "Create user successfully")
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

func (u *UserController) listHandler(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(ctx.DefaultQuery("size", "5"))
	username := ctx.Query("username")

	if username != "" {
		user, err := u.usersUC.FindUsersByUsername(username)
		if err != nil {
			common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		common.SendSingleResponse(ctx, user, "Get users successfully")
		return
	} else {
		user, paging, err := u.usersUC.ListAll(page, size)
		if err != nil {
			common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
			return
		}
		var response []interface{}

		for _, v := range user {
			response = append(response, v)
		}
		common.SendPagedResponse(ctx, response, paging, "Ok")
	}
}

func (u *UserController) putHandler(ctx *gin.Context) {
	var payload dto.UpdateUserDTO
	var id = ctx.Param("id")
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, "Failed to bind data : "+err.Error())
		return
	}
	user, err := u.usersUC.UpdateUsers(payload, id)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}

	updatedAt, _ := time.Parse("2006-01-02 15:04:05", user.UpdatedAt)

	var userDto dto.UserDTO
	if err := copier.Copy(&userDto, &user); err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendUpdatedResponse(ctx, userDto, updatedAt.Format("Monday 02, January 2006 15:04:05"), "Update user successfully")
}

func (u *UserController) deleteHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	err := u.usersUC.DeleteUsers(id)

	log.Println(err)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			common.SendErrorResponse(ctx, http.StatusNotFound, "Employee with ID "+id+" not found. Delete data failed")
		} else {
			common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		}

		return
	}

	common.SendDeletedResponse(ctx, "Delete user successfully")
}

func (u *UserController) Route() {
	admin := u.rg.Group(config.AdminGroup)
	admin.GET(config.Users, u.authMiddleware.RequireToken("Admin"), u.listHandler)
	admin.GET(config.UsersByID, u.authMiddleware.RequireToken("Admin"), u.getByIdHandler)
	admin.POST(config.Users, u.authMiddleware.RequireToken("Admin"), u.createHandler)
	admin.PUT(config.UsersByID, u.authMiddleware.RequireToken("Admin"), u.putHandler)
	admin.DELETE(config.UsersByID, u.authMiddleware.RequireToken("Admin"), u.deleteHandler)

	customers := u.rg.Group(config.CustomerGroup)
	customers.GET(config.Users, u.authMiddleware.RequireToken("Customer"), u.listHandler)
	customers.GET(config.UsersByID, u.authMiddleware.RequireToken("Customer"), u.getByIdHandler)
	customers.PUT(config.UsersByID, u.authMiddleware.RequireToken("Customer"), u.putHandler)
	customers.DELETE(config.UsersByID, u.authMiddleware.RequireToken("Customer"), u.deleteHandler)
}

func NewUsersController(usersUC usecase.UsersUseCase, rg *gin.RouterGroup, authMiddleware middleware.AuthMiddleware) *UserController {
	return &UserController{
		usersUC:        usersUC,
		rg:             rg,
		authMiddleware: authMiddleware,
	}
}
