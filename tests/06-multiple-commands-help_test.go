package tests

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func Test06_MultipleCommandsHelp(t *testing.T) {
	rootCmd.AddCommand(oneCmd)
	oneCmd.AddCommand(twoCmd)
	twoCmd.AddCommand(threeCmd)

	// 1
	fmt.Println()
	args := "help"
	os.Args = append(os.Args, strings.Split(args, " ")...)
	rootCmd.Execute()

	// 2
	fmt.Println()
	args = "help oneCmd"
	os.Args = append(os.Args, strings.Split(args, " ")...)
	rootCmd.Execute()

	// 3
	fmt.Println()
	args = "help oneCmd twoCmd"
	os.Args = append(os.Args, strings.Split(args, " ")...)
	rootCmd.Execute()

	// 4
	fmt.Println()
	args = "help oneCmd twoCmd threeCmd"
	os.Args = append(os.Args, strings.Split(args, " ")...)
	rootCmd.Execute()
}
