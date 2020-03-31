package cmdflags

import (
	"fmt"

	clicmdflags "github.com/codemodify/systemkit-clicmdflags"
)

// ExtendedInfoCmdFlags -
type ExtendedInfoCmdFlags struct {
	DumpCPUInfo bool `flagName:"dumpCpuInfo" flagDescription:"Outputs also CPU info" flagDefault:"true"`
}

func init() {
	appRootCmd.AddCommand(&clicmdflags.Command{
		Name:        "info",
		Description: "Displays extended information",
		Flags:       ExtendedInfoCmdFlags{},
		Handler: func(command *clicmdflags.Command) {
			fmt.Println("SSD size is 1TB")

			if flags, ok := command.Flags.(ExtendedInfoCmdFlags); ok && flags.DumpCPUInfo {
				fmt.Println("CPU is 64bit capable")
			}
		},
	})
}
