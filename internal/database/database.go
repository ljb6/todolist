package database

import (
	"database/sql"
	"encoding/json"
	"log"

	"github.com/ljb6/todolist/internal/models"
	_ "github.com/mattn/go-sqlite3"
)

func OpenDatabase() (*sql.DB, error) {

	db, err := sql.Open("sqlite3", "./tasks.db")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func CreateTable(db *sql.DB) {

	query := `
	CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		text TEXT NOT NULL,
		done BOOLEAN NO NULL,
		time DATETIME DEFAULT CURRENT_TIMESTAMP
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

// ([]models.Task, error)
func GetTasks(db *sql.DB) ([]byte, error) {

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

	jsonTasks, err := json.Marshal(tasks)
	if err != nil {
		return nil, err
	}
	return jsonTasks, nil
}
