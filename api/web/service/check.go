package service

import (
	"ais/entities"
	"ais/repository"
)

type checkService struct {
	repoAdditional  repository.EmployeeRepo
	repoAdditional2 repository.CustomerCardRepo
	repo            repository.CheckRepo
}

func (s *checkService) Create(check entities.Check) (int, error) {
	_, err := s.repoAdditional.GetEmployeeById(check.IdEmployee)
	if err != nil {
		// throw err ""
	}
	_, err = s.repoAdditional2.GetCustomerCardByNumber(check.CardNumber)
	if err != nil {
		// throw err ""
	}
	err = IsUnsigned(check.SumTotal)
	if err != nil {
		return 0, err
	}
	return s.repo.CreateCheck(check)
}

//func (s *checkService) Update(upc string, product entities.Check) error {
//	return s.repo.(upc, product)
//}

func (s *checkService) Delete(num string) error {
	return s.repo.DeleteCheck(num)
}

func (s *checkService) GetByNumber(num string) (entities.Check, error) {
	return s.repo.GetCheckByNumber(num)
}

func (s *checkService) GetAll() ([]entities.Check, error) {
	return s.repo.GetAllChecks()
}

func NewCheckService(repo repository.CheckRepo, repoEmployee repository.EmployeeRepo, repoCc repository.CustomerCardRepo) *checkService {
	return &checkService{repo: repo, repoAdditional: repoEmployee, repoAdditional2: repoCc}
}
