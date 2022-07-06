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
	GetById(emplId string) (entities.Employee, error)
	GetAll() ([]entities.Employee, error)
	GetAllByCategory(role string) ([]entities.Employee, error)
}

type Category interface {
	Create(category entities.Category) (int, error)
	Update(categoryNumber int, category entities.CategoryInput) error
	Delete(categoryName string) error
	DeleteByNum(name int) error
	GetByName(categoryName string) (entities.Category, error)
	GetByNumber(categoryNumber int) (entities.Category, error)
	GetAll() ([]entities.Category, error)
}

type Product interface {
	Create(product entities.Product) (int, error)
	Update(productId int, product entities.Product) error
	Delete(productId int) error
	GetByName(productName string) (entities.Product, error) // ? add func getByID
	GetByNumber(productId int) (entities.Product, error)
	GetAll() ([]entities.Product, error)
}

type StoreProduct interface {
	Create(product entities.StoreProduct) (int, error)
	Update(upc string, product entities.StoreProduct) error
	Delete(upc string) error
	GetByName(productName string) (entities.StoreProduct, error) // ? add func getByID
	GetAll() ([]entities.StoreProduct, error)
}

type Sale interface {
	Create(sale entities.Sale) (int, error) // we don't need to update sale Update(categoryNumber int, category entities.CategoryInput) error
	Delete(upc, checkNum string) error
	GetByUpcCheck(upc, checkNum string) (entities.Sale, error)
	GetAll() ([]entities.Sale, error)
}

type CustomerCard interface {
	Create(cc entities.CustomerCard) (int, error)
	Delete(name string) error
	Update(ccNum string, cc entities.CustomerCard) error
	GetByNumber(ccNumber string) (entities.CustomerCard, error)
	GetAll() ([]entities.CustomerCard, error)
}

type Check interface {
	Create(check []entities.CheckInput) (int, error)
	Delete(num string) error
	// Update(ccNum string, cc entities.CustomerCard) error
	GetByNumber(num string) (entities.Check, error)
	GetAll() ([]entities.Check, error)
}

type Role interface {
	GetByPhone(phone string) (entities.SignIn, error)
	GetByIdEmployee(id string) (entities.SignIn, error)
}

type Zvit interface {
	GetPricesByCategory() ([]entities.PriceByCat, error)
	GetChecksByCategory(category string) ([]entities.CheckByCat, error)
	GetCountByCities() ([]entities.CountCustomersCities, error)
	GetChecksByPrice(price int) ([]entities.SecondStruct, error)
}

type Service struct {
	//Authorization
	Employee
	Category
	Product
	StoreProduct
	Sale
	CustomerCard
	Check
	Role
	Zvit
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Employee:     NewEmployeeService(repos.EmployeeRepo),
		Category:     NewCategoryService(repos.CategoryRepo),
		Product:      NewProductService(repos.ProductRepo, repos.CategoryRepo),
		StoreProduct: NewStoreProductService(repos.ProductRepo, repos.StoreProductRepo),
		Sale:         NewSaleService(repos.SaleRepo, repos.StoreProductRepo, repos.CheckRepo),
		CustomerCard: NewCustomerCardServiceService(repos.CustomerCardRepo),
		Check:        NewCheckService(repos.CheckRepo, repos.EmployeeRepo, repos.CustomerCardRepo, repos.StoreProductRepo),
		Role:         NewRoleService(repos.RoleRepo),
		Zvit:         NewZvitService(repos.Zvit),
	}
}
