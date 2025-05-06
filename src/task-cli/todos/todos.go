package todos

import "time"

type TaskStatus string

const (
	TODO        TaskStatus = "todo"
	IN_PROGRESS TaskStatus = "in-progress"
	DONE        TaskStatus = "done"
)

type Task struct {
	Id          string
	Description string
	Status      TaskStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
