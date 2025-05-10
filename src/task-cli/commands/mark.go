package commands

import (
	"fmt"

	"github.com/S3ergio31/task-cli/todos"
)

type markCommand struct {
	baseCommand
}

func (m markCommand) Handle() {
	fmt.Printf("# Marking a task as in progress, todo or done\n")

	task, err := todos.Mark(m.arguments())

	if err != nil {
		fmt.Printf("%s", err.Error())

		return
	}

	fmt.Printf("# Output: Task (ID: %s) was marked as '%s'\n", task.Id, task.Status)
}
