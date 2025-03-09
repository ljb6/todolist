package database

import (
	"database/sql"
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
		time DATETIME DEFAUL CURRENT_TIMESTAMP
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
