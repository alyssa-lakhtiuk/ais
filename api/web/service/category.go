package service

import (
	"ais/entities"
	"ais/repository"
)

type categoryService struct {
	repo repository.CategoryRepo
}

func (s *categoryService) Update(categoryNum int, category entities.CategoryInput) error {
	_, err := s.repo.GetCategoryByName(category.Name)
	if err != nil {
		return err
	}
	return s.repo.UpdateCategory(categoryNum, category)
}

func (s *categoryService) Delete(name string) error {
	_, err := s.repo.GetCategoryByName(name)
	if err != nil {
		return err
	}
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
