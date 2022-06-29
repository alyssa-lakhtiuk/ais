package repository

import (
	"ais/entities"
	"github.com/jmoiron/sqlx"
)

const (
	createStoreProduct = "INSERT INTO " + storeProductTable + " (upc, selling_price, promotional_product, product_number," +
		" upc_prom, fk_id_product) " +
		"VALUES ($1, $2, $3, $4, $5, $6);"
	updateStoreProduct = "UPDATE " + storeProductTable + " SET selling_price=$2, promotional_product=$3, product_number=$4," +
		" upc_prom=$5, fk_id_product=$6 " +
		"WHERE upc=$1;"
	deleteStoreProduct    = "DELETE FROM " + storeProductTable + " WHERE upc = $1;"
	getStoreProductByName = "SELECT * FROM " + storeProductTable + " WHERE upc=$1;"
	getAllStoreProducts   = "SELECT * FROM " + storeProductTable + ";"
)

type storeProductPostgres struct {
	db *sqlx.DB
}

func NewStoreProductRepo(db *sqlx.DB) *storeProductPostgres {
	return &storeProductPostgres{db: db}
}

func (er *storeProductPostgres) CreateStoreProduct(product entities.StoreProduct) (int, error) {
	var id int
	row := er.db.QueryRow(createStoreProduct, product.UPC, product.SellingPrice, product.PromotionalProduct,
		product.ProductsNumber, product.UPCProm, product.IDProduct)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (er *storeProductPostgres) UpdateStoreProduct(upc string, product entities.StoreProduct) error {
	_, err := er.db.Exec(updateStoreProduct, upc, product.SellingPrice, product.PromotionalProduct, product.ProductsNumber, product.UPC)
	return err
	return nil
}

func (er *storeProductPostgres) DeleteStoreProduct(id int) error {
	_, err := er.db.Exec(deleteStoreProduct, id)
	return err
}

func (er *storeProductPostgres) GetStoreProductByName(name string) (entities.StoreProduct, error) {
	var st_product entities.StoreProduct
	if err := er.db.Get(&st_product, getStoreProductByName, name); err != nil {
		return entities.StoreProduct{}, err
	}
	return st_product, nil
}

func (er *storeProductPostgres) GetAllStoreProducts() ([]entities.StoreProduct, error) {
	var st_products []entities.StoreProduct
	if err := er.db.Select(&st_products, getAllStoreProducts); err != nil {
		return []entities.StoreProduct{}, err
	}
	return st_products, nil
}
