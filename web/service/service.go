package service

import (
	"ais/entities"
	"ais/repository"
)

type Authorization interface {
	CreateEmployee(employee entities.Employee) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Employee interface {
	Create(employee entities.Employee) (int, error)
	Update(employeeId string, employee entities.EmployeeInput) error
	Delete(employeeId string) error
	GetByName(employeeName string) (entities.Employee, error) // ?
	GetAll() ([]entities.Employee, error)
}

type Category interface {
	Create(category entities.Category) (int, error)
	Update(categoryNumber int, category entities.CategoryInput) error
	Delete(categoryName string) error
	GetByName(categoryName string) (entities.Category, error)
	GetAll() ([]entities.Category, error)
}

type Product interface {
	Create(product entities.Product) (int, error)
	Update(productId int, product entities.ProductInput) error
	Delete(productId int) error
	GetByName(productName string) (entities.Product, error) // ? add func getByID
	GetAll() ([]entities.Product, error)
}

type Service struct {
	//Authorization
	Employee
	Category
	Product
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Employee: NewEmployeeService(repos.EmployeeRepo),
		Category: NewCategoryService(repos.CategoryRepo),
		Product:  NewProductService(repos.ProductRepo),
	}
}
