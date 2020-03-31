package clicmdflags

import (
	"fmt"
	"reflect"
	"strings"
)

var helpCmd = &Command{
	Name: "help",
}

func (thisRef *Command) showUsage() {
	definedFlags := thisRef.getDefinedFlags()
	areTheseGlobalFlags := (thisRef.parentCommand == nil)

	flagNames := []string{}
	for _, flag := range definedFlags {
		flagNames = append(flagNames, flag.Name)
	}

	flags := strings.Join(flagNames, " VALUE ")
	if len(flags) > 0 {
		flags = flags + " VALUE"
	}
	usageString := fmt.Sprintf(" %s %s", thisRef.Name, flags)
	cmd := thisRef.parentCommand
	for {
		if cmd == nil {
			break
		}

		usageString = fmt.Sprintf(" %s", cmd.Name) + usageString

		cmd = cmd.parentCommand
	}

	fmt.Println()

	fmt.Println(fmt.Sprintf("%s, %s", strings.ToUpper(thisRef.Name), thisRef.Description))

	fmt.Println("~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~")
	fmt.Println(fmt.Sprintf("USAGE   | %s", strings.TrimSpace(usageString)))

	if len(definedFlags) > 0 {
		fmt.Println("~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~")
		fmt.Print(fmt.Sprintf("FLAGS   |"))
		for i, flag := range definedFlags {
			if i == 0 {
				fmt.Println(fmt.Sprintf(" %s, type=%s, default=%s, %s", flag.Name, flag.Type, flag.Default, flag.Description))
			} else {
				fmt.Println(fmt.Sprintf("        | %s, type=%s, default=%s, %s", flag.Name, flag.Type, flag.Default, flag.Description))
			}
		}
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
			fmt.Println("        | ~~~~  global flags  ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~")
			for _, flag := range definedFlags {
				fmt.Println(fmt.Sprintf("        | %s, type=%s, default=%s, %s", flag.Name, flag.Type, flag.Default, flag.Description))
			}
		}
	}

	if len(thisRef.Examples) > 0 {
		fmt.Println("~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~")
		fmt.Println(fmt.Sprintf("EXAMPLE | %s", strings.Join(thisRef.Examples, "\n")))
	}

	if len(thisRef.subCommands) > 0 {
		fmt.Println("~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~")
		fmt.Print(fmt.Sprintf("SUBCMD  |"))
		firstOnePrinted := false
		for _, c := range thisRef.subCommands {
			if c != helpCmd {
				if !firstOnePrinted {
					fmt.Println(fmt.Sprintf(" %s, %s", c.Name, c.Description))
					firstOnePrinted = true
				} else {
					fmt.Println(fmt.Sprintf("        | %s, %s", c.Name, c.Description))
				}
			}
		}
	}

	fmt.Println()
}

func (thisRef *Command) getDefinedFlags() []Flag {
	result := []Flag{}

	runtimeStructRef, err := thisRef.getRealRuntimeStruct(reflect.ValueOf(thisRef.Flags))
	if err != nil {
		return result
	}

	for i := 0; i < runtimeStructRef.NumField(); i++ {
		result = append(result, Flag{
			Name:        flagPattern + runtimeStructRef.Type().Field(i).Tag.Get("flagName"),
			Description: runtimeStructRef.Type().Field(i).Tag.Get("flagDescription"),
			Type:        runtimeStructRef.Type().Field(i).Type.Name(),
			Default:     runtimeStructRef.Type().Field(i).Tag.Get("flagDefault"),
		})
	}

	return result
}
