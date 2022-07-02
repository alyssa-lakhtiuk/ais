package repository

import (
	"ais/entities"
	"github.com/jmoiron/sqlx"
)

const (
	createSale = "INSERT INTO " + saleTable + " (fk_upc, fk_check_number, product_number, selling_price) " +
		"VALUES ($1, $2, $3, $4);"
	updateSale = "UPDATE " + saleTable + " SET product_number=$3, selling_price=$4 " +
		"WHERE fk_upc=$1 AND fk_check_number=$2;"
	deleteSale           = "DELETE FROM " + saleTable + " WHERE fk_upc = $1 AND fk_check_number=$2;"
	getSaleByUpcCheckNum = "SELECT * FROM " + saleTable + " WHERE fk_upc = $1 AND fk_check_number=$2;"
	getAllSales          = "SELECT * FROM " + saleTable + ";"
)

type SalePostgres struct {
	db *sqlx.DB
}

func (p *SalePostgres) CreateSale(sale entities.Sale) (int, error) {
	var id int
	row := p.db.QueryRow(createSale, sale.UPC, sale.CheckNumber, sale.ProductNumber, sale.SellingPrice)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (p *SalePostgres) UpdateSale(upc string, checkNumber string, sale entities.Sale) error {
	_, err := p.db.Exec(updateSale, upc, checkNumber, sale.ProductNumber, sale.SellingPrice)
	return err
}

func (p *SalePostgres) DeleteSale(upc string, checkNumber string) error {
	_, err := p.db.Exec(deleteSale, upc, checkNumber)
	return err
}

func (p *SalePostgres) GetSaleByUpcCheck(upc, checkNumber string) (entities.Sale, error) {
	var product entities.Sale
	if err := p.db.Get(&product, getSaleByUpcCheckNum, upc, checkNumber); err != nil {
		return entities.Sale{}, err
	}
	return product, nil
}

func (p *SalePostgres) GetAllSales() ([]entities.Sale, error) {
	var sales []entities.Sale
	rows, err := p.db.Query(getAllChecks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		sale := entities.Sale{}
		err := rows.Scan(&sale.UPC, &sale.CheckNumber, &sale.ProductNumber, &sale.SellingPrice)
		if err != nil {
			return nil, err
		}
		sales = append(sales, sale)
	}
	//if err := p.db.Select(&sales, getAllSales); err != nil {
	//	return []entities.Sale{}, err
	//}
	return sales, nil
}

func NewSalePostgres(db *sqlx.DB) *SalePostgres {
	return &SalePostgres{db: db}
}
