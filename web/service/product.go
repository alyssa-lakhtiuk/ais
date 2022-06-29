package service

import (
	"ais/entities"
	"ais/repository"
)

type ProductService struct {
	repo repository.ProductRepo
}

func (s *ProductService) Create(product entities.Product) (int, error) {
	return s.repo.CreateProduct(product)
}

func (s *ProductService) Update(productId int, product entities.ProductInput) error {
	return s.repo.UpdateProduct(productId, product)
}

func (s *ProductService) Delete(productId int) error {
	return s.repo.DeleteProduct(productId)
}

func (s *ProductService) GetByName(productName string) (entities.Product, error) {
	return s.repo.GetProductByName(productName)
}

func (s *ProductService) GetAll() ([]entities.Product, error) {
	return s.repo.GetAllProducts()
}

func NewProductService(repo repository.ProductRepo) *ProductService {
	return &ProductService{repo: repo}
}
