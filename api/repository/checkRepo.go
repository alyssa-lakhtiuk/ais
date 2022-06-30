package repository

import (
	"ais/entities"
	"github.com/jmoiron/sqlx"
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

func (er *checkPostgres) CreateCheck(check entities.Check) (int, error) {
	var id int
	row := er.db.QueryRow(createCheck, check.Number, check.PrintDate, check.SumTotal, check.Vat, check.IdEmployee, check.CardNumber)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (er *checkPostgres) UpdateCheck(numberCheck string, check entities.Check) error {
	_, err := er.db.Exec(updateCheck, numberCheck, check.PrintDate, check.SumTotal, check.Vat, check.IdEmployee, check.CardNumber)
	return err
	return nil
}

func (er *checkPostgres) DeleteCheck(name string) error {
	_, err := er.db.Exec(deleteCheck, name)
	return err
}

func (er *checkPostgres) GetCheckByNumber(num string) (entities.Check, error) {
	var check entities.Check
	if err := er.db.Get(&check, getCheckByName, num); err != nil {
		return entities.Check{}, err
	}
	return check, nil
}

func (er *checkPostgres) GetAllChecks() ([]entities.Check, error) {
	var checks []entities.Check
	if err := er.db.Select(&checks, getAllChecks); err != nil {
		return []entities.Check{}, err
	}
	return checks, nil
}
