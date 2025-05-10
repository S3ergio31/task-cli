package commands

import (
	"fmt"

	"github.com/S3ergio31/task-cli/todos"
)

type updateCommand struct {
	baseCommand
}

func (u updateCommand) Handle() {
	fmt.Printf("# Action: Updating task\n")

	task, err := todos.Update(u.arguments())

	if err != nil {
		fmt.Printf("%s", err.Error())

		return
	}

	fmt.Printf("# Output: Task updated successfully (ID: %s)\n", task.Id)
}
