package main

import (
	"log"

	"github.com/ZainJavedDev/catch-up-check/database"
	"github.com/ZainJavedDev/catch-up-check/storage"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	err := storage.CheckStorage()
	if err != nil {
		log.Fatal(err)
	}
	err = database.Check()
	if err != nil {
		log.Fatal(err)
	}
}
