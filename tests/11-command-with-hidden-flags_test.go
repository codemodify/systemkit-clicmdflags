package tests

import (
	"os"
	"strings"
	"testing"

	clicmdflags "github.com/codemodify/systemkit-clicmdflags"
)

func Test11_CommandWithHiddenFlags_test(t *testing.T) {
	args := "help"
	os.Args = append(os.Args, strings.Split(args, " ")...)

	if err := fiveCmd.Execute(); err != nil {
		t.Fatal(err.Error())
	}

	clicmdflags.DEBUGDumpCommandFlags(fiveCmd)
}
