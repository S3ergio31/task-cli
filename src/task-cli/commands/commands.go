package commands

import (
	"errors"
	"fmt"
	"os"
	"slices"
)

const (
	ADD    = "add"
	UPDATE = "update"
	REMOVE = "remove"
	MARK   = "mark"
	LIST   = "list"
)

var commands = []string{ADD, UPDATE, REMOVE, MARK, LIST}

func Command() string {
	return os.Args[1]
}

func Arguments() []string {
	return os.Args[2:]
}

func ValidateCommand(command string) error {
	if !slices.Contains(commands, command) {
		message := fmt.Sprintf("validateCommand: Invalid command -> %s", command)

		return errors.New(message)
	}

	return nil
}
