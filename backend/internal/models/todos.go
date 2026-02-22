package models

type ToDo struct {
	ID          string `json:"todo_id"`
	Name        string `json:"name"`
	IsCompleted bool   `json:"is_completed"`
	CreatedAt   string `json:"created_at"`
}
