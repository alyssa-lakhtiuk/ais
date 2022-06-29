package main

import (
	"ais/config"
	"ais/repository"
	"ais/web/controller"
	"ais/web/service"
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
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handler := controller.NewHandler(services)

	router := handler.NewRoutes()
	router.Run()
	//if err = http.ListenAndServe(jsonConfig.ListenUrl, handler); err != nil {
	//	log.Fatal(err.Error())
	//}
}
