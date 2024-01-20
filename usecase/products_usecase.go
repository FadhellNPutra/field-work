package usecase

import (
  "log"
  
	"field_work/entity"
	"field_work/entity/dto"
	"field_work/repository"
	"field_work/shared/model"
	
	"github.com/jinzhu/copier"
)

type ProductsUseCase interface {
  CreateNewProduct(payload entity.Products) (dto.ProductsDTO, error)
  ListProducts(page, size int) ([]entity.Products, model.Paging, error)
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

func NewProductsUseCase(productsRepository repository.ProductsRepository) ProductsUseCase {
  return &productsUseCase{productsRepository}
}