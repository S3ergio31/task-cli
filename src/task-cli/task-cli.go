package main

import (
	"fmt"

	"github.com/S3ergio31/task-cli/commands"
	"github.com/S3ergio31/task-cli/todos"
	"github.com/S3ergio31/task-cli/todosRepository"
)

func main() {
	todosRepository.LoadTodos(&todos.TodoList)

	command, err := commands.Build()

	if err != nil {
		fmt.Printf("%s", err.Error())

		return
	}

	command.Handle()

	todosRepository.SaveTodos(&todos.TodoList)
}
