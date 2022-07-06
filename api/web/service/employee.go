package service

import (
	"ais/entities"
	"ais/repository"
	"fmt"
)

type EmployeeService struct {
	repo     repository.EmployeeRepo
	repoRole repository.RoleRepo
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

func (s *EmployeeService) GetById(emplId string) (entities.Employee, error) {
	return s.repo.GetEmployeeById(emplId)
}

func (s *EmployeeService) GetAll() ([]entities.Employee, error) {
	return s.repo.GetAllEmployees()
}

func (s *EmployeeService) GetAllByCategory(role string) ([]entities.Employee, error) {
	return s.repo.GetEmployeeByRole(role)
}

func NewEmployeeService(repo repository.EmployeeRepo, repoRole repository.RoleRepo) *EmployeeService {
	return &EmployeeService{repo: repo, repoRole: repoRole}
}

func (s *EmployeeService) Create(employee entities.Employee) (int, error) {
	for true {
		emplId := GenerateRandomStr(10)
		cat, _ := s.repo.GetEmployeeById(emplId)
		if cat.ID != emplId {
			employee.ID = emplId
			break
		}
	}
	password := generatePasswordHash("secret")
	s.repoRole.CreateUserRole(password, employee.ID, employee.Role, employee.PhoneNumber)
	//if time.Now().Year()-employee.DateOfBirth.Year() < 18 {
	//	return 0, fmt.Errorf("employee must be adult")
	//}
	phone := employee.PhoneNumber
	err := ValidPhone(phone)
	if err != nil {
		return 0, fmt.Errorf("employee phone is invalid")
	}
	err = IsUnsigned(int(employee.Salary))
	if err != nil {
		return 0, fmt.Errorf("must employee pay us")
	}
	return s.repo.CreateEmployee(employee)
}
