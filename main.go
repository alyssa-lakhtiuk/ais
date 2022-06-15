package ais

import (
	"ais/config"
	"fmt"
	"log"
)

const (
	fileName = "config.json"
)

func main() {
	jsonConfig, err := config.New(fileName)
	fmt.Println(jsonConfig)
	if err != nil {
		log.Fatal(err)
	}
}
