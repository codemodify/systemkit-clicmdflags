package tests

import (
	"os"
	"strings"
	"testing"

	clicmdflags "github.com/codemodify/systemkit-clicmdflags"
)

func Test08_OneCommandMissingFlagValue(t *testing.T) {
	args := "-rootCmdFlag1 true -rootCmdFlag2 true -rootCmdFlag3 true -rootCmdFlag4 true"
	os.Args = append(os.Args, strings.Split(args, " ")...)

	args = "oneCmd -oneCmdFlag1 true -oneCmdFlag2 true -oneCmdFlag3 true -oneCmdFlag4 true"
	os.Args = append(os.Args, strings.Split(args, " ")...)

	rootCmd.AddCommand(oneCmd)

	rootCmd.Execute()

	clicmdflags.DEBUGDumpCommandFlags(rootCmd)
}
