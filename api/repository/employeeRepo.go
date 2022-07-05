package repository

import (
	"ais/entities"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	createEmployee = "INSERT INTO " + employeeTable + " (id_employee, empl_surname, empl_name, empl_patronymic, " +
		"empl_role, salary, date_of_birth, date_of_start, phone_number, city, street, zip_code) " +
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12);"
	createEmployeeWithoutPatronymic = "INSERT INTO " + employeeTable + " (id_employee, empl_surname, empl_name, " +
		"empl_role, salary, date_of_birth, date_of_start, phone_number, city, street, zip_code) " +
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);"
	updateEmployee = "UPDATE " + employeeTable + " SET empl_surname=$2, empl_name=$3, empl_patronymic=$4, " +
		"empl_role=$5, salary=$6, date_of_birth=$7, date_of_start=$8, phone_number=$9, city=$10, street=$11, zip_code=$12 " +
		"WHERE id_employee=$1;"
	deleteEmployee    = "DELETE FROM " + employeeTable + " WHERE id_employee = $1 ;"
	getEmployeeByName = "SELECT id_employee, empl_surname, empl_name, empl_patronymic, empl_role, salary, date_of_birth, " +
		"date_of_start, phone_number, city, street, zip_code FROM " + employeeTable + " WHERE empl_name = $1 ;"
	getEmployeeById = "SELECT id_employee, empl_surname, empl_name, empl_patronymic, empl_role, salary, date_of_birth, " +
		"date_of_start, phone_number, city, street, zip_code FROM " + employeeTable + " WHERE id_employee=$1 ;"
	getAllEmployees = "SELECT id_employee, empl_surname, empl_name, empl_patronymic, empl_role, salary, date_of_birth, " +
		"date_of_start, phone_number, city, street, zip_code FROM " + employeeTable + " ;"
	getAllEmployeesByCategory = "SELECT id_employee, empl_surname, empl_name, empl_patronymic, empl_role, salary, date_of_birth, " +
		"date_of_start, phone_number, city, street, zip_code FROM " + employeeTable + " WHERE empl_role = $1;"
)

type employeePostgres struct {
	db *sqlx.DB
}

func NewEmployeePostgres(db *sqlx.DB) *employeePostgres {
	return &employeePostgres{db: db}
}

func (er *employeePostgres) CreateEmployee(employee entities.Employee) (int, error) {
	var id int
	row := er.db.QueryRow(createEmployee, employee.ID, employee.SurName, employee.FirstName, employee.Patronymic.String,
		employee.Role, employee.Salary, employee.DateOfBirth, employee.DateOfStart, employee.PhoneNumber, employee.City,
		employee.Street, employee.ZipCode)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (er *employeePostgres) UpdateEmployee(idEmployee string, employee entities.EmployeeInput) error {
	_, err := er.db.Exec(updateEmployee, idEmployee, employee.SurName, employee.FirstName, employee.Patronymic.String,
		employee.Role, employee.Salary, employee.DateOfBirth, employee.DateOfStart, employee.PhoneNumber, employee.City,
		employee.Street, employee.ZipCode)
	return err
}

func (er *employeePostgres) DeleteEmployee(id string) error {
	_, err := er.db.Exec(deleteEmployee, id)
	return err
}

func (er *employeePostgres) GetEmployeeByName(name string) (entities.Employee, error) {
	var employee entities.Employee
	//if err := er.db.Get(&employee, getEmployeeByName, name); err != nil {
	//	return entities.Employee{}, err
	//}
	//return employee, nil
	row := er.db.QueryRow(getEmployeeByName, name)
	err := row.Scan(&employee.ID, &employee.SurName, &employee.FirstName, &employee.Patronymic, &employee.Role,
		&employee.Salary, &employee.DateOfBirth, &employee.DateOfStart, &employee.PhoneNumber, &employee.City,
		&employee.Street, &employee.ZipCode)
	if err != nil {
		return employee, err
	}

	//return employee, nil
	//if err := er.db.Get(&employee, getEmployeeByName, name); err != nil {
	//	return entities.Employee{}, err
	//}
	return employee, nil
}

func (er *employeePostgres) GetEmployeeById(id string) (entities.Employee, error) {
	var employee entities.Employee
	row := er.db.QueryRow(getEmployeeById, id)
	err := row.Scan(&employee.ID, &employee.SurName, &employee.FirstName, &employee.Patronymic, &employee.Role,
		&employee.Salary, &employee.DateOfBirth, &employee.DateOfStart, &employee.PhoneNumber, &employee.City,
		&employee.Street, &employee.ZipCode)
	if err != nil {
		return employee, err
	}
	return employee, nil
	//if err := er.db.Get(&employee, getEmployeeById, id); err != nil {
	//	return entities.Employee{}, err
	//}
	//return employee, nil
}

func (er *employeePostgres) GetAllEmployees() ([]entities.Employee, error) {
	var employees []entities.Employee
	rows, err := er.db.Query(getAllEmployees)
	if err != nil {
		return nil, fmt.Errorf("unable to execute the query")
	}
	defer rows.Close()
	for rows.Next() {
		employee := entities.Employee{}
		err := rows.Scan(&employee.ID, &employee.SurName, &employee.FirstName, &employee.Patronymic, &employee.Role,
			&employee.Salary, &employee.DateOfBirth, &employee.DateOfStart, &employee.PhoneNumber, &employee.City,
			&employee.Street, &employee.ZipCode)
		if err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}
	return employees, nil
}

func (er *employeePostgres) GetEmployeeByRole(role string) ([]entities.Employee, error) {
	var employees []entities.Employee
	rows, err := er.db.Query(getAllEmployeesByCategory, role)
	if err != nil {
		return nil, fmt.Errorf("unable to execute the query")
	}
	defer rows.Close()
	for rows.Next() {
		employee := entities.Employee{}
		err := rows.Scan(&employee.ID, &employee.SurName, &employee.FirstName, &employee.Patronymic, &employee.Role,
			&employee.Salary, &employee.DateOfBirth, &employee.DateOfStart, &employee.PhoneNumber, &employee.City,
			&employee.Street, &employee.ZipCode)
		if err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}
	return employees, nil
}
