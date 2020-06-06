package cmdflags

import (
	"fmt"
	"os"
	"path/filepath"

	clicmdflags "github.com/codemodify/systemkit-clicmdflags"
)

func init() {

	cmd0 := &clicmdflags.Command{
		Description:   "Device",
		IsPassThrough: true,
	}

	cmd1 := &clicmdflags.Command{
		Name:        "add",
		Description: "Adds a device",
		Examples: []string{
			filepath.Base(os.Args[0]) + " device add",
		},
		Handler: func(command *clicmdflags.Command) {
			fmt.Println("add")
		},
	}
	cmd2 := &clicmdflags.Command{
		Name:        "remove",
		Description: "Remove a device",
		Examples: []string{
			filepath.Base(os.Args[0]) + " device remove",
		},
		Handler: func(command *clicmdflags.Command) {
			fmt.Println("remove")
		},
	}
	cmd3 := &clicmdflags.Command{
		Name:        "transfer",
		Description: "Transfers a device",
		Examples: []string{
			filepath.Base(os.Args[0]) + "device transfer",
		},
		Handler: func(command *clicmdflags.Command) {
			fmt.Println("transfer")
		},
	}

	cmd0.AddCommand(cmd1)
	cmd0.AddCommand(cmd2)
	cmd0.AddCommand(cmd3)

	appRootCmd.AddCommand(cmd0)
}
