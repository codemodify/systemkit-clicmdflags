package clicmdflags

import (
	"fmt"
	"strconv"
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

	var constHorizontalLine = string('\u2500')
	var constHalfCrossDownLine = string('\u252C')
	var constCrossLine = string('\u253C')
	var constVerticalLine = string('\u2502')
	var constHalfCrossRightLine = string('\u251C')
	var constHalfCrossUpLine = string('\u2534')

	fmt.Println()
	fmt.Println(strings.Repeat(constHorizontalLine, 94))
	fmt.Println(fmt.Sprintf(" %s", thisRef.Description))

	fmt.Println(strings.Repeat(constHorizontalLine, 10) + constHalfCrossDownLine + strings.Repeat(constHorizontalLine, 83))
	fmt.Println(fmt.Sprintf("    Usage %s %s", constVerticalLine, strings.TrimSpace(usageString)))
	fmt.Println(fmt.Sprintf("          ") + constVerticalLine)

	if len(definedFlags) > 0 {
		fmt.Println(strings.Repeat(constHorizontalLine, 10) + constCrossLine + strings.Repeat(constHorizontalLine, 83))
		fmt.Print(fmt.Sprintf("    Flags " + constVerticalLine))

		pDefinedFlags := paddedFlags(definedFlags)
		for i, definedFlag := range pDefinedFlags {
			if i == 0 {
				fmt.Println(fmt.Sprintf(" %s, type=%s, required=%s, default=%s, %s", definedFlag.name, definedFlag.typeName, definedFlag.isRequired, definedFlag.defaultValue, definedFlag.description))
			} else {
				fmt.Println(fmt.Sprintf("          %s %s, type=%s, required=%s, default=%s, %s", constVerticalLine, definedFlag.name, definedFlag.typeName, definedFlag.isRequired, definedFlag.defaultValue, definedFlag.description))
			}
		}
		fmt.Println(fmt.Sprintf("          ") + constVerticalLine)
	}

	if !areTheseGlobalFlags {
		rootCmd := thisRef.parentCommand
		for {
			if rootCmd.parentCommand == nil {
				break
			}
			rootCmd = rootCmd.parentCommand
		}

		pDefinedFlags := paddedFlags(rootCmd.getDefinedFlags())

		if len(definedFlags) > 0 {
			fmt.Println("          " + constHalfCrossRightLine + strings.Repeat(constHorizontalLine, 10) + " global flags " + strings.Repeat(constHorizontalLine, 59))
			for _, definedFlag := range pDefinedFlags {
				fmt.Println(fmt.Sprintf("          %s %s, type=%s, required=%s, default=%s, %s", constVerticalLine, definedFlag.name, definedFlag.typeName, definedFlag.isRequired, definedFlag.defaultValue, definedFlag.description))
			}
			fmt.Println(fmt.Sprintf("          ") + constVerticalLine)
		}
	}

	if len(thisRef.Examples) > 0 {
		fmt.Println(strings.Repeat(constHorizontalLine, 10) + constCrossLine + strings.Repeat(constHorizontalLine, 83))
		fmt.Print(fmt.Sprintf(" Examples " + constVerticalLine))
		for i, example := range thisRef.Examples {
			if i == 0 {
				fmt.Println(fmt.Sprintf(" %s", example))
			} else {
				fmt.Println(fmt.Sprintf("          %s %s", constVerticalLine, example))
			}
		}
		fmt.Println(fmt.Sprintf("          ") + constVerticalLine)
	}

	if len(thisRef.subCommands) > 1 {
		fmt.Println(strings.Repeat(constHorizontalLine, 10) + constCrossLine + strings.Repeat(constHorizontalLine, 83))
		fmt.Print(fmt.Sprintf(" Commands " + constVerticalLine))
		firstOnePrinted := false
		pSubCommands := paddedCommands(thisRef.subCommands)
		for _, c := range pSubCommands {
			if c.Name != helpCmd.Name {
				if !firstOnePrinted {
					fmt.Println(fmt.Sprintf(" %s, %s", c.Name, c.Description))
					firstOnePrinted = true
				} else {
					fmt.Println(fmt.Sprintf("          %s %s, %s", constVerticalLine, c.Name, c.Description))
				}
			}
		}
		fmt.Println(fmt.Sprintf("          ") + constVerticalLine)
	}

	fmt.Println(strings.Repeat(constHorizontalLine, 10) + constHalfCrossUpLine + strings.Repeat(constHorizontalLine, 83))
	fmt.Println()
}

