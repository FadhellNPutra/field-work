package controller

import (
  "fmt"
  "field_work/config"
  "field_work/entity"
  "field_work/usecase"
  "field_work/delivery/middleware"
  "field_work/helpers"
  "field_work/shared/common"
  "log"
  "net/http"
  "path/filepath"
  "strconv"
  "time"
  
  "github.com/gin-gonic/gin"
)

type productImagesController struct {
  productImagesUseCase usecase.ProductImagesUseCase
  rg *gin.RouterGroup
  authMiddleware middleware.AuthMiddleware
}

func (c *productImagesController) insertHandler(ctx *gin.Context) {
  productID := ctx.Query("product_id")
  isPrimary := ctx.Query("is_primary")
  payload := entity.ProductImages{}
  
  file, err := ctx.FormFile("image")
  if err != nil {
    log.Println("productImagesController: insertHandler.File Err :", err)
    common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
    return
  }
  
  allowedExtensions := []string{".jpg", ".jpeg", ".png", ".webp", ".aviv"}
  fileName := fmt.Sprintf("assets/images/%s-%s", productID, file.Filename)
  ext := filepath.Ext(fileName)
  
  if !payload.IsAllowedExtension(ext, allowedExtensions...) {
    log.Println("productImagesController: insertHandler.IsAllowedExtension Err : Extension Not Allowed")
    common.SendErrorResponse(ctx, http.StatusBadRequest, "Extension Not Allowed")
    return
  }
  
  if file.Size > 5<<20 {
    common.SendErrorResponse(ctx, http.StatusBadRequest, "Image size is too big!")
    return
  }
  
  if err := ctx.SaveUploadedFile(file, fileName); err != nil {
    log.Println("productImagesController: insertHandler.SaveUploadedFile Err : Extension Not Allowed")
    common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
    return
  }
  
  payload.ProductID = productID
  payload.FileName = fileName
  payload.IsPrimary, err = strconv.ParseBool(isPrimary)
  if err != nil {
    log.Println("productImagesController: ParseBool Err :", err)
    common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
    return
  }
  
  productImage, err := c.productImagesUseCase.CreateProductImage(payload)
  if err != nil {
    log.Println("productImagesController: CreateProductImage Err :", err)
    common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
    return
  }
  
  common.SendCreatedResponse(ctx, productImage, time.Now().In(helpers.Location()).Format(time.RFC850), "Create product image successfully")
}

func (c *productImagesController) Route() {
  admin := c.rg.Group(config.AdminGroup)
  admin.POST(config.ProductImages, c.authMiddleware.RequireToken("Admin"), c.insertHandler)
}

func NewProductImagesController(productImagesUseCase usecase.ProductImagesUseCase, rg *gin.RouterGroup, authMiddleware middleware.AuthMiddleware) *productImagesController {
  return &productImagesController{productImagesUseCase, rg, authMiddleware}
}