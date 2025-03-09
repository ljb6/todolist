package main

import (
	"fmt"
	"log"

	"github.com/ljb6/todolist/internal/database"
)

func main() {
	myDb, err := database.OpenDatabase()
	if err != nil {
        log.Fatal(err)
    }
	database.CreateTable(myDb)

	// task1 := models.Task{
	// 	Text: "Third task",
	// 	Done: false,
	// 	Time: time.Now(),
	// }

	// database.AddTask(myDb, task1)

	x, err := database.GetTasks(myDb)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(x), err)

    defer myDb.Close()
}