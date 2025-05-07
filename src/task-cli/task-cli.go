package main

import (
	"fmt"

	"github.com/S3ergio31/task-cli/commands"
	"github.com/S3ergio31/task-cli/todos"
	"github.com/S3ergio31/task-cli/todosRepository"
)

func main() {
	todosRepository.LoadTodos(&todos.TodoList)

	command, err := command()

	if err != nil {
		fmt.Printf("%s", err.Error())

		return
	}

	handleCommand(command)

	todosRepository.SaveTodos(&todos.TodoList)
}

func command() (string, error) {
	command := commands.Command()
	err := commands.ValidateCommand(command)

	return command, err
}

func handleCommand(command string) {
	switch command {
	case commands.ADD:
		add()
	case commands.UPDATE:
		update()
	case commands.REMOVE:
		remove()
	case commands.MARK:
		mark()
	case commands.LIST:
		list()
	}
}

func add() {
	fmt.Printf("%s\n", "# Action: Adding a new task")

	task, err := todos.Add(commands.Arguments())

	if err != nil {
		fmt.Printf("%s", err.Error())

		return
	}

	fmt.Printf("# Output: Task added successfully (ID: %s)", task.Id)
}

func update() {
	fmt.Printf("# Action: Updating task\n")

	task, err := todos.Update(commands.Arguments())

	if err != nil {
		fmt.Printf("%s", err.Error())

		return
	}

	fmt.Printf("# Output: Task updated successfully (ID: %s)\n", task.Id)
}

func remove() {
	fmt.Printf("# Action: Deleting task\n")

	task, err := todos.Remove(commands.Arguments())

	if err != nil {
		fmt.Printf("%s", err.Error())

		return
	}

	fmt.Printf("# Output: Task deleted successfully (ID: %s)\n", task.Id)
}

func mark() {
	fmt.Printf("# Marking a task as in progress, todo or done\n")

	task, err := todos.Mark(commands.Arguments())

	if err != nil {
		fmt.Printf("%s", err.Error())

		return
	}

	fmt.Printf("# Output: Task (ID: %s) was marked as '%s'\n", task.Id, task.Status)
}

func list() {
	fmt.Printf("# Action: Listing task\n")

	filteredTodos := todos.List(commands.Arguments())

	if len(filteredTodos) == 0 {
		fmt.Printf("# Output: There are not tasks to list\n")
	}

	for _, task := range filteredTodos {
		fmt.Printf("- %s -> %s -> %s\n", task.Id, task.Description, task.Status)
	}
}
