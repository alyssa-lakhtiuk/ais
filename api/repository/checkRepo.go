package repository

import (
	"ais/entities"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"time"
)

const (
	createCheck = "INSERT INTO " + checkTable + " (bill_number, print_date, sum_total, vat, fk_id_employee, fk_card_number) " +
		"VALUES ($1, $2, $3, $4, $5, $6);"
	updateCheck = "UPDATE " + checkTable + " SET print_date=$2, sum_total=$3, vat=$4, fk_id_employee=$5, fk_card_number=$6 " +
		"WHERE bill_number=$1;"
	deleteCheck    = "DELETE FROM " + checkTable + " WHERE bill_number = $1;"
	getCheckByName = "SELECT * FROM " + checkTable + " WHERE bill_number=$1;"
	getAllChecks   = "SELECT * FROM " + checkTable + ";"
)

type checkPostgres struct {
	db *sqlx.DB
}

func NewCheckRepo(db *sqlx.DB) *checkPostgres {
	return &checkPostgres{db: db}
}

func (er *checkPostgres) CreateCheck(randomStr string, checkInput []entities.CheckInput) (int, error) {
	var id int
	tx, err := er.db.Begin()
	if err != nil {
		return 0, err
	}
	var sumTotal float64
	for i := 0; i < len(checkInput); i++ {
		currentProductUPC := checkInput[i].UPC
		var stProduct entities.StoreProduct
		row := er.db.QueryRow(getStoreProductByName, currentProductUPC)
		err := row.Scan(&stProduct.UPC, &stProduct.SellingPrice, &stProduct.PromotionalProduct, &stProduct.ProductsNumber,
			&stProduct.UPCProm, &stProduct.IDProduct)
		if err != nil {
			return 0, err
		}
		sumTotal += float64(checkInput[i].ProductNumber) * stProduct.SellingPrice
	}
	//var chechInDB entities.Check
	//checkInDB := entities.Check{Number: randomStr, PrintDate: time.Now(), SumTotal: sumTotal,
	//	Vat: sumTotal * 0.2, CardNumber: checkInput[0].CustomerNumber, IdEmployee: checkInput[0].IDEmployee}
	_ = tx.QueryRow(createCheck, randomStr, time.Now(), sumTotal, sumTotal*0.2, checkInput[0].IDEmployee, checkInput[0].CustomerNumber)
	//err2 := row.Scan(&id)
	//if err2 != nil {
	//	tx.Rollback()
	//	return 0, err
	//}
	//for j := 0; j < len(checkInput); j++ {
	//	var stProduct entities.StoreProduct
	//	//if err := er.db.Get(&stProduct, getStoreProductByName, checkInput[j].UPC); err != nil {
	//	//	return 0, err
	//	//}
	//	row := er.db.QueryRow(getStoreProductByName, checkInput[j].UPC)
	//	err := row.Scan(&stProduct.UPC, &stProduct.SellingPrice, &stProduct.PromotionalProduct, &stProduct.ProductsNumber,
	//		&stProduct.UPCProm, &stProduct.IDProduct)
	//	saleInCheck := entities.Sale{UPC: checkInput[j].UPC, SellingPrice: stProduct.SellingPrice,
	//		CheckNumber: checkInDB.Number, ProductNumber: checkInput[j].ProductNumber}
	//	_, err = tx.Exec(createSale, saleInCheck.ProductNumber, saleInCheck.SellingPrice, saleInCheck.UPC, saleInCheck.CheckNumber)
	//	if err != nil {
	//		tx.Rollback()
	//		return 0, err
	//	}
	//}
	//var id int
	//row := er.db.QueryRow(createCheck, check.Number, check.PrintDate, check.SumTotal, check.Vat, check.IdEmployee, check.CardNumber)
	//if err := row.Scan(&id); err != nil {
	//	return 0, err
	//}
	return id, tx.Commit()
}

func (er *checkPostgres) UpdateCheck(numberCheck string, check entities.Check) error {
	// ????????????????????
	_, err := er.db.Exec(updateCheck, numberCheck, check.PrintDate, check.SumTotal, check.Vat, check.IdEmployee, check.CardNumber)
	return err
	return nil
}

func (er *checkPostgres) DeleteCheck(num string) error {
	_, err := er.db.Exec(deleteCheck, num)
	return err
}

func (er *checkPostgres) GetCheckByNumber(num string) (entities.Check, error) {
	var check entities.Check
	row := er.db.QueryRow(getCheckByName, num)
	err := row.Scan(&check.Number, &check.PrintDate, &check.SumTotal, &check.Vat, &check.IdEmployee, &check.CardNumber)
	if err != nil {
		return check, err
	}
	return check, nil
	//if err := er.db.Get(&check, getCheckByName, num); err != nil {
	//	return entities.Check{}, err
	//}
	//return check, nil
}

func (er *checkPostgres) GetAllChecks() ([]entities.Check, error) {
	var checks []entities.Check

	rows, err := er.db.Query(getAllChecks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		check := entities.Check{}
		err := rows.Scan(&check.Number, &check.PrintDate, &check.SumTotal, &check.Vat, &check.IdEmployee, &check.CardNumber)
		if err != nil {
			return nil, err
		}
		checks = append(checks, check)
	}
	//if err := er.db.Select(&checks, getAllChecks); err != nil {
	//	return []entities.Check{}, err
	//}
	return checks, nil
}
