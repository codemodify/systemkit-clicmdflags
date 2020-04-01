package clicmdflags

import (
	"fmt"
	"strings"
)

var helpCmd = &Command{
	Name: "help",
}

func (thisRef *Command) showUsage() {
	definedFlags := thisRef.getDefinedFlags()
	areTheseGlobalFlags := (thisRef.parentCommand == nil)

	usageString := fmt.Sprintf(" %s COMMAND(s) %sFLAG(s)", thisRef.Name, flagPatterns[0])
	cmd := thisRef.parentCommand
	for {
		if cmd == nil {
			break
		}

		usageString = fmt.Sprintf(" %s", cmd.Name) + usageString

		cmd = cmd.parentCommand
	}

	fmt.Println()
	fmt.Println("~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~")
	fmt.Println(fmt.Sprintf(" %s", thisRef.Description))

	fmt.Println("~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~")
	fmt.Println(fmt.Sprintf("    Usage | %s", strings.TrimSpace(usageString)))

	if len(definedFlags) > 0 {
		fmt.Println("~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~")
		fmt.Print(fmt.Sprintf("    Flags |"))
		for i, definedFlag := range definedFlags {
			if i == 0 {
				fmt.Println(fmt.Sprintf(" %s, type=%s, required=%s, default=%s, %s", flagPatterns[0]+definedFlag.name, definedFlag.typeName, definedFlag.isRequired, definedFlag.defaultValue, definedFlag.description))
			} else {
				fmt.Println(fmt.Sprintf("          | %s, type=%s, required=%s, default=%s, %s", flagPatterns[0]+definedFlag.name, definedFlag.typeName, definedFlag.isRequired, definedFlag.defaultValue, definedFlag.description))
			}
		}
		fmt.Println(fmt.Sprintf("          |"))
	}

	if !areTheseGlobalFlags {
		rootCmd := thisRef.parentCommand
		for {
			if rootCmd.parentCommand == nil {
				break
			}
			rootCmd = rootCmd.parentCommand
		}

		definedFlags = rootCmd.getDefinedFlags()

		if len(definedFlags) > 0 {
			fmt.Println("          | ~~~~ global flags ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~")
			for _, definedFlag := range definedFlags {
				fmt.Println(fmt.Sprintf("          | %s, type=%s, required=%s, default=%s, %s", flagPatterns[0]+definedFlag.name, definedFlag.typeName, definedFlag.isRequired, definedFlag.defaultValue, definedFlag.description))
			}
			fmt.Println(fmt.Sprintf("          |"))
		}
	}

	if len(thisRef.Examples) > 0 {
		fmt.Println("~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~")
		fmt.Print(fmt.Sprintf(" Examples |"))
		for i, example := range thisRef.Examples {
			if i == 0 {
				fmt.Println(fmt.Sprintf(" %s", example))
			} else {
				fmt.Println(fmt.Sprintf("          | %s", example))
			}
		}
		fmt.Println(fmt.Sprintf("          |"))
	}

	if len(thisRef.subCommands) > 0 {
		fmt.Println("~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~")
		fmt.Print(fmt.Sprintf(" Commands |"))
		firstOnePrinted := false
		for _, c := range thisRef.subCommands {
			if c != helpCmd {
				if !firstOnePrinted {
					fmt.Println(fmt.Sprintf(" %s, %s", c.Name, c.Description))
					firstOnePrinted = true
				} else {
					fmt.Println(fmt.Sprintf("          | %s, %s", c.Name, c.Description))
				}
			}
		}
		fmt.Println(fmt.Sprintf("          |"))
	}

	fmt.Println("~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~")
	fmt.Println()
}
