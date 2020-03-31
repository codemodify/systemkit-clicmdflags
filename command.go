package clicmdflags

import (
	"os"
)

// Command -
type Command struct {
	Name        string
	Description string
	Flags       interface{}
	Examples    []string
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
func (thisRef *Command) Execute() {
	// 1. start at root & populate all flags values for all commands
	rootCmd := thisRef
	for {
		if rootCmd.parentCommand == nil {
			break
		}
		rootCmd = rootCmd.parentCommand
	}

	thisRef.flagedForExecute = true
	thisRef.subCommands = append([]*Command{helpCmd}, thisRef.subCommands...)

	thisRef.flagCmdForExecuteAndPopulateTheirFlags(os.Args[1:])

	// 2. find root and ask to parse flags
	commandToExecute := thisRef.getLastSubcommandFlagedForExecute()
	if helpCmd.flagedForExecute { // if `help` was asked
		commandToExecute.showUsage()
	} else if commandToExecute.Handler != nil {
		commandToExecute.Handler(commandToExecute)
	}
}
