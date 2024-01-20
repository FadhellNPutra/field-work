package repository

import (
  "database/sql"
  "log"
  
  
	"field_work/config"
	"field_work/entity"
)

type ProductsRepository interface {
  Insert(payload entity.Products) (entity.Products, error)
}

type productsRepository struct {
	db *sql.DB
}

func (r *productsRepository) Insert(payload entity.Products) (entity.Products, error) {
  var product entity.Products
  
  if err := r.db.QueryRow(config.InsertProduct, payload.ProductName, payload.Quantity, payload.Price, payload.Material, payload.Description).Scan(
    &product.ID,
    &product.ProductName,
    &product.Quantity,
    &product.Price,
    &product.Material,
    &product.Description,
  ); err != nil {
    log.Println("InsertProduct.QueryRow Err :", err)
    return entity.Products{}, err
  }
  
  return product, nil
}

func NewProductsRepository(db *sql.DB) ProductsRepository {
  return &productsRepository{db}
}