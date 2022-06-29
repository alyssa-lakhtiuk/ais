package service

import (
	"ais/entities"
	"ais/repository"
)

type productService struct {
	repo repository.ProductRepo
}

func (s *productService) Create(product entities.Product) (int, error) {
	return s.repo.CreateProduct(product)
}

func (s *productService) Update(productId int, product entities.ProductInput) error {
	return s.repo.UpdateProduct(productId, product)
}

func (s *productService) Delete(productId int) error {
	return s.repo.DeleteProduct(productId)
}

func (s *productService) GetByName(productName string) (entities.Product, error) {
	return s.repo.GetProductByName(productName)
}

func (s *productService) GetAll() ([]entities.Product, error) {
	return s.repo.GetAllProducts()
}

func NewProductService(repo repository.ProductRepo) *productService {
	return &productService{repo: repo}
}
