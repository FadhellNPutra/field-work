package repository

import (
  "database/sql"
  "field_work/config"
  "field_work/entity"
  "log"
)

type ProductImagesRepository interface {
  Insert(payload entity.ProductImages) (entity.ProductImages, error)
}

type productImagesRepository struct {
  db *sql.DB
}

func (r *productImagesRepository) Insert(payload entity.ProductImages) (entity.ProductImages, error) {
  var productImages entity.ProductImages
  
  if err := r.db.QueryRow(config.InsertProductImage, payload.ProductID, payload.FileName, payload.IsPrimary).Scan(
    &productImages.ID,
    &productImages.ProductID,
    &productImages.FileName,
    &productImages.IsPrimary,
  ); err != nil {
    log.Println("productImagesRepository: InsertProductImages.QueryRow.Scan Err :", err)
    return entity.ProductImages{}, err
  }
  
  return productImages, nil
}

func NewProductImagesRepository(db *sql.DB) ProductImagesRepository {
  return &productImagesRepository{db}
}