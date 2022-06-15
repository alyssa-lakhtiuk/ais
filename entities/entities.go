package entities

import (
	"time"
)

type Employee struct {
	ID          string
	FirstName   string `json:"firstname" validate:"required"`
	SurName     string `json:"lastname" validate:"required"`
	Patronymic  string `json:"lastname" validate:"required"`
	Role        string `json:"role" validate:"required"`
	Salary      int
	DateOfBirth time.Time
	DateOfStart time.Time
	PhoneNumber string
	City        string
	Street      string
	ZipCode     string
	//Password    string `json:"password" validate:"required"`
}

type Product struct {
	Id              int
	CategoryNum     int
	Name            string
	Characteristics string
}

type StoreProduct struct {
	UPC       string
	UPCProm   string
	IDProduct int
}

type Category struct {
	Number int
	Name   string
}

type Check struct {
	Number     int
	IdEmployee string
	CardNumber string
	PrintDate  time.Time
	SumTotal   int
	Vat        int
}

type CustomerCard struct {
	Number             int
	CustomerSurname    string
	CustomerName       string
	CustomerPatronymic string
	PhoneNumber        string
	City               string
	Street             string
	ZipCode            string
	Percent            int
}

type Sale struct {
	UPC           string
	CheckNumber   string
	ProductNumber int
	SellingPrice  int
}
