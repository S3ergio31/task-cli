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

type Command interface {
	Handle()
	arguments() []string
}

type baseCommand struct{}

func (bc baseCommand) arguments() []string {
	return os.Args[2:]
}

func validateCommand(command string) error {
	if !slices.Contains(commands, command) {
		message := fmt.Sprintf("validateCommand: Invalid command -> %s", command)

		return errors.New(message)
	}

	return nil
}

func Build() (Command, error) {
	command := os.Args[1]
	err := validateCommand(command)

	if err != nil {
		return nil, err
	}

	switch command {
	case ADD:
		return addCommand{}, nil
	case UPDATE:
		return updateCommand{}, nil
	case REMOVE:
		return removeCommand{}, nil
	case MARK:
		return markCommand{}, nil
	case LIST:
		return listCommand{}, nil
	}

	return nil, nil
}
