package service

import (
	"ais/entities"
	"ais/repository"
)

type categoryService struct {
	repo repository.CategoryRepo
}

func (s *categoryService) Update(categoryNum int, employee entities.CategoryInput) error {
	return s.repo.UpdateCategory(categoryNum, employee)
}

func (s *categoryService) Delete(name string) error {
	return s.repo.DeleteCategory(name)
}

func (s *categoryService) GetByName(categoryName string) (entities.Category, error) {
	return s.repo.GetCategoryByName(categoryName)
}

func (s *categoryService) GetAll() ([]entities.Category, error) {
	return s.repo.GetAllCategories()
}

func NewCategoryService(repo repository.CategoryRepo) *categoryService {
	return &categoryService{repo: repo}
}

func (s *categoryService) Create(category entities.Category) (int, error) {
	return s.repo.CreateCategory(category)
}
