package service

import (
	"ais/entities"
	"ais/repository"
)

type EmployeeService struct {
	repo repository.EmployeeRepo
}

func (s *EmployeeService) Update(employeeId string, employee entities.EmployeeInput) error {
	return s.repo.UpdateEmployee(employeeId, employee)
}

func (s *EmployeeService) Delete(employeeId string) error {
	return s.repo.DeleteEmployee(employeeId)
}

func (s *EmployeeService) GetByName(employeeName string) (entities.Employee, error) {
	return s.repo.GetEmployeeByName(employeeName)
}

func (s *EmployeeService) GetAll() ([]entities.Employee, error) {
	return s.repo.GetAllEmployees()
}

func NewEmployeeService(repo repository.EmployeeRepo) *EmployeeService {
	return &EmployeeService{repo: repo}
}

func (s *EmployeeService) Create(employee entities.Employee) (int, error) {
	return s.repo.CreateEmployee(employee)
}
