package usecase

import (
  "field_work/entity"
  "field_work/repository"
)

type ProductImagesUseCase interface {
  CreateProductImage(payload entity.ProductImages) (entity.ProductImages, error)
}

type productImagesUseCase struct {
  productImagesRepository repository.ProductImagesRepository
}

func (u *productImagesUseCase) CreateProductImage(payload entity.ProductImages) (entity.ProductImages, error) {
  return u.productImagesRepository.Insert(payload)
}

func NewProductImagesUseCase(productImagesRepository repository.ProductImagesRepository) ProductImagesUseCase {
  return &productImagesUseCase{productImagesRepository}
}