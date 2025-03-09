package models

import "time"

type Task struct {
	Id   int       `json:"id"`
	Text string    `json:"text"`
	Done bool      `json:"done"`
	Time time.Time `json:"time"`
}
