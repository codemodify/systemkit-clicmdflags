package cmdflags

import (
	"fmt"
	"os"
	"path/filepath"

	clicmdflags "github.com/codemodify/systemkit-clicmdflags"
)

func init() {

	cmd0 := &clicmdflags.Command{
		Name:        "service",
		Description: "Services",
	}

	cmd1 := &clicmdflags.Command{
		Name:        "start",
		Description: "Star the service",
		Examples: []string{
			filepath.Base(os.Args[0]) + " service start",
		},
		Handler: func(command *clicmdflags.Command) {
			fmt.Println("start")
		},
	}
	cmd2 := &clicmdflags.Command{
		Name:        "stop",
		Description: "Stop the service",
		Examples: []string{
			filepath.Base(os.Args[0]) + " service stop",
		},
		Handler: func(command *clicmdflags.Command) {
			fmt.Println("stop")
		},
	}
	cmd3 := &clicmdflags.Command{
		Name:        "restart",
		Description: "Restart the service",
		Examples: []string{
			filepath.Base(os.Args[0]) + "service restart",
		},
		Handler: func(command *clicmdflags.Command) {
			fmt.Println("restart")
		},
	}

	cmd0.AddCommand(cmd1)
	cmd0.AddCommand(cmd2)
	cmd0.AddCommand(cmd3)

	appRootCmd.AddCommand(cmd0)
}
