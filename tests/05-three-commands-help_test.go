package tests

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func Test05_ThreeCommandsHelp(t *testing.T) {
	rootCmd.AddCommand(oneCmd)
	rootCmd.AddCommand(twoCmd)
	rootCmd.AddCommand(threeCmd)

	fmt.Println()
	args := "help"
	os.Args = append(os.Args, strings.Split(args, " ")...)
	rootCmd.Execute()
}
