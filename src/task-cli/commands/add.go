package commands

import (
	"fmt"

	"github.com/S3ergio31/task-cli/todos"
)

type addCommand struct {
	baseCommand
}

func (a addCommand) Handle() {
	fmt.Printf("%s\n", "# Action: Adding a new task")

	task, err := todos.Add(a.arguments())

	if err != nil {
		fmt.Printf("%s", err.Error())

		return
	}

	fmt.Printf("# Output: Task added successfully (ID: %s)", task.Id)
}
