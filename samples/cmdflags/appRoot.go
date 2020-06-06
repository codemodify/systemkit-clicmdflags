package cmdflags

import (
	"fmt"
	"os"
	"path/filepath"

	clicmdflags "github.com/codemodify/systemkit-clicmdflags"
)

// AppRootCmdFlags -
type AppRootCmdFlags struct {
	JSON    bool `flagName:"json"    flagDefault:"true" flagDescription:"Enables JSON output"`
	Verbose bool `flagName:"verbose" flagDescription:"Enables verbose output"`
}

var appRootCmd = &clicmdflags.Command{
	Name:        filepath.Base(os.Args[0]),
	Description: "Displays PC information",
	Examples: []string{
		filepath.Base(os.Args[0]) + " -json",
		filepath.Base(os.Args[0]) + " -json true",
	},
	Handler: func(command *clicmdflags.Command) {
		flags, ok := command.Flags.(AppRootCmdFlags)

		if ok && flags.JSON {
			fmt.Println("JSON is on")
		}
		if ok && flags.Verbose {
			fmt.Println("Verbose is on")
		}
	},
	Flags: AppRootCmdFlags{},
}

// Execute - this is a convenience call
func Execute() error {
	return appRootCmd.Execute()
}
