package usecase

import (
  "fmt"
  "log"
  "reflect"
  "strconv"
  
	"field_work/entity"
	"field_work/entity/dto"
	"field_work/repository"
	"field_work/shared/model"
	
	"github.com/jinzhu/copier"
)

type ProductsUseCase interface {
  CreateNewProduct(payload entity.Products) (dto.ProductsDTO, error)
  ListProducts(page, size int) ([]entity.Products, model.Paging, error)
  GetProductByID(id string) (entity.Products, error)
  GetProductsByProductName(productName string, page, size int) ([]entity.Products, model.Paging, error)
  UpdateProductByID(payload entity.Products, id string) (entity.Products, error)
  DeleteProductByID(id string) error
}

type productsUseCase struct {
	productsRepository repository.ProductsRepository
}

func (u *productsUseCase) CreateNewProduct(payload entity.Products) (dto.ProductsDTO, error) {
  product, err := u.productsRepository.Insert(payload)
  if err != nil {
    return dto.ProductsDTO{}, err
  }
  
  var productDTO dto.ProductsDTO
  if err := copier.Copy(&productDTO, &product); err != nil {
    log.Println("CreateNewProduct.copier.Copy Err :", err)
    return dto.ProductsDTO{}, err
  }
  
  return productDTO, nil
}

func (u *productsUseCase) ListProducts(page, size int) ([]entity.Products, model.Paging, error) {
  return u.productsRepository.FindAll(page, size)
}

func (u *productsUseCase) GetProductByID(id string) (entity.Products, error) {
  return u.productsRepository.FindByID(id)
}

func (u *productsUseCase) GetProductsByProductName(productName string, page, size int) ([]entity.Products, model.Paging, error) {
  return u.productsRepository.FindByProductName(productName, page, size)
}

func (u *productsUseCase) UpdateProductByID(payload entity.Products, id string) (entity.Products, error) {
  product, err := u.productsRepository.FindByID(id)
  if err != nil {
    return entity.Products{}, err
  }
  
  payloadMap := payload.ToMap()
  productMap := product.ToMap()
  
  for key, val := range payloadMap {
    if *val == "" || *val == "0" {
      *val = *productMap[key]
    }
  }
  
  for key, val := range payloadMap {
		field := reflect.ValueOf(&payload).Elem().FieldByName(key)

		if field.Kind() == reflect.String {
			fieldValue := *val
			field.SetString(fieldValue)
		}
	}
	
	if err := copier.Copy(&product, &payload); err != nil {
	  log.Println("productsUseCase: UpdateProductByID.Copy Err :", err)
		return entity.Products{}, fmt.Errorf("failed to copy product struct: %v", err.Error())
	}
	
	product.Quantity, _ = strconv.ParseInt(*payloadMap["Quantity"], 10, 64)
	product.Price, _ = strconv.ParseInt(*payloadMap["Price"], 10, 64)

	newProduct, err := u.productsRepository.UpdateByID(product, id)
	if err != nil {
	  log.Println("productsUseCase: UpdateProductByID.UpdateByID Err :", err)
		return entity.Products{}, err
	}

	return newProduct, nil
}

func (u *productsUseCase) DeleteProductByID(id string) error {
  return u.productsRepository.DeleteByID(id)
}

func NewProductsUseCase(productsRepository repository.ProductsRepository) ProductsUseCase {
  return &productsUseCase{productsRepository}
}