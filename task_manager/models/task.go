package models

type Status string

const (
	MISSED    Status = "missed"
	PENDING   Status = "pending"
	COMPLETED Status = "completed"
)

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     Date   `json:"due_date"`
	Status      string `json:"status"`
}
