package repository

import (
	"ais/entities"
	"github.com/jmoiron/sqlx"
)

// tableNames:
const (
	employeeTable     = "Employee"
	categoryTable     = "Category"
	productTable      = "Product"
	storeProductTable = "Store_Product"
	saleTable         = "Sale"
	customerCardTable = "Customer_Card"
	checkTable        = "Check"
)

// all repository interfaces:

type EmployeeRepo interface {
	CreateEmployee(employee entities.Employee) (int, error)
	UpdateEmployee(idEmployee string, employee entities.EmployeeInput) error
	DeleteEmployee(id string) error
	GetEmployeeByName(name string) (entities.Employee, error)
	GetAllEmployees() ([]entities.Employee, error)
}

type ProductRepo interface {
	CreateProduct(product entities.Product) (int, error)
	UpdateProduct(idProduct int, product entities.ProductInput) error
	DeleteProduct(productId int) error
	GetProductByName(name string) (entities.Product, error)
	GetAllProducts() ([]entities.Product, error)
}

type StoreProductRepo interface {
	createProduct(product entities.Product) (int, error)
	updateProduct(idProduct string, product entities.Product) error
	deleteProduct(idProduct string) error
	getProductByID(idProduct string) (entities.Product, error)
	getAllProducts() ([]entities.Product, error)
}

type CategoryRepo interface {
	CreateCategory(category entities.Category) (int, error)
	UpdateCategory(categoryNum int, category entities.CategoryInput) error
	DeleteCategory(name string) error
	GetCategoryByName(categoryName string) (entities.Category, error)
	GetAllCategories() ([]entities.Category, error)
}

type CheckRepo interface {
	createCheck(check entities.Check) (int, error)
	deleteCheck(check entities.Check) error
	getCheckByNumber(checkId string) (entities.Check, error)
	getAllChecks() ([]entities.Check, error)
}

type CustomerCardRepo interface {
	createCustomerCard(cc entities.CustomerCard) (int, error)
	updateCustomerCard(ccId string, category entities.CustomerCard) error
	deleteCustomerCard(ccId string) error
	getCustomerCardByName(ccId string) (entities.CustomerCard, error)
	getAllCustomerCards() ([]entities.CustomerCard, error)
}

type Repository struct {
	EmployeeRepo
	ProductRepo
	CategoryRepo
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		EmployeeRepo: NewEmployeePostgres(db),
		ProductRepo:  NewProductPostgres(db),
		CategoryRepo: NewCategoryRepo(db),
	}
}
