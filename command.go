package clicmdflags

import (
	"os"
	"strings"
)

// Command -
type Command struct {
	Name        string
	Description string
	Examples    []string
	Flags       interface{}
	Handler     func(command *Command)
	Hidden      bool
	PassThrough bool

	parentCommand    *Command
	subCommands      []*Command
	flagedForExecute bool
}

// AddCommand -
func (thisRef *Command) AddCommand(command *Command) {
	command.parentCommand = thisRef

	thisRef.subCommands = append(thisRef.subCommands, command)
}

// Execute -
func (thisRef *Command) Execute() error {
	return thisRef.parseAndExecute(true)
}

// ParseFlags -
func (thisRef *Command) ParseFlags() error {
	return thisRef.parseAndExecute(false)
}

// Execute -
func (thisRef *Command) parseAndExecute(execute bool) error {
	// 1. start at root & populate all flags values for all commands
	rootCommand := thisRef.getRootCommand()
	rootCommand.flagedForExecute = true

	argsToParse := os.Args[1:]
	helpsIsWanted := strings.HasSuffix(strings.Join(argsToParse, " "), "help")

	if err := rootCommand.flagNeededCommandsForExecuteAndPopulateTheirFlags(argsToParse); err != nil &&
		!helpsIsWanted {
		return err
	}

	// 2. find root and ask to parse flags
	commandToExecute := thisRef.getLastSubcommandFlagedForExecute()
	if execute && helpsIsWanted {
		commandToExecute.showUsage()
	} else if execute && commandToExecute.Handler != nil {
		commandToExecute.Handler(commandToExecute)
	}

	return nil
}
