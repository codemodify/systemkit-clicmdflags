package clicmdflags

import (
	"os"
)

// Command -
type Command struct {
	Name        string
	Description string
	Examples    []string
	Flags       interface{}
	Handler     func(command *Command)
	Hidden      bool

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
	rootCommand.subCommands = append([]*Command{helpCmd}, thisRef.subCommands...)

	if err := rootCommand.flagNeededCommandsForExecuteAndPopulateTheirFlags(os.Args[1:]); err != nil &&
		!helpCmd.flagedForExecute {

		return err
	}

	// 2. find root and ask to parse flags
	commandToExecute := thisRef.getLastSubcommandFlagedForExecute()
	if helpCmd.flagedForExecute && execute { // if `help` was asked
		commandToExecute.showUsage()
	} else if commandToExecute.Handler != nil && execute {
		commandToExecute.Handler(commandToExecute)
	}

	return nil
}