func paddedFlags(input []flag) []flag {
	definedFlagNameMaxLength := 0
	definedFlagTypeNameMaxLength := 0
	definedFlagIsRequiredMaxLength := 0
	definedFlagDefaultValueMaxLength := 0
	definedFlagDescriptionMaxLength := 0

	for _, val := range input {
		if len(flagPatterns[0]+val.name) > definedFlagNameMaxLength {
			definedFlagNameMaxLength = len(flagPatterns[0] + val.name)
		}
		if len(val.typeName) > definedFlagTypeNameMaxLength {
			definedFlagTypeNameMaxLength = len(val.typeName)
		}
		if len(val.isRequired) > definedFlagIsRequiredMaxLength {
			definedFlagIsRequiredMaxLength = len(val.isRequired)
		}
		if len(val.defaultValue) > definedFlagDefaultValueMaxLength {
			definedFlagDefaultValueMaxLength = len(val.defaultValue)
		}
		if len(val.description) > definedFlagDescriptionMaxLength {
			definedFlagDescriptionMaxLength = len(val.description)
		}
	}

	output := []flag{}
	for _, definedFlag := range input {
		definedFlagPaddedName := fmt.Sprintf("%"+strconv.Itoa(-definedFlagNameMaxLength)+"s", flagPatterns[0]+definedFlag.name)
		definedFlagPaddedTypeName := fmt.Sprintf("%"+strconv.Itoa(-definedFlagTypeNameMaxLength)+"s", definedFlag.typeName)
		definedFlagPaddedIsRequired := fmt.Sprintf("%"+strconv.Itoa(-definedFlagIsRequiredMaxLength)+"s", definedFlag.isRequired)
		definedFlagPaddedDefaultValue := fmt.Sprintf("%"+strconv.Itoa(-definedFlagDefaultValueMaxLength)+"s", definedFlag.defaultValue)
		definedFlagPaddedDescription := fmt.Sprintf("%"+strconv.Itoa(-definedFlagDescriptionMaxLength)+"s", definedFlag.description)

		output = append(output, flag{
			name:         definedFlagPaddedName,
			typeName:     definedFlagPaddedTypeName,
			isRequired:   definedFlagPaddedIsRequired,
			defaultValue: definedFlagPaddedDefaultValue,
			description:  definedFlagPaddedDescription,
		})
	}

	return output
}

func paddedCommands(input []*Command) []Command {
	definedCommandNameMaxLength := 0
	definedCommandDescriptionMaxLength := 0

	for _, val := range input {
		if val.Name != helpCmd.Name {
			if len(val.Name) > definedCommandNameMaxLength {
				definedCommandNameMaxLength = len(val.Name)
			}
			if len(val.Description) > definedCommandDescriptionMaxLength {
				definedCommandDescriptionMaxLength = len(val.Description)
			}
		}
	}

	output := []Command{}
	for _, val := range input {
		if val.Name != helpCmd.Name {
			definedCommandPaddedName := fmt.Sprintf("%"+strconv.Itoa(-definedCommandNameMaxLength)+"s", val.Name)
			definedCommandPaddedDescription := fmt.Sprintf("%"+strconv.Itoa(-definedCommandDescriptionMaxLength)+"s", val.Description)

			output = append(output, Command{
				Name:        definedCommandPaddedName,
				Description: definedCommandPaddedDescription,
			})
		}
	}

	return output
}
