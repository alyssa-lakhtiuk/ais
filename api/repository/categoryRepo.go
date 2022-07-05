package repository

import (
	"ais/entities"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	createCategory = "INSERT INTO " + categoryTable + " (category_number, category_name) " +
		"VALUES ($1, $2);"
	updateCategory = "UPDATE " + categoryTable + " SET category_name=$2 " +
		"WHERE category_number=$1;"
	deleteCategory      = "DELETE FROM " + categoryTable + " WHERE category_name = $1;"
	deleteCategoryByNum = "DELETE FROM " + categoryTable + " WHERE category_number = $1;"
	getCategoryByName   = "SELECT * FROM " + categoryTable + " WHERE category_name=$1;"
	getCategoryByNumber = "SELECT * FROM " + categoryTable + " WHERE category_number=$1;"
	getAllCategories    = "SELECT * FROM " + categoryTable + " ;"
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
}

func (er *categoryPostgres) DeleteCategory(name string) error {
	_, err := er.db.Exec(deleteCategory, name)
	return err
}

func (er *categoryPostgres) DeleteCategoryByNum(name int) error {
	_, err := er.db.Exec(deleteCategoryByNum, name)
	return err
}

func (er *categoryPostgres) GetCategoryByName(name string) (entities.Category, error) {
	var category entities.Category
	row := er.db.QueryRow(getCategoryByName, name)
	err := row.Scan(&category.Number, &category.Name)
	if err != nil {
		return category, err
	}
	return category, nil
	//if err := er.db.Get(&category, getCategoryByName, name); err != nil {
	//	return entities.Category{}, fmt.Errorf("such category doesn't exist")
	//}
	//return category, nil
}

func (er *categoryPostgres) GetCategoryByNumber(number int) (entities.Category, error) {
	var category entities.Category
	row := er.db.QueryRow(getCategoryByNumber, number)
	err := row.Scan(&category.Number, &category.Name)
	if err != nil {
		return category, err
	}
	return category, nil
	//if err := er.db.Get(&category, getCategoryByNumber, number); err != nil {
	//	return entities.Category{}, fmt.Errorf("such category doesn't exist")
	//}
	//return category, nil
}

func (er *categoryPostgres) GetAllCategories() ([]entities.Category, error) {
	var categories []entities.Category
	rows, err := er.db.Query(getAllCategories)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		category := entities.Category{}
		err := rows.Scan(&category.Number, &category.Name)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	//if err := er.db.Select(&categories, getAllCategories); err != nil {
	//	return []entities.Category{}, err
	//}

	return categories, nil
}
