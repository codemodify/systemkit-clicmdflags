package clicmdflags

import (
	"fmt"
	"strconv"
	"strings"
)

func (thisRef *Command) showUsage() {
	// get all commands
	commandsWithNoSubCommands := []*Command{}
	commandsWithSubCommands := []*Command{}
	subCommands := []*Command{}

	for _, cmd := range thisRef.subCommands {
		if len(cmd.subCommands) > 0 {
			commandsWithSubCommands = append(commandsWithSubCommands, cmd)
			for _, subCmd := range cmd.subCommands {
				subCommands = append(subCommands, subCmd)
			}
		} else {
			commandsWithNoSubCommands = append(commandsWithNoSubCommands, cmd)
		}
	}

	paddedCommandsWithNoSubCommands := paddedCommands(commandsWithNoSubCommands)
	paddedCommandsWithSubCommands := paddedCommands(commandsWithSubCommands)
	paddedSubCommands := paddedCommands(subCommands)

	// get all flags
	definedFlags := thisRef.getDefinedFlags()
	areTheseGlobalFlags := (thisRef.parentCommand == nil)
	updatedDefinedFlags := []flag{}
	updatedDefinedFlags = append(updatedDefinedFlags, flag{
		name:         "Name",
		typeName:     "Type",
		isRequired:   "Required",
		defaultValue: "Default",
		description:  "Description",
	})
	if len(definedFlags) > 0 {
		updatedDefinedFlags = append(updatedDefinedFlags, definedFlags...)
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
			updatedDefinedFlags = append(updatedDefinedFlags, flag{description: "-=GFLAGS=-"})
			updatedDefinedFlags = append(updatedDefinedFlags, definedFlags...)
		}
	}
	flags := []flag{}
	if len(updatedDefinedFlags) > 1 {
		flags = paddedFlags(updatedDefinedFlags)
	}

	// get the longest string line
	longestStringFromCommandsAndFlags := 0
	longestCommandOrFlagName := 0
	for _, sc := range paddedCommandsWithNoSubCommands {
		l := len(sc.Name) + len(sc.Description)
		if longestStringFromCommandsAndFlags < l {
			longestStringFromCommandsAndFlags = l
		}
		if longestCommandOrFlagName < len(sc.Name) {
			longestCommandOrFlagName = len(sc.Name)
		}
	}
	for _, sc := range paddedCommandsWithSubCommands {
		l := len(sc.Name) + len(sc.Description)
		if longestStringFromCommandsAndFlags < l {
			longestStringFromCommandsAndFlags = l
		}
		if longestCommandOrFlagName < len(sc.Name) {
			longestCommandOrFlagName = len(sc.Name)
		}
	}
	for _, sc := range paddedSubCommands {
		l := len(sc.Name) + len(sc.Description)
		if longestStringFromCommandsAndFlags < l {
			longestStringFromCommandsAndFlags = l
		}
		if longestCommandOrFlagName < len(sc.Name) {
			longestCommandOrFlagName = len(sc.Name)
		}
	}
	for _, fl := range flags {
		l := len(fl.name) + len(fl.typeName) + len(fl.isRequired) + len(fl.defaultValue) + len(fl.description)
		if longestStringFromCommandsAndFlags < l {
			longestStringFromCommandsAndFlags = l
		}
		if longestCommandOrFlagName < len(fl.name) {
			longestCommandOrFlagName = len(fl.name)
		}
	}

	//
	var constThinHorizontalLine = string('\u2500')
	var constThickHorizontalLine = string('\u2501')
	var constHalfCrossDownLine = string('\u252F')
	var constCrossLine = string('\u253C')
	var constCrossLine2 = string('\u253F')
	var constVerticalLine = string('\u2502')
	// var constHalfCrossRightLine = string('\u251C')
	var constMaxLineLength = longestStringFromCommandsAndFlags + (5 * 6)
	var constShortLineLength = constMaxLineLength - 2

	// START
	usageString := fmt.Sprintf(" %s COMMAND(s) %sFLAG(s)", thisRef.Name, flagPatterns[0])
	cmd := thisRef.parentCommand
	for {
		if cmd == nil {
			break
		}

		usageString = " " + cmd.Name + usageString

		cmd = cmd.parentCommand
	}

	// header
	fmt.Println()
	fmt.Println(fmt.Sprintf(" %s", thisRef.Description))

	fmt.Println(strings.Repeat(constThickHorizontalLine, 10) + constHalfCrossDownLine + strings.Repeat(constThickHorizontalLine, constMaxLineLength-11))
	fmt.Println(fmt.Sprintf("    Usage %s %s", constVerticalLine, strings.TrimSpace(usageString)))
	fmt.Println(strings.Repeat(constThickHorizontalLine, 10) + constCrossLine2 + strings.Repeat(constThickHorizontalLine, constMaxLineLength-11))

	// commands
	if len(thisRef.subCommands) > 1 {
		fmt.Print(fmt.Sprintf(" Commands " + constVerticalLine))
		firstOnePrinted := false

		for _, c := range paddedCommandsWithNoSubCommands {
			if !c.Hidden {
				if !firstOnePrinted {
					fmt.Println(fmt.Sprintf("  %s "+constThinHorizontalLine+" %s", fmt.Sprintf("%"+strconv.Itoa(-longestCommandOrFlagName)+"s", c.Name), c.Description))
					firstOnePrinted = true
				} else {
					fmt.Println(fmt.Sprintf("          %s  %s "+constThinHorizontalLine+" %s", constVerticalLine, fmt.Sprintf("%"+strconv.Itoa(-longestCommandOrFlagName)+"s", c.Name), c.Description))
				}
			}
		}

		for _, c := range paddedCommandsWithSubCommands {
			if !c.Hidden {
				commandDisplayData := c.Description
				if !c.PassThrough {
					commandDisplayData = c.Description + " (" + strings.TrimSpace(c.Name) + ")"
				}

				// pring sub-command as header
				if !firstOnePrinted {
					// fmt.Println(" " + strings.Repeat(constThinHorizontalLine, constShortLineLength-12))
					fmt.Println(fmt.Sprintf(" %s ", commandDisplayData))
					fmt.Println(fmt.Sprintf("          %s ", constVerticalLine) + strings.Repeat(constThinHorizontalLine, constShortLineLength-12))
					firstOnePrinted = true
				} else {
					fmt.Println(fmt.Sprintf("          %s ", constVerticalLine))
					// fmt.Println(fmt.Sprintf("          %s ", constVerticalLine) + strings.Repeat(constThinHorizontalLine, constShortLineLength-12))
					fmt.Println(fmt.Sprintf("          %s  %s ", constVerticalLine, commandDisplayData))
					fmt.Println(fmt.Sprintf("          %s ", constVerticalLine) + strings.Repeat(constThinHorizontalLine, constShortLineLength-12))
				}

				// for _, originalC := range commandsWithSubCommands {
				// 	if originalC.Name == strings.TrimSpace(c.Name) {
				paddedSubCommands := paddedCommands(c.subCommands)
				for _, subC := range paddedSubCommands {
					if !subC.Hidden {
						fmt.Println(fmt.Sprintf("          %s  %s %s %s", constVerticalLine, fmt.Sprintf("%"+strconv.Itoa(-longestCommandOrFlagName)+"s", subC.Name), constThinHorizontalLine, subC.Description))
					}
				}

				// 		break
				// 	}
				// }

			}
		}

		fmt.Println(strings.Repeat(" ", 10) + constVerticalLine + strings.Repeat(" ", constShortLineLength-11))
	}

	// flags
	if len(flags) > 0 {
		if len(thisRef.subCommands) > 1 {
			fmt.Println(strings.Repeat(constThinHorizontalLine, 10) + constCrossLine + strings.Repeat(constThinHorizontalLine, constShortLineLength-11))
		}
		fmt.Print(fmt.Sprintf("    Flags " + constVerticalLine))

		globalFlagsStarted := false
		for i, definedFlag := range flags {
			if definedFlag.isHidden == "true" {
				continue
			}

			definedFlag.name = fmt.Sprintf("%"+strconv.Itoa(-longestCommandOrFlagName)+"s", definedFlag.name)
			definedFlag.name = " " + definedFlag.name

			lenOfAllColumns := len("          ") + 5 + len(definedFlag.name) + 2 + len(definedFlag.typeName) + 2 + len(definedFlag.isRequired) + 2 + len(definedFlag.defaultValue) + 2
			if i == 0 {
				fmt.Println(fmt.Sprintf(" %s "+constVerticalLine+" %s "+constVerticalLine+" %s "+constVerticalLine+" %s "+constVerticalLine+" %s", definedFlag.name, definedFlag.typeName, definedFlag.isRequired, definedFlag.defaultValue, definedFlag.description))
				fmt.Println("          " + constVerticalLine + "  " +
					strings.Repeat(constThinHorizontalLine, len(definedFlag.name)) + constCrossLine +
					strings.Repeat(constThinHorizontalLine, len(definedFlag.typeName)+2) + constCrossLine +
					strings.Repeat(constThinHorizontalLine, len(definedFlag.isRequired)+2) + constCrossLine +
					strings.Repeat(constThinHorizontalLine, len(definedFlag.defaultValue)+2) + constCrossLine +
					strings.Repeat(constThinHorizontalLine, constShortLineLength-lenOfAllColumns) +
					"")
			} else {
				if strings.TrimSpace(definedFlag.description) == "-=GFLAGS=-" {
					globalFlagsStarted = true

					if i > 1 { // in case the command does not have any defined flags
						fmt.Println(fmt.Sprintf("          %s  %s%s%s%s%s%s%s%s%s",
							constVerticalLine, strings.Repeat(constThinHorizontalLine, len(definedFlag.name)), // constHalfCrossRightLine
							constCrossLine, strings.Repeat(constThinHorizontalLine, len(definedFlag.typeName)+2),
							constCrossLine, strings.Repeat(constThinHorizontalLine, len(definedFlag.isRequired)+2),
							constCrossLine, strings.Repeat(constThinHorizontalLine, len(definedFlag.defaultValue)+2),
							constCrossLine, strings.Repeat(constThinHorizontalLine, constShortLineLength-lenOfAllColumns),
						))
					}
				} else {
					if globalFlagsStarted {
						fmt.Println(fmt.Sprintf("  Globals %s %s "+constVerticalLine+" %s "+constVerticalLine+" %s "+constVerticalLine+" %s "+constVerticalLine+" %s", constVerticalLine, definedFlag.name, definedFlag.typeName, definedFlag.isRequired, definedFlag.defaultValue, definedFlag.description))
						globalFlagsStarted = false
					} else {
						fmt.Println(fmt.Sprintf("          %s %s "+constVerticalLine+" %s "+constVerticalLine+" %s "+constVerticalLine+" %s "+constVerticalLine+" %s", constVerticalLine, definedFlag.name, definedFlag.typeName, definedFlag.isRequired, definedFlag.defaultValue, definedFlag.description))
					}
				}
			}
		}
	}

	// examples
	if len(thisRef.Examples) > 0 {
		fmt.Println(strings.Repeat(constThinHorizontalLine, 10) + constCrossLine + strings.Repeat(constThinHorizontalLine, constShortLineLength-11))
		fmt.Print(fmt.Sprintf(" Examples " + constVerticalLine))
		for i, example := range thisRef.Examples {
			if i == 0 {
				fmt.Println(fmt.Sprintf(" %s", example))
			} else {
				fmt.Println(fmt.Sprintf("          %s %s", constVerticalLine, example))
			}
		}
	}

	fmt.Println(strings.Repeat(" ", 10) + constVerticalLine + strings.Repeat(" ", constShortLineLength-11))
	fmt.Println()
}

