package usecase

import (
	"go-api/model"
	"go-api/repository"
)

type ProductUseCase struct {
	// Repository
	repository repository.ProductRepository
}

func NewProductUseCase(repository repository.ProductRepository) ProductUseCase {
	return ProductUseCase{
		repository: repository,
	}
}

func (pu *ProductUseCase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *ProductUseCase) CreateProduct(product model.Product) (model.Product, error) {
	productId, err := pu.repository.CreateProduct(product)

	if err != nil {
		return model.Product{}, err
	}

	product.ID = productId

	return product, nil
}

func (pu *ProductUseCase) GetProductById(id int) (*model.Product, error) {
	product, err := pu.repository.GetProductById(id)

	if err != nil {
		return nil, err
	}

	return product, nil
}
