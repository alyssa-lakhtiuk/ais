package service

import (
	"ais/entities"
	"ais/repository"
)

type storeProductService struct {
	repoAdditional repository.ProductRepo
	repo           repository.StoreProductRepo
}

func (s *storeProductService) Create(product entities.StoreProduct) (int, error) {
	//_, err := s.repoAdditional.GetProductByNumber(product.IDProduct)
	//if err != nil {
	//	// throw err "Product of this store product doesn't exist"
	//}
	err := IsUnsigned(int(product.SellingPrice))
	if err != nil {
		return 0, err
	}
	return s.repo.CreateStoreProduct(product)
}

func (s *storeProductService) Update(upc string, product entities.StoreProduct) error {
	return s.repo.UpdateStoreProduct(upc, product)
}

func (s *storeProductService) Delete(upc string) error {
	return s.repo.DeleteStoreProduct(upc)
}

func (s *storeProductService) GetByName(upc string) (entities.StoreProduct, error) {
	return s.repo.GetStoreProductByUpc(upc)
}

func (s *storeProductService) GetAll() ([]entities.StoreProduct, error) {
	return s.repo.GetAllStoreProducts()
}

func NewStoreProductService(repo repository.ProductRepo, repoStore repository.StoreProductRepo) *storeProductService {
	return &storeProductService{repo: repoStore, repoAdditional: repo}
}

func (s *storeProductService) SearchUPC(upc string) ([]entities.StoreProduct, error) {
	return s.repo.SearchByUPC(upc)
}
