package main

import (
	"github.com/ljb6/todolist/internal/database"
	"github.com/ljb6/todolist/internal/routes"
)

func main() {
	database.OpenDatabase()
	database.CreateTable(database.DB)
	routes.InitializeServer()
}