func paddedFlags(input []flag) []flag {
	if input != nil && len(input) <= 0 {
		return []flag{}
	}

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
			isHidden:     definedFlag.isHidden,
		})
	}

	return output
}

func paddedCommands(input []*Command) []Command {
	if input != nil && len(input) <= 0 {
		return []Command{}
	}

	definedCommandNameMaxLength := 0
	definedCommandDescriptionMaxLength := 0

	for _, val := range input {
		if len(val.Name) > definedCommandNameMaxLength {
			definedCommandNameMaxLength = len(val.Name)
		}
		if len(val.Description) > definedCommandDescriptionMaxLength {
			definedCommandDescriptionMaxLength = len(val.Description)
		}
	}

	output := []Command{}
	for _, val := range input {
		definedCommandPaddedName := fmt.Sprintf("%"+strconv.Itoa(-definedCommandNameMaxLength)+"s", val.Name)
		definedCommandPaddedDescription := fmt.Sprintf("%"+strconv.Itoa(-definedCommandDescriptionMaxLength)+"s", val.Description)

		output = append(output, Command{
			Name:        definedCommandPaddedName,
			Description: definedCommandPaddedDescription,
			Hidden:      val.Hidden,
			PassThrough: val.PassThrough,
			subCommands: val.subCommands,
		})
	}

	return output
}
