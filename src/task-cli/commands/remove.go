package commands

import (
	"fmt"

	"github.com/S3ergio31/task-cli/todos"
)

type removeCommand struct {
	baseCommand
}

func (r removeCommand) Handle() {
	fmt.Printf("# Action: Deleting task\n")

	task, err := todos.Remove(r.arguments())

	if err != nil {
		fmt.Printf("%s", err.Error())

		return
	}

	fmt.Printf("# Output: Task deleted successfully (ID: %s)\n", task.Id)
}
