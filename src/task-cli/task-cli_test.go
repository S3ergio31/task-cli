package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/S3ergio31/task-cli/todos"
	"github.com/S3ergio31/task-cli/todosRepository"
)

func clearTodosJson() {
	todoList := make(map[string]todos.Task)
	todosRepository.SaveTodos(&todoList)
	todos.TodoList = todoList
}

func assertOutput(t *testing.T, arguments []string, expectedOutput string) {
	originalArgs := os.Args
	originalStdout := os.Stdout
	os.Args = arguments

	var buf bytes.Buffer
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	buf.ReadFrom(r)
	os.Stdout = originalStdout
	os.Args = originalArgs

	output := buf.String()

	if !strings.Contains(output, expectedOutput) {
		setTestError(t, output, expectedOutput)
	}

	t.Cleanup(clearTodosJson)
}

func setTestError(t *testing.T, output string, expectedOutput string) {
	t.Errorf("Expected %q but got %q", expectedOutput, output)
}

func TestUndefinedCommand(t *testing.T) {
	arguments := []string{"task-cli", "test"}
	expectedOutput := "validateCommand: Invalid command -> test"
	assertOutput(t, arguments, expectedOutput)
}

func TestAddCommandOk(t *testing.T) {
	arguments := []string{"task-cli", "add", "Check emails"}
	expectedOutput := "# Action: Adding a new task\n# Output: Task added successfully"
	assertOutput(t, arguments, expectedOutput)
}

func TestAddCommandWithoutDescription(t *testing.T) {
	arguments := []string{"task-cli", "add"}
	expectedOutput := "# Action: Adding a new task\n# Error: add -> invalid description"
	assertOutput(t, arguments, expectedOutput)
}

func TestUpdateCommandOk(t *testing.T) {
	task, _ := todos.Add([]string{"Test task"})
	newDescription := "Updated task"
	arguments := []string{"task-cli", "update", task.Id, newDescription}
	expectedOutput := "# Action: Updating task\n# Output: Task updated successfully"
	assertOutput(t, arguments, expectedOutput)
	task, _ = todos.Find(task.Id)
	if task.Description != newDescription {
		setTestError(t, task.Description, newDescription)
	}
}

func TestUpdateCommandMissingArguments(t *testing.T) {
	arguments := []string{"task-cli", "update"}
	expectedOutput := "# Action: Updating task\n# Error: update -> missing arguments"
	assertOutput(t, arguments, expectedOutput)
}

func TestUpdateCommandTaskNotFound(t *testing.T) {
	arguments := []string{"task-cli", "update", "a-b-c", "Test task"}
	expectedOutput := "# Action: Updating task\n# Error: remove -> task with (ID: a-b-c) not found\n"
	assertOutput(t, arguments, expectedOutput)
}

func TestRemoveCommandOk(t *testing.T) {
	task, _ := todos.Add([]string{"Test task"})
	arguments := []string{"task-cli", "remove", task.Id}
	expectedOutput := fmt.Sprintf("# Action: Deleting task\n# Output: Task deleted successfully (ID: %s)\n", task.Id)
	assertOutput(t, arguments, expectedOutput)
}

func TestRemoveCommandMissingArgumentTaskId(t *testing.T) {
	arguments := []string{"task-cli", "remove"}
	expectedOutput := "# Action: Deleting task\n# Error: remove -> Task id is not defined"
	assertOutput(t, arguments, expectedOutput)
}

func TestRemoveCommandTaskNotFound(t *testing.T) {
	arguments := []string{"task-cli", "remove", "a-b-c"}
	expectedOutput := "# Action: Deleting task\n# Error: remove -> task with (ID: a-b-c) not found\n"
	assertOutput(t, arguments, expectedOutput)
}

func TestMarkCommandOk(t *testing.T) {
	task, _ := todos.Add([]string{"Test task"})
	newStatus := string(todos.IN_PROGRESS)
	arguments := []string{"task-cli", "mark", task.Id, newStatus}
	expectedOutput := fmt.Sprintf("# Marking a task as in progress, todo or done\n# Output: Task (ID: %s) was marked as '%s'\n", task.Id, newStatus)
	assertOutput(t, arguments, expectedOutput)
	task, _ = todos.Find(task.Id)
	if task.Status != todos.TaskStatus(newStatus) {
		setTestError(t, string(task.Status), newStatus)
	}
}

func TestMarkCommandMissingArguments(t *testing.T) {
	arguments := []string{"task-cli", "mark"}
	expectedOutput := "# Marking a task as in progress, todo or done\n# Error: mark -> missing arguments"
	assertOutput(t, arguments, expectedOutput)
}

func TestMarkCommandTaskNotFound(t *testing.T) {
	arguments := []string{"task-cli", "mark", "a-b-c", string(todos.DONE)}
	expectedOutput := "# Marking a task as in progress, todo or done\n# Error: remove -> task with (ID: a-b-c) not found\n"
	assertOutput(t, arguments, expectedOutput)
}

func TestMarkCommandInvalidStatus(t *testing.T) {
	arguments := []string{"task-cli", "mark", "a-b-c", "a-status"}
	expectedOutput := "# Marking a task as in progress, todo or done\n# Error: mark -> Invalid status 'a-status'\n"
	assertOutput(t, arguments, expectedOutput)
}

func TestListAllCommandOk(t *testing.T) {
	task, _ := todos.Add([]string{"Test task"})
	arguments := []string{"task-cli", "list"}
	expectedOutput := fmt.Sprintf("# Action: Listing task\n- %s -> %s -> %s\n", task.Id, task.Description, task.Status)
	assertOutput(t, arguments, expectedOutput)
}

func TestListEmptyCommand(t *testing.T) {
	arguments := []string{"task-cli", "list"}
	expectedOutput := "# Action: Listing task\n# Output: There are not tasks to list\n"
	assertOutput(t, arguments, expectedOutput)
}
