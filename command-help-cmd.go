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
	var constMaxLineLength = 89

	fmt.Println()
	fmt.Println(strings.Repeat(constHorizontalLine, constMaxLineLength))
	fmt.Println(fmt.Sprintf(" %s", thisRef.Description))

	fmt.Println(strings.Repeat(constHorizontalLine, 10) + constHalfCrossDownLine + strings.Repeat(constHorizontalLine, constMaxLineLength-11))
	fmt.Println(fmt.Sprintf("    Usage %s %s", constVerticalLine, strings.TrimSpace(usageString)))

	if len(definedFlags) > 0 {
		updatedDefinedFlags := []flag{}
		updatedDefinedFlags = append(updatedDefinedFlags, flag{
			name:         "Name",
			typeName:     "Type",
			isRequired:   "Required",
			defaultValue: "Default",
			description:  "Description",
		})
		updatedDefinedFlags = append(updatedDefinedFlags, definedFlags...)

		if !areTheseGlobalFlags {
			rootCmd := thisRef.parentCommand
			for {
				if rootCmd.parentCommand == nil {
					break
				}
				rootCmd = rootCmd.parentCommand
			}
			updatedDefinedFlags = append(updatedDefinedFlags, flag{description: "-=GFLAGS=-"})
			updatedDefinedFlags = append(updatedDefinedFlags, rootCmd.getDefinedFlags()...)
		}

		pDefinedFlags := paddedFlags(updatedDefinedFlags)

		fmt.Println(strings.Repeat(constHorizontalLine, 10) + constCrossLine + strings.Repeat(constHorizontalLine, constMaxLineLength-11))
		fmt.Print(fmt.Sprintf("    Flags " + constVerticalLine))

		for i, definedFlag := range pDefinedFlags {
			definedFlag.name = " " + definedFlag.name
			if i == 0 {
				lenOfAllColumns := len("          ") + 5 + len(definedFlag.name) + 2 + len(definedFlag.typeName) + 2 + len(definedFlag.isRequired) + 2 + len(definedFlag.defaultValue) + 2
				fmt.Println(fmt.Sprintf(" %s "+constVerticalLine+" %s "+constVerticalLine+" %s "+constVerticalLine+" %s "+constVerticalLine+" %s", definedFlag.name, definedFlag.typeName, definedFlag.isRequired, definedFlag.defaultValue, definedFlag.description))
				fmt.Println("          " + constVerticalLine + "  " +
					strings.Repeat(constHorizontalLine, len(definedFlag.name)) + constCrossLine +
					strings.Repeat(constHorizontalLine, len(definedFlag.typeName)+2) + constCrossLine +
					strings.Repeat(constHorizontalLine, len(definedFlag.isRequired)+2) + constCrossLine +
					strings.Repeat(constHorizontalLine, len(definedFlag.defaultValue)+2) + constCrossLine +
					strings.Repeat(constHorizontalLine, constMaxLineLength-lenOfAllColumns) +
					"")
			} else {
				if strings.TrimSpace(definedFlag.description) == "-=GFLAGS=-" {
					fmt.Println("          " + constHalfCrossRightLine + strings.Repeat(constHorizontalLine, constMaxLineLength-11))
				} else {
					fmt.Println(fmt.Sprintf("          %s %s "+constVerticalLine+" %s "+constVerticalLine+" %s "+constVerticalLine+" %s "+constVerticalLine+" %s", constVerticalLine, definedFlag.name, definedFlag.typeName, definedFlag.isRequired, definedFlag.defaultValue, definedFlag.description))
				}
			}
		}
	}

	// if !areTheseGlobalFlags {
	// 	rootCmd := thisRef.parentCommand
	// 	for {
	// 		if rootCmd.parentCommand == nil {
	// 			break
	// 		}
	// 		rootCmd = rootCmd.parentCommand
	// 	}

	// 	updatedDefinedFlags := []flag{}
	// 	updatedDefinedFlags = append(updatedDefinedFlags, flag{
	// 		name:         "Name",
	// 		typeName:     "Type",
	// 		isRequired:   "Required",
	// 		defaultValue: "Default",
	// 		description:  "Description",
	// 	})
	// 	updatedDefinedFlags = append(updatedDefinedFlags, rootCmd.getDefinedFlags()...)

	// 	pDefinedFlags := paddedFlags(updatedDefinedFlags)

	// 	if len(definedFlags) > 0 {
	// 		fmt.Println("          " + constHalfCrossRightLine + strings.Repeat(constHorizontalLine, constMaxLineLength-9))
	// 		fmt.Println("          " + constHalfCrossRightLine + strings.Repeat(constHorizontalLine, 10) + " global flags " + strings.Repeat(constHorizontalLine, constMaxLineLength-35))
	// 		fmt.Println("          " + constHalfCrossRightLine + strings.Repeat(constHorizontalLine, constMaxLineLength-9))
	// 		for _, definedFlag := range pDefinedFlags {
	// 			fmt.Println(fmt.Sprintf("          %s %s | %s | %s | %s | %s", constVerticalLine, definedFlag.name, definedFlag.typeName, definedFlag.isRequired, definedFlag.defaultValue, definedFlag.description))
	// 		}
	// 	}
	// }

	if len(thisRef.Examples) > 0 {
		fmt.Println(strings.Repeat(constHorizontalLine, 10) + constCrossLine + strings.Repeat(constHorizontalLine, constMaxLineLength-11))
		fmt.Print(fmt.Sprintf(" Examples " + constVerticalLine))
		for i, example := range thisRef.Examples {
			if i == 0 {
				fmt.Println(fmt.Sprintf(" %s", example))
			} else {
				fmt.Println(fmt.Sprintf("          %s %s", constVerticalLine, example))
			}
		}
	}

	if len(thisRef.subCommands) > 1 {
		fmt.Println(strings.Repeat(constHorizontalLine, 10) + constCrossLine + strings.Repeat(constHorizontalLine, constMaxLineLength-11))
		fmt.Print(fmt.Sprintf(" Commands " + constVerticalLine))
		firstOnePrinted := false
		pSubCommands := paddedCommands(thisRef.subCommands)
		for _, c := range pSubCommands {
			if c.Name != helpCmd.Name {
				if !firstOnePrinted {
					fmt.Println(fmt.Sprintf(" %s | %s", c.Name, c.Description))
					firstOnePrinted = true
				} else {
					fmt.Println(fmt.Sprintf("          %s %s | %s", constVerticalLine, c.Name, c.Description))
				}
			}
		}
	}

	fmt.Println(strings.Repeat(constHorizontalLine, 10) + constHalfCrossUpLine + strings.Repeat(constHorizontalLine, constMaxLineLength-11))
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
	for i, definedFlag := range input {
		flagPrefix := ""
		if i != 0 {
			flagPrefix = flagPatterns[0]
		}

		definedFlagPaddedName := fmt.Sprintf("%"+strconv.Itoa(-definedFlagNameMaxLength)+"s", flagPrefix+definedFlag.name)
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
