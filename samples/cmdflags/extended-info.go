package cmdflags

// ExtendedInfoCmdFlags -
type ExtendedInfoCmdFlags struct {
	DumpCPUInfo bool `flagName:"dumpCpuInfo" flagRequired:"true" flagDescription:"Outputs also CPU info"`
}

func init() {
	// appRootCmd.AddCommand(&clicmdflags.Command{
	// 	Name:        "info",
	// 	Description: "Displays extended information",
	// 	Examples: []string{
	// 		filepath.Base(os.Args[0]) + " info",
	// 		filepath.Base(os.Args[0]) + " info -dumpCpuInfo",
	// 	},
	// 	Flags: ExtendedInfoCmdFlags{},
	// 	Handler: func(command *clicmdflags.Command) {
	// 		fmt.Println("SSD size is 1TB")

	// 		flags, ok := command.Flags.(ExtendedInfoCmdFlags)
	// 		if ok && flags.DumpCPUInfo {
	// 			fmt.Println("CPU is 64bit capable")
	// 		}
	// 	},
	// })
}
