package service

import (
	"ais/entities"
	"ais/repository"
)

type customerCardService struct {
	repo repository.CustomerCardRepo
}

func (s *customerCardService) Update(ccNum string, cc entities.CustomerCard) error {
	_, err := s.repo.GetCustomerCardByNumber(ccNum)
	if err != nil {
		return err
	}
	return s.repo.UpdateCustomerCard(ccNum, cc)
}

func (s *customerCardService) Delete(name string) error {
	//_, err := s.repo.
	//if err != nil {
	//	return err
	//}
	return s.repo.DeleteCustomerCard(name)
}

func (s *customerCardService) GetByNumber(ccNumber string) (entities.CustomerCard, error) {
	return s.repo.GetCustomerCardByNumber(ccNumber)
}

func (s *customerCardService) GetAll() ([]entities.CustomerCard, error) {
	return s.repo.GetAllCustomerCards()
}

func NewCustomerCardServiceService(repo repository.CustomerCardRepo) *customerCardService {
	return &customerCardService{repo: repo}
}

func (s *customerCardService) Create(cc entities.CustomerCard) (int, error) {
	phone := cc.PhoneNumber
	err := ValidPhone(phone)
	if err != nil {
		return 0, err
	}
	err = IsUnsigned(cc.Percent)
	if err != nil {
		return 0, err
	}
	return s.repo.CreateCustomerCard(cc)
}
