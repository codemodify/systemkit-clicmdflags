package tests

import (
	"os"
	"strings"
	"testing"

	clicmdflags "github.com/codemodify/systemkit-clicmdflags"
)

func Test10_CommandWithRequriedFlags_test(t *testing.T) {
	args := ""
	// args := "-fourCmdFlags1 -fourCmdFlags3 -fourCmdFlags4 true"
	// args := "help"
	os.Args = append(os.Args, strings.Split(args, " ")...)

	if err := fourCmd.Execute(); err != nil {
		t.Fatal(err.Error())
	}

	clicmdflags.DEBUGDumpCommandFlags(fourCmd)
}
