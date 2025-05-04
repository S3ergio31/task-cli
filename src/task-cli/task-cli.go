package main

import (
	"errors"
	"fmt"
	"os"
	"slices"
)

const ADD = "add"
const UPDATE = "update"
const DELETE = "delete"
const MARK = "mark"
const LIST = "list"

var commands = []string{ADD, UPDATE, DELETE, MARK, LIST}

type task struct {
	id          int
	description string
	status      string
	createdAt   string
	updatedAt   string
}

func main() {
	command := command()
	err := validateCommand(command)

	if err != nil {
		fmt.Printf(err.Error())
	}
}

func command() string {
	return os.Args[1]
}

func validateCommand(command string) error {
	if !slices.Contains(commands, command) {
		message := fmt.Sprintf("validateCommand: Invalid command -> %s", command)
		return errors.New(message)
	}

	return nil
}
