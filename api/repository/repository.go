package repository

import (
	"ais/entities"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// tableNames:
const (
	employeeTable     = "employee"
	categoryTable     = "category"
	productTable      = "product"
	storeProductTable = "store_product"
	saleTable         = "sale"
	customerCardTable = "customer_card"
	checkTable        = "bill"
	rolesTable        = "roles"
)

// all repository interfaces:

type EmployeeRepo interface {
	CreateEmployee(employee entities.Employee) (int, error)
	UpdateEmployee(idEmployee string, employee entities.EmployeeInput) error
	DeleteEmployee(id string) error
	GetEmployeeByName(name string) (entities.Employee, error)
	GetEmployeeById(id string) (entities.Employee, error)
	GetAllEmployees() ([]entities.Employee, error)
	GetEmployeeByRole(role string) ([]entities.Employee, error)
}

type ProductRepo interface {
	CreateProduct(product entities.Product) (int, error)
	UpdateProduct(idProduct int, product entities.Product) error
	DeleteProduct(productId int) error
	GetProductByName(name string) (entities.Product, error)
	GetProductByNumber(number int) (entities.Product, error)
	GetAllProducts() ([]entities.Product, error)
}

type StoreProductRepo interface {
	CreateStoreProduct(product entities.StoreProduct) (int, error)
	UpdateStoreProduct(upc string, product entities.StoreProduct) error
	DeleteStoreProduct(upc string) error
	GetStoreProductByUpc(name string) (entities.StoreProduct, error)
	GetAllStoreProducts() ([]entities.StoreProduct, error)
}

type CategoryRepo interface {
	CreateCategory(category entities.Category) (int, error)
	UpdateCategory(categoryNum int, category entities.CategoryInput) error
	DeleteCategory(name string) error
	DeleteCategoryByNum(name int) error
	GetCategoryByName(categoryName string) (entities.Category, error)
	GetCategoryByNumber(categoryNumber int) (entities.Category, error)
	GetAllCategories() ([]entities.Category, error)
}

type CheckRepo interface {
	CreateCheck(randomStr string, checkInput []entities.CheckInput) (int, error)
	DeleteCheck(name string) error
	GetCheckByNumber(checkId string) (entities.Check, error)
	GetAllChecks() ([]entities.Check, error)
}

type CustomerCardRepo interface {
	CreateCustomerCard(cc entities.CustomerCard) (int, error)
	UpdateCustomerCard(ccId string, category entities.CustomerCard) error
	DeleteCustomerCard(ccId string) error
	GetCustomerCardByName(ccId string) (entities.CustomerCard, error)
	GetCustomerCardByNumber(num string) (entities.CustomerCard, error)
	GetAllCustomerCards() ([]entities.CustomerCard, error)
}

type SaleRepo interface {
	CreateSale(sale entities.Sale) (int, error)
	DeleteSale(upc string, checkNumber string) error
	GetSaleByUpcCheck(upc, checkNumber string) (entities.Sale, error)
	GetAllSales() ([]entities.Sale, error)
}

type RoleRepo interface {
	CreateUserRole(password string, emplId string, role string, phone string) (int, error)
	GetRoleByPhone(phone string) (entities.SignIn, error)
	GetRoleByIdEmployee(id string) (entities.SignIn, error)
}

type Zvit interface {
	GetQuantitiesByCategories() ([]entities.QuantityByCat, error)
	GetChecksByCat(category string) ([]entities.CheckByCat, error)
	/////
	CountCities() ([]entities.CountCustomersCities, error)
	ChecksByPrice(price int) ([]entities.SecondStruct, error)
}

type Repository struct {
	EmployeeRepo
	ProductRepo
	StoreProductRepo
	CategoryRepo
	CheckRepo
	CustomerCardRepo
	SaleRepo
	RoleRepo
	Zvit
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		EmployeeRepo:     NewEmployeePostgres(db),
		ProductRepo:      NewProductPostgres(db),
		StoreProductRepo: NewStoreProductRepo(db),
		CategoryRepo:     NewCategoryRepo(db),
		CheckRepo:        NewCheckRepo(db),
		CustomerCardRepo: NewCustomerCardPostgres(db),
		SaleRepo:         NewSalePostgres(db),
		RoleRepo:         NewRolesPostgres(db),
		Zvit:             NewZvit(db),
	}
}
