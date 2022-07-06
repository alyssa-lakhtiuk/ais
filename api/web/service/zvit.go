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

func (s *zvitService) GetQuantitiesByCategory() ([]entities.QuantityByCat, error) {
	return s.repo.GetQuantitiesByCategories()
}

func (s *zvitService) GetChecksByCategory(category string) ([]entities.CheckByCat, error) {
	return s.repo.GetChecksByCat(category)
}

func (s *zvitService) GetCountByCities() ([]entities.CountCustomersCities, error) {
	return s.repo.CountCities()
}

func (s *zvitService) GetChecksByPrice(price int) ([]entities.SecondStruct, error) {
	return s.repo.ChecksByPrice(price)
}
