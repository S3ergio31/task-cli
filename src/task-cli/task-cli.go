package main

import (
	"fmt"

	"github.com/S3ergio31/task-cli/commands"
	"github.com/S3ergio31/task-cli/todos"
	"github.com/S3ergio31/task-cli/todosRepository"
)

func main() {
	todosRepository.LoadTodos(&todos.TodoList)
	command := commands.Command()
	err := commands.ValidateCommand(command)

	if err != nil {
		fmt.Printf("%s", err.Error())

		return
	}

	switch command {
	case commands.ADD:
		fmt.Printf("%s\n", "# Action: Adding a new task")

		task, err := todos.Add(commands.Arguments())

		if err != nil {
			fmt.Printf("%s", err.Error())

			return
		}

		fmt.Printf("# Output: Task added successfully (ID: %s)", task.Id)
	case commands.UPDATE:
		fmt.Printf("# Action: # Updating task\n")

		task, err := todos.Update(commands.Arguments())

		if err != nil {
			fmt.Printf("%s", err.Error())

			return
		}

		fmt.Printf("# Output: Task updated successfully (ID: %s)\n", task.Id)
	case commands.REMOVE:
		fmt.Printf("# Action: # Deleting task\n")

		task, err := todos.Remove(commands.Arguments())

		if err != nil {
			fmt.Printf("%s", err.Error())

			return
		}

		fmt.Printf("# Output: Task deleted successfully (ID: %s)\n", task.Id)
	case commands.MARK:
		fmt.Printf("# Marking a task as in progress, todo or done\n")

		task, err := todos.Mark(commands.Arguments())

		if err != nil {
			fmt.Printf("%s", err.Error())

			return
		}

		fmt.Printf("# Output: Task (ID: %s) was marked as '%s'\n", task.Id, task.Status)
	case commands.LIST:
		fmt.Printf("# Action: # Listing task\n")

		filteredTodos := todos.List(commands.Arguments())

		if len(filteredTodos) == 0 {
			fmt.Printf("# Output: # There are not tasks to list\n")
		}

		for _, task := range filteredTodos {
			fmt.Printf("- %s -> %s -> %s\n", task.Id, task.Description, task.Status)
		}
	}
	todosRepository.SaveTodos(&todos.TodoList)
}
