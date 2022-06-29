package repository

import (
	"ais/config"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

func NewDB(jsonConfig config.Postgresql) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf(
		"host=%s port=%v password=%s dbname=%s user=%s sslmode=disable",
		jsonConfig.Host, jsonConfig.Port, jsonConfig.Password, jsonConfig.DBName,
		jsonConfig.User))
	//fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
	//	jsonConfig.User,
	//	jsonConfig.Password,
	//	jsonConfig.Host,
	//	jsonConfig.Port,
	//	jsonConfig.DBName)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}
