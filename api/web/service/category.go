package service

import (
	"ais/entities"
	"ais/repository"
	"math/rand"
)

type categoryService struct {
	repo repository.CategoryRepo
}

const IdRange = 100

func (s *categoryService) Update(categoryNum int, category entities.CategoryInput) error {
	//_, err := s.repo.GetCategoryByName(category.Name)
	//if err != nil {
	//	return fmt.Errorf("there is no such category")
	//}
	return s.repo.UpdateCategory(categoryNum, category)
}

func (s *categoryService) Delete(name string) error {
	//_, err := s.repo.GetCategoryByName(name)
	//if err != nil {
	//	return err
	//}
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
	for true {
		randId := rand.Intn(IdRange)
		cat, _ := s.repo.GetCategoryByNumber(randId)
		if cat.Number != randId {
			category.Number = randId
			break
		}
	}
	//category.Number = 97
	return s.repo.CreateCategory(category)
}
