package models

type Task struct {
	Id   int       `json:"id"`
	Text string    `json:"text"`
	Done bool      `json:"done"`
	Time string `json:"time"`
}
