package repository

import (
	"ais/entities"
	"github.com/jmoiron/sqlx"
)

const (
	createCategory = "INSERT INTO " + categoryTable + " (category_number, category_name) " +
		"VALUES ($1, $2);"
	updateCategory = "UPDATE " + categoryTable + " SET category_name=$2 " +
		"WHERE category_number=$1;"
	deleteCategory    = "DELETE FROM " + categoryTable + " WHERE category_name = $1;"
	getCategoryByName = "SELECT * FROM " + categoryTable + " WHERE category_name=$1;"
	getAllCategories  = "SELECT * FROM " + categoryTable + ";"
)

type categoryPostgres struct {
	db *sqlx.DB
}

func NewCategoryRepo(db *sqlx.DB) *categoryPostgres {
	return &categoryPostgres{db: db}
}

func (er *categoryPostgres) CreateCategory(category entities.Category) (int, error) {
	var id int
	row := er.db.QueryRow(createCategory, category.Number, category.Name)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (er *categoryPostgres) UpdateCategory(numberCategory int, category entities.CategoryInput) error {
	_, err := er.db.Exec(updateCategory, numberCategory, category.Name)
	return err
	return nil
}

func (er *categoryPostgres) DeleteCategory(name string) error {
	_, err := er.db.Exec(deleteCategory, name)
	return err
}

func (er *categoryPostgres) GetCategoryByName(name string) (entities.Category, error) {
	var employee entities.Category
	if err := er.db.Get(&employee, getCategoryByName, name); err != nil {
		return entities.Category{}, err
	}
	return employee, nil
}

func (er *categoryPostgres) GetAllCategories() ([]entities.Category, error) {
	var categories []entities.Category
	if err := er.db.Select(&categories, getAllCategories); err != nil {
		return []entities.Category{}, err
	}
	return categories, nil
}
