package todos

import "time"

type TaskStatus string

type Task struct {
	Id          string
	Description string
	Status      TaskStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
