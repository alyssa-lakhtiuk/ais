package repository

import (
	"ais/entities"
	"github.com/jmoiron/sqlx"
)

const (
	createStoreProduct = "INSERT INTO " + storeProductTable + " (upc, selling_price, promotional_product, product_number," +
		" upc_prom, fk_id_product) " +
		"VALUES ($1, $2, $3, $4, $5, $6);"
	createStoreProductWithoutUpc = "INSERT INTO " + storeProductTable + " (upc, selling_price, promotional_product, product_number," +
		" fk_id_product) " +
		"VALUES ($1, $2, $3, $4, $5);"
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
	//var id int
	if product.UPCProm == "" {
		_, err := er.db.Exec(createStoreProductWithoutUpc, product.UPC, product.SellingPrice, product.PromotionalProduct,
			product.ProductsNumber, product.IDProduct)
		if err != nil {
			return 0, err
		}
	} else {
		_, err := er.db.Exec(createStoreProduct, product.UPC, product.SellingPrice, product.PromotionalProduct,
			product.ProductsNumber, product.UPCProm, product.IDProduct)
		if err != nil {
			return 0, err
		}
	}
	//row := er.db.QueryRow(createStoreProduct, product.UPC, product.SellingPrice, product.PromotionalProduct,
	//	product.ProductsNumber, product.UPCProm, product.IDProduct)
	//if err := row.Scan(&id); err != nil {
	//	return 0, err
	//}

	return 1, nil
}

func (er *storeProductPostgres) UpdateStoreProduct(upc string, product entities.StoreProduct) error {
	_, err := er.db.Exec(updateStoreProduct, upc, product.SellingPrice, product.PromotionalProduct,
		product.ProductsNumber, product.UPC, product.IDProduct)
	return err
	return nil
}

func (er *storeProductPostgres) DeleteStoreProduct(upc string) error {
	_, err := er.db.Exec(deleteStoreProduct, upc)
	return err
}

func (er *storeProductPostgres) GetStoreProductByUpc(upc string) (entities.StoreProduct, error) {
	var stProduct entities.StoreProduct
	if err := er.db.Get(&stProduct, getStoreProductByName, upc); err != nil {
		return entities.StoreProduct{}, err
	}
	return stProduct, nil
}

func (er *storeProductPostgres) GetAllStoreProducts() ([]entities.StoreProduct, error) {
	var stProducts []entities.StoreProduct

	rows, err := er.db.Query(getAllChecks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		stProd := entities.StoreProduct{}
		err := rows.Scan(&stProd.UPC, &stProd.SellingPrice, &stProd.PromotionalProduct, &stProd.ProductsNumber,
			&stProd.UPCProm, &stProd.IDProduct)
		if err != nil {
			return nil, err
		}
		stProducts = append(stProducts, stProd)
	}
	//if err := er.db.Select(&stProducts, getAllStoreProducts); err != nil {
	//	return []entities.StoreProduct{}, err
	//}
	return stProducts, nil
}
