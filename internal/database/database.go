package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/ljb6/todolist/internal/models"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func OpenDatabase() error {
	var err error
	DB, err = sql.Open("sqlite3", "./tasks.db")
	if err != nil {
		return err
	}
	return nil
}

func CreateTable(db *sql.DB) {

	query := `
	CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		text TEXT NOT NULL,
		done BOOLEAN NO NULL,
		time TEXT NOT NULL
	);
	`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func AddTask(db *sql.DB, task models.Task) (int64, error) {

	query := `INSERT INTO tasks (text, done, time) VALUES (?, ?, ?)`
	result, err := db.Exec(query, task.Text, task.Done, task.Time)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	return id, err
}

func GetTasks(db *sql.DB) ([]models.Task, error) {

	var tasks []models.Task

	row, err := db.Query("SELECT * FROM tasks")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()

	for row.Next() {
		item := models.Task{}
		err := row.Scan(&item.Id, &item.Text, &item.Done, &item.Time)
		if err != nil {
			log.Fatal(err)
		}
		tasks = append(tasks, item)
	}

	return tasks, nil
}

func DeleteTask(db *sql.DB, id string) {
	query := "DELETE FROM tasks WHERE id = ?" 
	result, err := db.Exec(query, id)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result.LastInsertId())
}

func MarkTaskAsDone(db *sql.DB, id string) {
	query := "UPDATE tasks SET done = ? WHERE id = ?"
	result, err := db.Exec(query, 1, id)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}

func DeleteAllTasks(db *sql.DB) {
	query := "DELETE FROM tasks"
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}	