package service

import (
	"ais/entities"
	"ais/repository"
	"time"
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
	if time.Now().Year()-employee.DateOfBirth.Year() < 18 {
		// return err "Employee must be adult"
	}
	phone := employee.PhoneNumber
	err := ValidPhone(phone)
	if err != nil {
		return 0, err
	}
	err = IsUnsigned(int(employee.Salary))
	if err != nil {
		return 0, err
	}
	return s.repo.CreateEmployee(employee)
}
