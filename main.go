package main

import (
	"log"

	"github.com/ljb6/todolist/internal/database"
)

func main() {
	myDb, err := database.OpenDatabase()
	if err != nil {
        log.Fatal(err)
    }
	database.CreateTable(myDb)
    defer myDb.Close()
}