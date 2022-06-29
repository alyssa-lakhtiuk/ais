package entities

import (
	"time"
)

type Employee struct {
	ID          string    //`json:"id_employee" validate:"required"`
	FirstName   string    `json:"firstname" validate:"required"`
	SurName     string    `json:"lastname" validate:"required"`
	Patronymic  string    `json:"patronymic" validate:"required"`
	Role        string    `json:"role" validate:"required"`
	Salary      int       `json:"salary" validate:"required"`
	DateOfBirth time.Time `json:"DateOfBirth" validate:"required"`
	DateOfStart time.Time `json:"DateOfStart" validate:"required"`
	PhoneNumber string    `json:"phoneNumber" validate:"required"`
	City        string    `json:"city" validate:"required"`
	Street      string    `json:"street" validate:"required"`
	ZipCode     string    `json:"zipCode" validate:"required"`
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
	UPC                string  `json:"upc" validate:"required"`
	SellingPrice       float64 `json:"selling_price" validate:"required"`
	PromotionalProduct bool    `json:"promotional_product" validate:"required"`
	UPCProm            string  `json:"UPCProm" validate:"required"`
	IDProduct          int     `json:"IDProduct" validate:"required"`
	ProductsNumber     int     `json:"products_number" validate:"required"`
}

type Category struct {
	Number int    `json:"number" validate:"required"`
	Name   string `json:"name" validate:"required"`
}

type Check struct {
	Number     int       `json:"number" validate:"required"`
	IdEmployee string    `json:"idEmployee" validate:"required"`
	CardNumber string    `json:"cardNumber" validate:"required"`
	PrintDate  time.Time `json:"printDate" validate:"required"`
	SumTotal   int       `json:"sumTotal" validate:"required"`
	Vat        int       `json:"vat" validate:"required"`
}

type CustomerCard struct {
	Number             int    `json:"number" validate:"required"`
	CustomerSurname    string `json:"customerSurname" validate:"required"`
	CustomerName       string `json:"customerName" validate:"required"`
	CustomerPatronymic string `json:"customerPatronymic" validate:"required"`
	PhoneNumber        string `json:"phoneNumber" validate:"required"`
	City               string `json:"city" validate:"required"`
	Street             string `json:"street" validate:"required"`
	ZipCode            string `json:"zipCode" validate:"required"`
	Percent            int    `json:"percent" validate:"required"`
}

type Sale struct {
	UPC           string `json:"upc" validate:"required"`
	CheckNumber   string `json:"checkNumber" validate:"required"`
	ProductNumber int    `json:"productNumber" validate:"required"`
	SellingPrice  int    `json:"sellingPrice" validate:"required"`
}

// entities for update:

type EmployeeInput struct {
	FirstName   string    `json:"firstname"`
	SurName     string    `json:"lastname"`
	Patronymic  string    `json:"patronymic"`
	Role        string    `json:"role"`
	Salary      int       `json:"salary"`
	DateOfBirth time.Time `json:"DateOfBirth"`
	DateOfStart time.Time `json:"DateOfStart"`
	PhoneNumber string    `json:"phoneNumber"`
	City        string    `json:"city"`
	Street      string    `json:"street"`
	ZipCode     string    `json:"zipCode"`
}

type CategoryInput struct {
	Name string `json:"name" validate:"required"`
}

type ProductInput struct {
	CategoryNum     int    `json:"categoryNum"`
	Name            string `json:"name"`
	Characteristics string `json:"Characteristics"`
}
