package controller

import (
	"ais/repository"
)

// це тільки контролер, який колає відповідні сервіси!
// мейбі навіть зробити тут один файлик

// ввести відомості про новий товар
// оновити відомості про товар
// вилучити відомості про товар

type ProductService struct {
	repo repository.ProductRepo
}

func NewProductService(repo repository.ProductRepo) *ProductService {
	return &ProductService{repo: repo}
}

func (es *ProductService) Create() {

}
