package tests

import (
	"os"
	"strings"
	"testing"

	clicmdflags "github.com/codemodify/systemkit-clicmdflags"
)

func Test07_RootCommandMissingFlagValue(t *testing.T) {
	args := "-rootCmdFlag1 -rootCmdFlag2 true -rootCmdFlag3 true -rootCmdFlag4 true"
	os.Args = append(os.Args, strings.Split(args, " ")...)

	rootCmd.Execute()

	clicmdflags.DEBUGDumpCommandFlags(rootCmd)
}
