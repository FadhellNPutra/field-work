package usecase

import (
  "log"
  
	"field_work/entity"
	"field_work/entity/dto"
	"field_work/repository"
	
	"github.com/jinzhu/copier"
)

type ProductsUseCase interface {
  CreateNewProduct(payload entity.Products) (dto.ProductsDTO, error)
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

func NewProductsUseCase(productsRepository repository.ProductsRepository) ProductsUseCase {
  return &productsUseCase{productsRepository}
}