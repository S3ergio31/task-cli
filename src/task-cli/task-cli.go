package main

import (
	"errors"
	"fmt"
	"slices"
	"time"

	"github.com/S3ergio31/task-cli/commands"
	"github.com/S3ergio31/task-cli/todos"
	"github.com/S3ergio31/task-cli/todosRepository"
	"github.com/google/uuid"
)

var todoList = make(map[string]todos.Task)

func main() {
	todosRepository.LoadTodos(&todoList)
	command := commands.Command()
	err := commands.ValidateCommand(command)

	if err != nil {
		fmt.Printf("%s", err.Error())

		return
	}

	switch command {
	case commands.ADD:
		fmt.Printf("%s\n", "# Action: Adding a new task")

		task, err := add(commands.Arguments())

		if err != nil {
			fmt.Printf("%s", err.Error())

			return
		}

		fmt.Printf("# Output: Task added successfully (ID: %s)", task.Id)
	case commands.UPDATE:
		fmt.Printf("# Action: # Updating task\n")

		task, err := update(commands.Arguments())

		if err != nil {
			fmt.Printf("%s", err.Error())

			return
		}

		fmt.Printf("# Output: Task updated successfully (ID: %s)\n", task.Id)
	case commands.REMOVE:
		fmt.Printf("# Action: # Deleting task\n")

		task, err := remove(commands.Arguments())

		if err != nil {
			fmt.Printf("%s", err.Error())

			return
		}

		fmt.Printf("# Output: Task deleted successfully (ID: %s)\n", task.Id)
	case commands.MARK:
		fmt.Printf("# Marking a task as in progress, todo or done\n")

		task, err := mark(commands.Arguments())

		if err != nil {
			fmt.Printf("%s", err.Error())

			return
		}

		fmt.Printf("# Output: Task (ID: %s) was marked as '%s'\n", task.Id, task.Status)
	case commands.LIST:
		fmt.Printf("# Action: # Listing task\n")

		filteredTodos := list(commands.Arguments())

		if len(filteredTodos) == 0 {
			fmt.Printf("# Output: # There are not tasks to list\n")
		}

		for _, task := range filteredTodos {
			fmt.Printf("- %s -> %s -> %s\n", task.Id, task.Description, task.Status)
		}
	}
	todosRepository.SaveTodos(&todoList)
}

/*
 * Domain
 */
func add(arguments []string) (todos.Task, error) {
	if len(arguments) == 0 {
		return todos.Task{}, errors.New("# Error: add -> invalid description")
	}

	description := arguments[0]

	task := todos.Task{
		Id:          uuid.NewString(),
		Description: description,
		Status:      todos.TODO,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	todoList[task.Id] = task

	return task, nil
}

func update(arguments []string) (todos.Task, error) {
	if len(arguments) != 2 {
		return todos.Task{}, errors.New("# Error: update -> missing arguments")
	}

	taskId := arguments[0]
	newDescription := arguments[1]

	task, err := find(taskId)

	if err != nil {
		return todos.Task{}, err
	}

	task.Description = newDescription

	todoList[taskId] = task

	return task, nil
}

func remove(arguments []string) (todos.Task, error) {
	if len(arguments) == 0 {
		return todos.Task{}, errors.New("# Error: remove -> Task id is not defined")
	}

	taskId := arguments[0]

	task, err := find(taskId)

	if err != nil {
		return todos.Task{}, err
	}

	delete(todoList, task.Id)

	return task, nil
}

func find(taskId string) (todos.Task, error) {
	task, ok := todoList[taskId]

	if !ok {
		message := fmt.Sprintf("# Error: remove -> task with (ID: %s) not found\n", taskId)

		return todos.Task{}, errors.New(message)
	}

	return task, nil
}

func mark(arguments []string) (todos.Task, error) {
	var newStatus todos.TaskStatus

	if len(arguments) != 2 {
		return todos.Task{}, errors.New("# Error: mark -> missing arguments")
	}

	taskId := arguments[0]
	newStatus = todos.TaskStatus(arguments[1])

	if !existsStatus(newStatus) {
		message := fmt.Sprintf("# Error: mark -> Invalid status '%s'\n", newStatus)

		return todos.Task{}, errors.New(message)
	}

	task, err := find(taskId)

	if err != nil {
		return todos.Task{}, err
	}

	task.Status = todos.TaskStatus(newStatus)

	todoList[taskId] = task

	return task, nil
}

func list(arguments []string) []todos.Task {
	var status todos.TaskStatus
	var filtered []todos.Task

	if len(arguments) != 0 {
		status = todos.TaskStatus(arguments[0])
	}

	if status != "" && !existsStatus(status) {
		fmt.Printf("# Error: list -> Invalid status '%s'\n", status)
		return filtered
	}

	for _, task := range todoList {
		if status == "" || status == task.Status {
			filtered = append(filtered, task)
		}
	}

	return filtered
}

func existsStatus(status todos.TaskStatus) bool {
	statuses := []todos.TaskStatus{todos.TODO, todos.IN_PROGRESS, todos.DONE}
	return slices.Contains(statuses, status)
}
