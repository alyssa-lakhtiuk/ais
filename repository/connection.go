package repository

import (
	"ais/config"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

func NewDB(jsonConfig config.Postgresql) (*sqlx.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		jsonConfig.Host, jsonConfig.Port, jsonConfig.User, jsonConfig.Password, jsonConfig.DBName)
	db, err := sqlx.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}
