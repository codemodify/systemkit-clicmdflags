package cmdflags

import (
	"os"
	"path/filepath"

	clicmdflags "github.com/codemodify/systemkit-clicmdflags"
)

// AppRootCmdFlags -
type AppRootCmdFlags struct {
	JSON    bool `flagName:"json"    flagDescription:"Enables JSON output"    flagDefault:"false"`
	Verbose bool `flagName:"verbose" flagDescription:"Enables verbose output" flagDefault:"false"`
}

var appRootCmd = &clicmdflags.Command{
	Name:        filepath.Base(os.Args[0]),
	Description: "Displays PC information",
	Flags:       AppRootCmdFlags{},
}

// Execute - this is a convenience call
func Execute() {
	appRootCmd.Execute()
}
