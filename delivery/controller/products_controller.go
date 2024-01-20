package controller

import (
  "field_work/config"
  "field_work/delivery/middleware"
  "field_work/entity"
  "field_work/helpers"
  "field_work/shared/common"
  "field_work/usecase"
  "net/http"
  "strconv"
  "time"
  
  "github.com/gin-gonic/gin"
)

type productsController struct {
	productsUseCase usecase.ProductsUseCase
	rg             *gin.RouterGroup
	authMiddleware middleware.AuthMiddleware
}

func (c *productsController) insertHandler(ctx *gin.Context) {
  var payload entity.Products
  if err := ctx.ShouldBindJSON(&payload); err != nil {
    common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
    return
  }
  
  productDTO, err := c.productsUseCase.CreateNewProduct(payload)
  if err != nil {
    common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
    return
  }
  
  common.SendCreatedResponse(ctx, productDTO, time.Now().In(helpers.Location()).Format(time.RFC850), "Successfully create product")
}

func (c *productsController) listHandler(ctx *gin.Context) {
  page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
  size, _ := strconv.Atoi(ctx.DefaultQuery("size", "5"))
  
  products, paging, err := c.productsUseCase.ListProducts(page, size)
  if err != nil {
    common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
    return
  }
  
  var response []any
  for _, value := range products {
    value.TimeFormat("CreatedAt")
    response = append(response, value)
  }
  
  common.SendPagedResponse(ctx, response, paging, "List Products")
}

func (c *productsController) Route() {
  admin := c.rg.Group(config.AdminGroup)
  admin.POST(config.Products, c.authMiddleware.RequireToken("Admin"), c.insertHandler)
  admin.GET(config.Products, c.authMiddleware.RequireToken("Admin"), c.listHandler)
  
  customer := c.rg.Group(config.CustomerGroup)
  customer.GET(config.Products, c.authMiddleware.RequireToken("Customer"), c.listHandler)
}

func NewProductsController(productsUseCase usecase.ProductsUseCase, rg *gin.RouterGroup, authMiddleware middleware.AuthMiddleware) *productsController {
  return &productsController{productsUseCase, rg, authMiddleware}
}