package service

import (
	"ais/entities"
	"ais/repository"
)

type zvitService struct {
	repo repository.Zvit
}

func NewZvitService(repo repository.Zvit) *zvitService {
	return &zvitService{repo: repo}
}

func (s *zvitService) GetPricesByCategory() ([]entities.PriceByCat, error) {
	return s.repo.GetPricesByCategories()
}

func (s *zvitService) GetChecksByCategory(category string) ([]entities.CheckByCat, error) {
	return s.repo.GetChecksByCat(category)
}

//func (s *zvitService) GetСountByCities() ( error) {
//	return s.repo.GetСountByCities()
//}
//
//func (s *zvitService) GetChecksByPrice() ( error) {
//	return s.repo.GetChecksByPrice()
//}
