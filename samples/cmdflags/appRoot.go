package cmdflags

import (
	"os"
	"path/filepath"

	clicmdflags "github.com/codemodify/systemkit-clicmdflags"
)

// AppRootCmdFlags -
type AppRootCmdFlags struct {
	JSON    bool `flagName:"json"    flagDefault:"false" flagDescription:"Enables JSON output"`
	Verbose bool `flagName:"verbose" flagDefault:"false" flagDescription:"Enables verbose output"`
}

var appRootCmd = &clicmdflags.Command{
	Name:        filepath.Base(os.Args[0]),
	Description: "Displays PC information",
	Examples: []string{
		filepath.Base(os.Args[0]) + " -json",
		filepath.Base(os.Args[0]) + " -json true",
	},
	Flags: AppRootCmdFlags{},
}

// Execute - this is a convenience call
func Execute() error {
	return appRootCmd.Execute()
}
