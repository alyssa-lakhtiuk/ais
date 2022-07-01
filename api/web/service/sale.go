package service

import (
	"ais/entities"
	"ais/repository"
)

type saleService struct {
	repo          repository.SaleRepo
	repoAddition  repository.StoreProductRepo
	repoAddition2 repository.CheckRepo
}

//func (s *saleService) Update(categoryNum int, category entities.Sale) error {
//	_, err := s.repo.GetCategoryByName(category.Name)
//	if err != nil {
//		return err
//	}
//	return s.repo. (categoryNums, category)
//}

func (s *saleService) Delete(upc, checkNum string) error {
	_, err := s.repo.GetSaleByUpcCheck(upc, checkNum)
	if err != nil {
		return err
	}
	return s.repo.DeleteSale(upc, checkNum)
}

func (s *saleService) GetByUpcCheck(upc, checkNum string) (entities.Sale, error) {
	return s.repo.GetSaleByUpcCheck(upc, checkNum)
}

func (s *saleService) GetAll() ([]entities.Sale, error) {
	return s.repo.GetAllSales()
}

func NewSaleService(repo repository.SaleRepo, repoStProduct repository.StoreProductRepo, repoCheck repository.CheckRepo) *saleService {
	return &saleService{repo: repo, repoAddition: repoStProduct, repoAddition2: repoCheck}
}

func (s *saleService) Create(sale entities.Sale) (int, error) {
	_, err := s.repoAddition.GetStoreProductByUpc(sale.UPC)
	if err != nil {
		// throw there is no such store product
	}
	_, err = s.repoAddition2.GetCheckByNumber(sale.CheckNumber)
	if err != nil {
		// throw there is no such check
	}
	return s.repo.CreateSale(sale)
}
