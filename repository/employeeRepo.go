package repository

import (
	"ais/entities"
	"github.com/jmoiron/sqlx"
)

const (
	createEmployee = "INSERT INTO " + employeeTable + " (id_employee, empl_surname, empl_name, empl_patronymic, " +
		"empl_role, salary, date_of_birth, date_of_start, phone_number, city, street, zip_code) " +
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12);"
	updateEmployee = "UPDATE " + employeeTable + " SET empl_surname=$2, empl_name=$3, empl_patronymic=$4, " +
		"empl_role=$5, salary=$6, date_of_birth=$7, date_of_start=$8, phone_number=$9, city=$10, street=$11, zip_code=$12 " +
		"WHERE id_employee=$1;"
	deleteEmployee    = "DELETE FROM " + employeeTable + " WHERE id_employee = $1;"
	getEmployeeByName = "SELECT * FROM " + employeeTable + " WHERE empl_name=$1;"
	getAllEmployees   = "SELECT * FROM " + employeeTable + ";"
)

type EmployeePostgres struct {
	db *sqlx.DB
}

func NewEmployeePostgres(db *sqlx.DB) *EmployeePostgres {
	return &EmployeePostgres{db: db}
}

func (er *EmployeePostgres) CreateEmployee(employee entities.Employee) (int, error) {
	var id int
	row := er.db.QueryRow(createEmployee, employee.ID, employee.FirstName, employee.SurName, employee.Patronymic,
		employee.Role, employee.Salary, employee.DateOfBirth, employee.DateOfStart, employee.PhoneNumber, employee.Street, employee.ZipCode)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (er *EmployeePostgres) UpdateEmployee(idEmployee string, employee entities.EmployeeInput) error {
	_, err := er.db.Exec(updateEmployee, idEmployee, employee.SurName, employee.FirstName, employee.Patronymic,
		employee.Role, employee.Salary, employee.DateOfBirth, employee.DateOfStart, employee.PhoneNumber, employee.City,
		employee.Street, employee.ZipCode)
	return err
}

func (er *EmployeePostgres) DeleteEmployee(id string) error {
	_, err := er.db.Exec(deleteEmployee, id)
	return err
}

func (er *EmployeePostgres) GetEmployeeByName(name string) (entities.Employee, error) {
	var employee entities.Employee
	if err := er.db.Get(&employee, getEmployeeByName, name); err != nil {
		return entities.Employee{}, err
	}
	return employee, nil
}

func (er *EmployeePostgres) GetAllEmployees() ([]entities.Employee, error) {
	var employees []entities.Employee
	if err := er.db.Select(&employees, getAllEmployees); err != nil {
		return []entities.Employee{}, err
	}
	return employees, nil
}
