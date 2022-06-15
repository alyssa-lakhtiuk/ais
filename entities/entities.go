package entities

import "github.com/google/uuid"

type Employee struct {
	ID        uuid.UUID
	FirstName string `json:"firstname" validate:"required"`
	LastName  string `json:"lastname" validate:"required"`
	Password  string `json:"password" validate:"required"`
	Role      string `json:"role" validate:"required"`
}

type Product struct {
}

type StoreProduct struct {
}

type Category struct {
}

type Check struct {
}

type CustomerCard struct {
}

type Sale struct {
}
