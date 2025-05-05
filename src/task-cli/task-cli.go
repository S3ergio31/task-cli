package main

import (
	"errors"
	"fmt"
	"os"
	"slices"
	"time"

	"github.com/google/uuid"
)

// Actions
const ADD = "add"
const UPDATE = "update"
const REMOVE = "remove"
const MARK = "mark"
const LIST = "list"

// Task statuses
type TaskStatus string

const (
	TODO        TaskStatus = "todo"
	IN_PROGRESS TaskStatus = "in-progress"
	DONE        TaskStatus = "done"
)

type Task struct {
	id          string
	description string
	status      TaskStatus
	createdAt   string
	updatedAt   string
}

var commands = []string{ADD, UPDATE, REMOVE, MARK, LIST}
var todos = make(map[string]*Task)

func main() {
	command := command()
	err := validateCommand(command)

	if err != nil {
		fmt.Printf("%s", err.Error())

		return
	}

	switch command {
	case ADD:
		fmt.Printf("%s\n", "# Action: Adding a new task")
		task, err := add(arguments())

		if err != nil {
			fmt.Printf("%s", err.Error())

			return
		}

		message := fmt.Sprintf("# Output: Task added successfully (ID: %s)", task.id)
		fmt.Printf("%s", message)
	case UPDATE:
	case REMOVE:
	case MARK:
	case LIST:
	}

}

func handleError(err *error) {

}

/*
 * Commands
 */
func command() string {
	return os.Args[1]
}

func arguments() []string {
	return os.Args[2:]
}

func validateCommand(command string) error {
	if !slices.Contains(commands, command) {
		message := fmt.Sprintf("validateCommand: Invalid command -> %s", command)

		return errors.New(message)
	}

	return nil
}

/*
 * Domain
 */
func add(arguments []string) (*Task, error) {
	if len(arguments) == 0 {
		return nil, errors.New("# Error: add -> invalid description")
	}

	description := arguments[0]

	task := &Task{
		id:          uuid.NewString(),
		description: description,
		status:      TODO,
		createdAt:   time.Now().GoString(),
		updatedAt:   time.Now().GoString(),
	}

	todos[task.id] = task

	return task, nil
}

func update(task Task, description string) {

}

func remove(task Task) {
	delete(todos, task.id)
}

func mark() {

}

func list() {

}
