package service

import (
	"ais/entities"
	"ais/repository"
	"math/rand"
	"time"
)

type productService struct {
	repo         repository.ProductRepo
	repoCategory repository.CategoryRepo
}

func (s *productService) Create(product entities.Product) (int, error) {
	for true {
		x1 := rand.NewSource(time.Now().UnixNano())
		y1 := rand.New(x1)
		randId := y1.Intn(900 - IdRange)
		cat, _ := s.repo.GetProductByNumber(randId)
		if cat.Id != randId {
			product.Id = randId
			break
		}
	}
	return s.repo.CreateProduct(product)
}

func (s *productService) Update(productId int, product entities.Product) error {
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

func NewProductService(repo repository.ProductRepo, repoCategory repository.CategoryRepo) *productService {
	return &productService{repo: repo, repoCategory: repoCategory}
}
