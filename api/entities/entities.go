package entities

import (
	"database/sql"
	_ "github.com/lib/pq"
	"time"
)

type Employee struct {
	ID          string         //`json:"id_employee" validate:"required"`
	SurName     string         `json:"lastname" validate:"required"`
	FirstName   string         `json:"firstname" validate:"required"`
	Patronymic  sql.NullString `json:"patronymic"`
	Role        string         `json:"role" validate:"required"`
	Salary      float64        `json:"salary" validate:"required"`
	DateOfBirth string         `json:"DateOfBirth" validate:"required"`
	DateOfStart string         `json:"DateOfStart" validate:"required"`
	PhoneNumber string         `json:"phoneNumber" validate:"required"`
	City        string         `json:"city" validate:"required"`
	Street      string         `json:"street" validate:"required"`
	ZipCode     string         `json:"zipCode" validate:"required"`
	//Password    string `json:"password" validate:"required"`
	// or separate db relation for password
}

type Product struct {
	Id              int    `json:"id" validate:"required"`
	CategoryNum     int    `json:"categoryNum" validate:"required"`
	Name            string `json:"name" validate:"required"`
	Characteristics string `json:"Characteristics" validate:"required"`
}

type StoreProduct struct {
	UPC                string         `json:"upc" validate:"required"`
	SellingPrice       float64        `json:"selling_price" validate:"required"`
	PromotionalProduct bool           `json:"promotional_product" validate:"required"`
	UPCProm            sql.NullString `json:"UPCProm" validate:"required"`
	IDProduct          int            `json:"IDProduct" validate:"required"`
	ProductsNumber     int            `json:"products_number" validate:"required"`
}

type Category struct {
	Number int    `json:"number"`
	Name   string `json:"name" validate:"required"`
}

type Check struct {
	Number     string         `json:"number" validate:"required"`
	IdEmployee string         `json:"idEmployee" validate:"required"`
	CardNumber sql.NullString `json:"cardNumber"`
	PrintDate  time.Time      `json:"printDate" validate:"required"`
	SumTotal   float64        `json:"sumTotal" validate:"required"`
	Vat        float64        `json:"vat"`
}

type CustomerCard struct {
	Number             string         `json:"number"`
	CustomerSurname    string         `json:"customerSurname" validate:"required"`
	CustomerName       string         `json:"customerName" validate:"required"`
	CustomerPatronymic sql.NullString `json:"customerPatronymic" validate:"required"`
	PhoneNumber        string         `json:"phoneNumber" validate:"required"`
	City               sql.NullString `json:"city"`
	Street             sql.NullString `json:"street"`
	ZipCode            sql.NullString `json:"zipCode"`
	Percent            int            `json:"percent" validate:"required"`
}

type Sale struct {
	UPC           string  `json:"upc" validate:"required"`
	CheckNumber   string  `json:"checkNumber" validate:"required"`
	ProductNumber int     `json:"productNumber" validate:"required"`
	SellingPrice  float64 `json:"sellingPrice" validate:"required"`
}

// entities for update:

type EmployeeInput struct {
	FirstName   string         `json:"firstname"`
	SurName     string         `json:"lastname"`
	Patronymic  sql.NullString `json:"patronymic"`
	Role        string         `json:"role"`
	Salary      float64        `json:"salary"`
	DateOfBirth string         `json:"DateOfBirth"`
	DateOfStart string         `json:"DateOfStart"`
	PhoneNumber string         `json:"phoneNumber"`
	City        string         `json:"city"`
	Street      string         `json:"street"`
	ZipCode     string         `json:"zipCode"`
}

type CategoryInput struct {
	Name string `json:"name" validate:"required"`
}

type ProductInput struct {
	CategoryNum     int    `json:"categoryNum"`
	Name            string `json:"name"`
	Characteristics string `json:"Characteristics"`
}

type CheckInput struct {
	// те, що вводить касир при створенні чеку
	UPC            string `json:"upc" validate:"required"`
	ProductNumber  int    `json:"productNumber" validate:"required"`
	CustomerNumber string `json:"CustomerNumber"`
	IDEmployee     string `json:"id_employee" validate:"required"`
}

type SignIn struct {
	IdEmployee string
	Role       string
	Phone      string
	Password   string
}

type Message struct {
	Mess string
}

type CategoryProduct struct {
	C []Category
	P Product
}

type ProductStoreProductUpc struct {
	Pr  []Product
	Stp StoreProduct
	UpC []StoreProduct
}

type QuantityByCat struct {
	Cat      string  `json:"cat" validate:"required"`
	Quantity float64 `json:"quantity" validate:"required"`
}

type CheckByCat struct {
	Check string `json:"check" validate:"required"`
	Cat   string `json:"cat"`
}

type CountCustomersCities struct {
	City            string
	CheckCardNumber int
}

type SecondStruct struct {
	Check string
}
