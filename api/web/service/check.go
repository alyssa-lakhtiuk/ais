package service

import (
	"ais/entities"
	"ais/repository"
)

type checkService struct {
	repoAdditional        repository.EmployeeRepo
	repoAdditional2       repository.CustomerCardRepo
	repoAdditionalProduct repository.StoreProductRepo
	repoAdditionalSale    repository.SaleRepo
	repo                  repository.CheckRepo
}

func (s *checkService) Create(check []entities.CheckInput) (int, error) {
	if check[0].IDEmployee != "" {
		_, err := s.repoAdditional.GetEmployeeById(check[0].IDEmployee)
		if err != nil {
			// throw err ""
		}
	}
	_, err := s.repoAdditional2.GetCustomerCardByNumber(check[0].CustomerNumber)
	if err != nil {
		// throw err ""
	}
	_, err = s.repo.CreateCheck(GenerateRandomStr(10), check)
	if err != nil {
		return 0, err
	}
	return 0, nil
}

func (s *checkService) Delete(num string) error {
	return s.repo.DeleteCheck(num)
}

func (s *checkService) GetByNumber(num string) (entities.Check, error) {
	return s.repo.GetCheckByNumber(num)
}

func (s *checkService) GetAll() ([]entities.Check, error) {
	return s.repo.GetAllChecks()
}

func NewCheckService(repo repository.CheckRepo, repoEmployee repository.EmployeeRepo, repoCc repository.CustomerCardRepo,
	repoStPr repository.StoreProductRepo) *checkService {
	return &checkService{repo: repo, repoAdditional: repoEmployee, repoAdditional2: repoCc, repoAdditionalProduct: repoStPr}
}
