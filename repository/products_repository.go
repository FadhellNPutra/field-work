package repository

import (
  "database/sql"
  "log"
  "math"
  
	"field_work/config"
	"field_work/entity"
	"field_work/shared/model"
)

type ProductsRepository interface {
  Insert(payload entity.Products) (entity.Products, error)
  FindAll(page, size int) ([]entity.Products, model.Paging, error)
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

func (r *productsRepository) FindAll(page, size int) ([]entity.Products, model.Paging, error) {
  var products []entity.Products
  offset := (page - 1) * size

  rows, err := r.db.Query(config.SelectAllProducts, size, offset)
  if err != nil {
    log.Println("productsRepository: FindAll.Query Err :", err)
    return nil, model.Paging{}, err
  }

  for rows.Next() {
    var product entity.Products
    if err := rows.Scan(
      &product.ID,
      &product.ProductName,
      &product.Quantity,
      &product.Price,
      &product.Material,
      &product.Description,
      &product.CreatedAt,
    ); err != nil {
      log.Println("productsRepository: FindAll.rows.Scan Err :", err)
      return nil, model.Paging{}, err
    }

    products = append(products, product)
  }

  totalRows := 0
  if err := r.db.QueryRow("SELECT COUNT(*) totalRows FROM products").Scan(&totalRows); err != nil {
    log.Println("productsRepository: FindAll.QueryRow.totalRows Err :", err)
    return nil, model.Paging{}, err
  }
  
  paging := model.Paging{
    Page: page,
    RowsPerPage: size,
    TotalRows: totalRows,
    TotalPages: int(math.Ceil(float64(totalRows) / float64(size))),
  }
  
  return products, paging, nil
}

func NewProductsRepository(db *sql.DB) ProductsRepository {
  return &productsRepository{db}
}