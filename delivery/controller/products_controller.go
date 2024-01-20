package controller

import (
  "database/sql"
  "errors"
  "field_work/config"
  "field_work/delivery/middleware"
  "field_work/entity"
  "field_work/helpers"
  "field_work/shared/common"
  "field_work/usecase"
  "fmt"
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

func (c *productsController) getHandler(ctx *gin.Context) {
  id := ctx.Param("id")
  product, err := c.productsUseCase.GetProductByID(id)
  if err != nil {
    if errors.Is(err, sql.ErrNoRows) {
      common.SendErrorResponse(ctx, http.StatusNotFound, fmt.Sprintf("Product with id '%s' not found", id))
    } else {
      common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
    }
    return
  }
  
  product.TimeFormat("CreatedAt", "UpdatedAt")
  
  common.SendSingleResponse(ctx, product, "Get product successfully")
}

func (c *productsController) Route() {
  admin := c.rg.Group(config.AdminGroup)
  admin.POST(config.Products, c.authMiddleware.RequireToken("Admin"), c.insertHandler)
  admin.GET(config.Products, c.authMiddleware.RequireToken("Admin"), c.listHandler)
  admin.GET(config.ProductByID, c.authMiddleware.RequireToken("Admin"), c.getHandler)
  
  customer := c.rg.Group(config.CustomerGroup)
  customer.GET(config.Products, c.authMiddleware.RequireToken("Customer"), c.listHandler)
  customer.GET(config.ProductByID, c.authMiddleware.RequireToken("Customer"), c.getHandler)
}

func NewProductsController(productsUseCase usecase.ProductsUseCase, rg *gin.RouterGroup, authMiddleware middleware.AuthMiddleware) *productsController {
  return &productsController{productsUseCase, rg, authMiddleware}
}