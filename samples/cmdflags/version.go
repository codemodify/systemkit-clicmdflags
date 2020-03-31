package cmdflags

import (
	"fmt"

	clicmdflags "github.com/codemodify/systemkit-clicmdflags"
)

func init() {
	appRootCmd.AddCommand(&clicmdflags.Command{
		Name:        "version",
		Description: "Displays product version",
		Handler: func(command *clicmdflags.Command) {
			fmt.Println("v1.0")
		},
	})
}
