package repository

import (
	"ais/entities"
	"github.com/jmoiron/sqlx"
)

const (
	createProduct = "INSERT INTO " + productTable + " (id_product, category_number, product_name, description) " +
		"VALUES ($1, $2, $3, $4);"
	updateProduct = "UPDATE " + productTable + " SET category_number=$2, product_name=$3, description=$4 " +
		"WHERE id_product=$1;"
	deleteProduct      = "DELETE FROM " + productTable + " WHERE id_product = $1;"
	getProductByName   = "SELECT * FROM " + productTable + " WHERE product_name=$1;"
	getProductByNumber = "SELECT * FROM " + productTable + " WHERE id_product=$1;"
	getAllProducts     = "SELECT * FROM " + productTable + ";"
)

type ProductPostgres struct {
	db *sqlx.DB
}

func (p *ProductPostgres) CreateProduct(product entities.Product) (int, error) {
	var id int
	row := p.db.QueryRow(createProduct, product.Id, product.CategoryNum, product.Name, product.Characteristics)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (p *ProductPostgres) UpdateProduct(idProduct int, product entities.Product) error {
	_, err := p.db.Exec(updateProduct, idProduct, product.CategoryNum, product.Name, product.Characteristics)
	return err
}

func (p *ProductPostgres) DeleteProduct(productId int) error {
	_, err := p.db.Exec(deleteProduct, productId)
	return err
}

func (p *ProductPostgres) GetProductByName(name string) (entities.Product, error) {
	var product entities.Product
	if err := p.db.Get(&product, getProductByName, name); err != nil {
		return entities.Product{}, err
	}
	return product, nil
}

func (p *ProductPostgres) GetProductByNumber(number int) (entities.Product, error) {
	var product entities.Product
	if err := p.db.Get(&product, getProductByNumber, number); err != nil {
		return entities.Product{}, err
	}
	return product, nil
}

func (p *ProductPostgres) GetAllProducts() ([]entities.Product, error) {
	var products []entities.Product
	rows, err := p.db.Query(getAllChecks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		prod := entities.Product{}
		err := rows.Scan(&prod.Id, &prod.Name, &prod.Characteristics, &prod.CategoryNum)
		if err != nil {
			return nil, err
		}
		products = append(products, prod)
	}
	//if err := p.db.Select(&products, getAllProducts); err != nil {
	//	return []entities.Product{}, err
	//}
	return products, nil
}

func NewProductPostgres(db *sqlx.DB) *ProductPostgres {
	return &ProductPostgres{db: db}
}
