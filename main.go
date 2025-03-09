package main

import (
	"fmt"
	"time"

	"github.com/ljb6/todolist/internal/models"
)

func main() {
	task1 := models.Task{Id:1, Text: "Compras", Done: false, Time: time.Now()}
	fmt.Println(task1)
}