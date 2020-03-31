package tests

import (
	"os"
	"strings"
	"testing"

	clicmdflags "github.com/codemodify/systemkit-clicmdflags"
)

func Test04_ThreeCommands(t *testing.T) {
	args := "-rootCmdFlag1 true -rootCmdFlag2 true -rootCmdFlag3 true -rootCmdFlag4 true"
	os.Args = append(os.Args, strings.Split(args, " ")...)

	args = "oneCmd -oneCmdFlag1 true -oneCmdFlag2 true -oneCmdFlag3 true -oneCmdFlag4 true"
	os.Args = append(os.Args, strings.Split(args, " ")...)

	args = "twoCmd -twoCmdFlag1 true -twoCmdFlag2 true -twoCmdFlag3 true -twoCmdFlag4 true"
	os.Args = append(os.Args, strings.Split(args, " ")...)

	args = "threeCmd -threeCmdFlag1 true -threeCmdFlag2 true -threeCmdFlag3 true -threeCmdFlag4 true"
	os.Args = append(os.Args, strings.Split(args, " ")...)

	rootCmd.AddCommand(oneCmd)
	rootCmd.AddCommand(twoCmd)
	rootCmd.AddCommand(threeCmd)

	rootCmd.Execute()

	clicmdflags.DEBUGDumpCommandFlags(rootCmd)
}
