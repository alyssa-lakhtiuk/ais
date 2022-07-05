package repository

import (
	"ais/entities"
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	priceByCategory = "SELECT category_name, SUM(product_number) " +
		"FROM (category INNER JOIN product ON category.category_number=product.fk_category_number) " +
		"INNER JOIN store_product ON store_product.fk_id_product=product.id_product " +
		"GROUP BY category_name"

	checksByCat = "SELECT bill_number " +
		"FROM sale SL " +
		"WHERE NOT EXISTS (SELECT * " +
		"FROM sale" +
		" WHERE upc IN (SELECT upc" +
		" FROM store_product " +
		"WHERE id_product NOT IN (SELECT id_product FROM product" +
		" WHERE category_number IN (SELECT category_number " +
		"FROM category WHERE category_name = $1))) AND SL.bill_number = bill_number);"

	countCities = "SELECT city, COUNT(*)" +
		"FROM customer_card INNER JOIN bill ON  customer_card.number = bill.card_number" +
		"GROUP BY customer_card.city"

	checksByPrice = "SELECT bill_number" +
		"FROM bill" +
		"WHERE NOT EXIST (SELECT *" +
		"FROM Sale" +
		"WHERE selling_ price < 250 AND Check.check_number = Sale.check_number); "
)

type zvit struct {
	db *sqlx.DB
}

func NewZvit(db *sqlx.DB) *zvit {
	return &zvit{db: db}
}

func (er *zvit) GetPricesByCategories() ([]entities.PriceByCat, error) {
	var ccs []entities.PriceByCat

	rows, err := er.db.Query(priceByCategory)
	if err != nil {
		return nil, fmt.Errorf("unable to execute the query")
	}
	defer rows.Close()
	for rows.Next() {
		cc := entities.PriceByCat{}
		err := rows.Scan(&cc.Cat, &cc.Price)
		if err != nil {
			return nil, err
		}
		ccs = append(ccs, cc)
	}
	return ccs, nil
}

func (er *zvit) GetChecksByCat(category string) ([]entities.CheckByCat, error) {
	var ccs []entities.CheckByCat

	rows, err := er.db.Query(checksByCat, category)
	if err != nil {
		return nil, fmt.Errorf("unable to execute the query")
	}
	defer rows.Close()
	for rows.Next() {
		cc := entities.CheckByCat{}
		err := rows.Scan(&cc.Check)
		if err != nil {
			return nil, err
		}
		ccs = append(ccs, cc)
	}
	return ccs, nil
}
