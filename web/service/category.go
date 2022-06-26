package service

import (
	"ais/entities"
	"ais/repository"
)

type CategoryService struct {
	repo repository.CategoryRepo
}

func (s *CategoryService) Update(categoryNum int, employee entities.CategoryInput) error {
	return s.repo.UpdateCategory(categoryNum, employee)
}

func (s *CategoryService) Delete(name string) error {
	return s.repo.DeleteCategory(name)
}

func (s *CategoryService) GetByName(categoryName string) (entities.Category, error) {
	return s.repo.GetCategoryByName(categoryName)
}

func (s *CategoryService) GetAll() ([]entities.Category, error) {
	return s.repo.GetAllCategories()
}

func NewCategoryService(repo repository.CategoryRepo) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) Create(category entities.Category) (int, error) {
	return s.repo.CreateCategory(category)
}
