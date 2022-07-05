package repository

import (
	"ais/entities"
	"github.com/jmoiron/sqlx"
)

const (
	createRole = "INSERT INTO " + rolesTable + " (empl_id, role, password, phone) " +
		"VALUES ($1, $2, $3, $4);"
	deleteRole          = "DELETE FROM " + rolesTable + " WHERE  phone = $1;"
	getRoleByPhone      = "SELECT * FROM " + rolesTable + " WHERE phone = $1;"
	getRoleByIdEmployee = "SELECT * FROM " + rolesTable + " WHERE empl_id = $1;"
)

type rolesPostgres struct {
	db *sqlx.DB
}

func NewRolesPostgres(db *sqlx.DB) *rolesPostgres {
	return &rolesPostgres{db: db}
}

func (er *rolesPostgres) GetRoleByPhone(phone string) (entities.SignIn, error) {
	var role entities.SignIn
	row := er.db.QueryRow(getRoleByPhone, phone)
	err := row.Scan(&role.IdEmployee, &role.Role, &role.Password, &role.Phone)
	if err != nil {
		return role, err
	}
	return role, nil
}

func (er *rolesPostgres) GetRoleByIdEmployee(id string) (entities.SignIn, error) {
	var role entities.SignIn
	row := er.db.QueryRow(getRoleByIdEmployee, id)
	err := row.Scan(&role.IdEmployee, &role.Role, &role.Password, &role.Phone)
	if err != nil {
		return role, err
	}
	return role, nil
}
