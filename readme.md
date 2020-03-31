# ![](https://fonts.gstatic.com/s/i/materialicons/bookmarks/v4/24px.svg) CLI Flags
[![GoDoc](https://godoc.org/github.com/codemodify/systemkit-logging?status.svg)](https://godoc.org/github.com/codemodify/systemkit-events)
[![0-License](https://img.shields.io/badge/license-0--license-brightgreen)](https://github.com/codemodify/TheFreeLicense)
[![Go Report Card](https://goreportcard.com/badge/github.com/codemodify/systemkit-logging)](https://goreportcard.com/report/github.com/codemodify/systemkit-logging)
[![Test Status](https://github.com/danawoodman/systemservice/workflows/Test/badge.svg)](https://github.com/danawoodman/systemservice/actions)
![code size](https://img.shields.io/github/languages/code-size/codemodify/SystemKit?style=flat-square)

#### Robust CLI commands and flags for your Go app. Elegant + the smallest footprint there is.

# ![](https://fonts.gstatic.com/s/i/materialicons/bookmarks/v4/24px.svg) Why this is better than `libX` or `libY`
- In comparison to `spf13/cobra` + `spf13/pflag`
	- __clean, lean, simple and small footprint code for similar functionality__
	- __global flags work in front and at the end__
	- uses native __Go structs__ defined by the user
	- multiple comands in one line, each with its own flags
	- memory is freed after you execute a command, it matters with 10+ commands and the app is a daemon

- In comparison to `Golang flag`
	- __clean, lean, simple and small footprint code for similar functionality__
	- __global flags work in front and at the end__
	- uses native __Go structs__ defined by the user
	- multiple comands in one line, each with its own flags
	- has __command__ and __sub-commands__ concep


# ![](https://fonts.gstatic.com/s/i/materialicons/bookmarks/v4/24px.svg) Install
```go
go get github.com/codemodify/systemkit-clicmdflags
```

# ![](https://fonts.gstatic.com/s/i/materialicons/bookmarks/v4/24px.svg) Example
```go
import (
	"fmt"

	clicmdflags "github.com/codemodify/systemkit-clicmdflags"
)

// AppRootCmdFlags -
type AppRootCmdFlags struct {
	JSON    bool `flagName:"json"    flagDescription:"Enables JSON output"    flagDefault:"false"`
	Verbose bool `flagName:"verbose" flagDescription:"Enables verbose output" flagDefault:"false"`
}

// ExtendedInfoCmdFlags -
type ExtendedInfoCmdFlags struct {
	DumpCPUInfo bool `flagName:"dumpCpuInfo" flagDescription:"Outputs also CPU info" flagDefault:"true"`
}

func main() {
	var appRootCmd = &clicmdflags.Command{
		Name:        filepath.Base(os.Args[0]),
		Description: "Displays PC information",
		Flags:       AppRootCmdFlags{},
	}

	appRootCmd.AddCommand(&clicmdflags.Command{
		Name:        "version",
		Description: "Displays product version",
		Handler: func(command *clicmdflags.Command) {
			fmt.Println("v1.0")
		},
	})

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

	appRootCmd.Execute()
}
```

![alt text](https://raw.githubusercontent.com/codemodify/systemkit-clicmdflags/master/.dox/01.png)
![alt text](https://raw.githubusercontent.com/codemodify/systemkit-clicmdflags/master/.dox/02.png)
![alt text](https://raw.githubusercontent.com/codemodify/systemkit-clicmdflags/master/.dox/03.png)
