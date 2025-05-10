package commands

import (
	"fmt"

	"github.com/S3ergio31/task-cli/todos"
)

type listCommand struct {
	baseCommand
}

func (l listCommand) Handle() {
	fmt.Printf("# Action: Listing task\n")

	filteredTodos := todos.List(l.arguments())

	if len(filteredTodos) == 0 {
		fmt.Printf("# Output: There are not tasks to list\n")
	}

	for _, task := range filteredTodos {
		fmt.Printf("- %s -> %s -> %s\n", task.Id, task.Description, task.Status)
	}
}
