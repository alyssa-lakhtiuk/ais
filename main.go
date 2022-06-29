package main

import (
	"ais/config"
	"ais/repository"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

const (
	fileName = "config.json"
)

func main() {
	var (
		db *sqlx.DB
	)
	jsonConfig, err := config.New(fileName)
	//fmt.Println(jsonConfig)
	if err != nil {
		log.Fatal(err)
	}
	if db, err = repository.NewDB(jsonConfig.Postgresql); err != nil {
		log.Fatal()
	}
	fmt.Println(db)
}
