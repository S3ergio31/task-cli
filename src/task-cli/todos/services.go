package todos

import (
	"errors"
	"fmt"
	"slices"
	"time"

	"github.com/google/uuid"
)

var TodoList = make(map[string]Task)

func Add(arguments []string) (Task, error) {
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

	TodoList[task.Id] = task

	return task, nil
}

func Update(arguments []string) (Task, error) {
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

	TodoList[taskId] = task

	return task, nil
}

func Remove(arguments []string) (Task, error) {
	if len(arguments) == 0 {
		return Task{}, errors.New("# Error: remove -> Task id is not defined")
	}

	taskId := arguments[0]

	task, err := find(taskId)

	if err != nil {
		return Task{}, err
	}

	delete(TodoList, task.Id)

	return task, nil
}

func find(taskId string) (Task, error) {
	task, ok := TodoList[taskId]

	if !ok {
		message := fmt.Sprintf("# Error: remove -> task with (ID: %s) not found\n", taskId)

		return Task{}, errors.New(message)
	}

	return task, nil
}

func Mark(arguments []string) (Task, error) {
	var newStatus TaskStatus

	if len(arguments) != 2 {
		return Task{}, errors.New("# Error: mark -> missing arguments")
	}

	taskId := arguments[0]
	newStatus = TaskStatus(arguments[1])

	if !existsStatus(newStatus) {
		message := fmt.Sprintf("# Error: mark -> Invalid status '%s'\n", newStatus)

		return Task{}, errors.New(message)
	}

	task, err := find(taskId)

	if err != nil {
		return Task{}, err
	}

	task.Status = TaskStatus(newStatus)

	TodoList[taskId] = task

	return task, nil
}

func List(arguments []string) []Task {
	var status TaskStatus
	var filtered []Task

	if len(arguments) != 0 {
		status = TaskStatus(arguments[0])
	}

	if status != "" && !existsStatus(status) {
		fmt.Printf("# Error: list -> Invalid status '%s'\n", status)
		return filtered
	}

	for _, task := range TodoList {
		if status == "" || status == task.Status {
			filtered = append(filtered, task)
		}
	}

	return filtered
}

func existsStatus(status TaskStatus) bool {
	statuses := []TaskStatus{TODO, IN_PROGRESS, DONE}
	return slices.Contains(statuses, status)
}
