package main

import (
	"fmt"
	"log"
	"time"

	"github.com/ljb6/todolist/internal/database"
	"github.com/ljb6/todolist/internal/models"
)

func main() {
	myDb, err := database.OpenDatabase()
	if err != nil {
        log.Fatal(err)
    }
	database.CreateTable(myDb)

	task1 := models.Task{
		Text: "My second task",
		Done: false,
		Time: time.Now(),
	}

	id, err := database.AddTask(myDb, task1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(id)

    defer myDb.Close()
}