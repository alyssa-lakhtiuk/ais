package repository

import (
	"ais/entities"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	createCustomerCard = "INSERT INTO " + customerCardTable + " (card_number, cust_surname, cust_name, cust_patronymic, " +
		"phone_number, city, street, zip_code, discount) " +
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);"
	updateCustomerCard = "UPDATE " + customerCardTable + " SET cust_surname=$2, cust_name=$3, cust_patronymic=$4, " +
		"phone_number=$5, city=$6, street=$7, zip_code=$8, discount=$9 " +
		"WHERE card_number=$1;"
	deleteCustomerCard      = "DELETE FROM " + customerCardTable + " WHERE card_number = $1;"
	getCustomerCardByNumber = "SELECT * FROM " + customerCardTable + " WHERE card_number=$1;"
	getCustomerCardByName   = "SELECT * FROM " + customerCardTable + " WHERE cust_name=$1;"
	getAllCustomerCards     = "SELECT * FROM " + customerCardTable + " ;"
)

type customerCardPostgres struct {
	db *sqlx.DB
}

func NewCustomerCardPostgres(db *sqlx.DB) *customerCardPostgres {
	return &customerCardPostgres{db: db}
}

func (er *customerCardPostgres) CreateCustomerCard(customerCard entities.CustomerCard) (int, error) {
	var id int
	row := er.db.QueryRow(createCustomerCard, customerCard.Number, customerCard.CustomerSurname, customerCard.CustomerName,
		customerCard.CustomerPatronymic, customerCard.PhoneNumber, customerCard.City, customerCard.Street,
		customerCard.ZipCode, customerCard.Percent)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (er *customerCardPostgres) UpdateCustomerCard(cardNumber string, customerCard entities.CustomerCard) error {
	_, err := er.db.Exec(updateCustomerCard, cardNumber, customerCard.CustomerSurname, customerCard.CustomerName,
		customerCard.CustomerPatronymic, customerCard.PhoneNumber, customerCard.City, customerCard.Street,
		customerCard.ZipCode, customerCard.Percent)
	return err
}

func (er *customerCardPostgres) DeleteCustomerCard(num string) error {
	_, err := er.db.Exec(deleteCustomerCard, num)
	return err
}

func (er *customerCardPostgres) GetCustomerCardByNumber(num string) (entities.CustomerCard, error) {
	var cc entities.CustomerCard
	row := er.db.QueryRow(getCustomerCardByNumber, num)
	err := row.Scan(&cc.Number, &cc.CustomerSurname, &cc.CustomerName, &cc.CustomerPatronymic, &cc.PhoneNumber,
		&cc.City, &cc.Street, &cc.ZipCode, &cc.Percent)
	if err != nil {
		return cc, err
	}
	return cc, nil
	//if err := er.db.Get(&cc, getCustomerCardByNumber, num); err != nil {
	//	return entities.CustomerCard{}, err
	//}
	//return cc, nil
}

func (er *customerCardPostgres) GetCustomerCardByName(name string) (entities.CustomerCard, error) {
	var cc entities.CustomerCard
	if err := er.db.Get(&cc, getCustomerCardByName, name); err != nil {
		return entities.CustomerCard{}, err
	}
	return cc, nil
}

func (er *customerCardPostgres) GetAllCustomerCards() ([]entities.CustomerCard, error) {
	var cc []entities.CustomerCard
	rows, err := er.db.Query(getAllCustomerCards)
	if err != nil {
		return nil, fmt.Errorf("unable to execute the query")
	}
	defer rows.Close()
	for rows.Next() {
		customerCard := entities.CustomerCard{}
		err := rows.Scan(&customerCard.Number, &customerCard.CustomerSurname, &customerCard.CustomerName,
			&customerCard.CustomerPatronymic, &customerCard.PhoneNumber, &customerCard.City,
			&customerCard.Street, &customerCard.ZipCode, &customerCard.Percent)
		if err != nil {
			return nil, err
		}
		cc = append(cc, customerCard)
	}
	return cc, nil
}
