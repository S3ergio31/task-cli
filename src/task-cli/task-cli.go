package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
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
	Id          string
	Description string
	Status      TaskStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

var commands = []string{ADD, UPDATE, REMOVE, MARK, LIST}
var todos = make(map[string]Task)

func main() {
	loadTodos()
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

		fmt.Printf("# Output: Task added successfully (ID: %s)", task.Id)
	case UPDATE:
		fmt.Printf("# Action: # Updating task\n")

		task, err := update(arguments())

		if err != nil {
			fmt.Printf("%s", err.Error())

			return
		}

		fmt.Printf("# Output: Task updated successfully (ID: %s)\n", task.Id)
	case REMOVE:
		fmt.Printf("# Action: # Deleting task\n")

		task, err := remove(arguments())

		if err != nil {
			fmt.Printf("%s", err.Error())

			return
		}

		fmt.Printf("# Output: Task deleted successfully (ID: %s)\n", task.Id)
	case MARK:
	case LIST:
	}
	saveTodos()
}

func handleError(err *error) {

}

func loadTodos() {
	jsonFile, err := os.Open("todos.json")

	if err != nil {
		fmt.Printf("loadTodos: %s\n", err)

		return
	}

	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &todos)
}

func saveTodos() {
	bytes, jsonErr := json.Marshal(todos)

	if jsonErr != nil {
		fmt.Printf("saveTodos: %s\n", jsonErr)

		return
	}

	writeError := os.WriteFile("todos.json", bytes, 0644)

	if writeError != nil {
		fmt.Printf("saveTodos: %s\n", writeError)

		return
	}
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
func add(arguments []string) (Task, error) {
	if len(arguments) == 0 {
		return Task{}, errors.New("# Error: add -> invalid description")
	}

	description := arguments[0]

	task := Task{
		Id:          uuid.NewString(),
		Description: description,
		Status:      TODO,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	todos[task.Id] = task

	return task, nil
}

func update(arguments []string) (Task, error) {
	if len(arguments) != 2 {
		return Task{}, errors.New("# Error: update -> missing arguments")
	}

	taskId := arguments[0]
	newDescription := arguments[1]

	task, err := find(taskId)

	if err != nil {
		return Task{}, err
	}

	task.Description = newDescription

	todos[taskId] = task

	return task, nil
}

func remove(arguments []string) (Task, error) {
	if len(arguments) == 0 {
		return Task{}, errors.New("# Error: remove -> Task id is not defined")
	}

	taskId := arguments[0]

	task, err := find(taskId)

	if err != nil {
		return Task{}, err
	}

	delete(todos, task.Id)

	return task, nil
}

func find(taskId string) (Task, error) {
	task, ok := todos[taskId]

	if !ok {
		message := fmt.Sprintf("# Error: remove -> task with (ID: %s) not found\n", taskId)

		return Task{}, errors.New(message)
	}

	return task, nil
}

func mark() {

}

func list() {

}
