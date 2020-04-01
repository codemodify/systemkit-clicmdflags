# ![](https://fonts.gstatic.com/s/i/materialicons/bookmarks/v4/24px.svg) CLI Flags
[![GoDoc](https://godoc.org/github.com/codemodify/systemkit-logging?status.svg)](https://godoc.org/github.com/codemodify/systemkit-events)
[![0-License](https://img.shields.io/badge/license-0--license-brightgreen)](https://github.com/codemodify/TheFreeLicense)
[![Go Report Card](https://goreportcard.com/badge/github.com/codemodify/systemkit-logging)](https://goreportcard.com/report/github.com/codemodify/systemkit-logging)
[![Test Status](https://github.com/danawoodman/systemservice/workflows/Test/badge.svg)](https://github.com/danawoodman/systemservice/actions)
![code size](https://img.shields.io/github/languages/code-size/codemodify/systemkit-clicmdflags?style=flat-square)

#### Robust CLI commands and flags for your Go app. Elegant + the smallest footprint there is.

# ![](https://fonts.gstatic.com/s/i/materialicons/bookmarks/v4/24px.svg) What it does
my-cli-app `help` __PATH/TO/COMMAND__


# ![](https://fonts.gstatic.com/s/i/materialicons/bookmarks/v4/24px.svg) Why this is better than `libX` or `libY`
- In comparison to `spf13/cobra` + `spf13/pflag`
	- __clean, lean, simple and small footprint code with similar functionality__
	- __global flags are recognised regardless of the position: front, middle or at the end__
	- uses native __Go structs__ defined by the user
	- multiple comands in one line, each with its own flags
	- memory is freed after you execute a command, it matters if 10+ commands in a daemon

- In comparison to `Golang flag`
	- __clean, lean, simple and small footprint code with similar functionality__
	- __global flags are recognised regardless of the position: front, middle or at the end__
	- uses native __Go structs__ defined by the user
	- has __command__ and __sub-commands__ concept


# ![](https://fonts.gstatic.com/s/i/materialicons/bookmarks/v4/24px.svg) Install
```go
go get github.com/codemodify/systemkit-clicmdflags
```

# ![](https://fonts.gstatic.com/s/i/materialicons/bookmarks/v4/24px.svg) Example
```go
import (
	"fmt"
	"os"
	"path/filepath"
	"log"

	clicmdflags "github.com/codemodify/systemkit-clicmdflags"
)

// AppRootCmdFlags -
type AppRootCmdFlags struct {
	JSON    bool `flagName:"json"    flagDefault:"false" flagDescription:"Enables JSON output"`
	Verbose bool `flagName:"verbose" flagDefault:"false" flagDescription:"Enables verbose output"`
}

// ExtendedInfoCmdFlags -
type ExtendedInfoCmdFlags struct {
	DumpCPUInfo bool `flagName:"dumpCpuInfo" flagRequired:"true" flagDescription:"Outputs also CPU info"`
}

func main() {
	var appRootCmd = &clicmdflags.Command{
		Name:        filepath.Base(os.Args[0]),
		Description: "Displays PC information",
		Examples: []string{
			filepath.Base(os.Args[0]) + " -json",
			filepath.Base(os.Args[0]) + " -json true",
		},
		Flags: AppRootCmdFlags{},
	}

	appRootCmd.AddCommand(&clicmdflags.Command{
		Name:        "version",
		Description: "Displays product version",
		Examples: []string{
			filepath.Base(os.Args[0]) + " version",
		},
		Handler: func(command *clicmdflags.Command) {
			fmt.Println("v1.0")
		},
	})

	appRootCmd.AddCommand(&clicmdflags.Command{
		Name:        "info",
		Description: "Displays extended information",
		Examples: []string{
			filepath.Base(os.Args[0]) + " info",
			filepath.Base(os.Args[0]) + " info -dumpCpuInfo",
		},
		Flags: ExtendedInfoCmdFlags{},
		Handler: func(command *clicmdflags.Command) {
			fmt.Println("SSD size is 1TB")

			flags, ok := command.Flags.(ExtendedInfoCmdFlags)
			if ok && flags.DumpCPUInfo {
				fmt.Println("CPU is 64bit capable")
			}
		},
	})

	if err := appRootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
```

![](https://raw.githubusercontent.com/codemodify/systemkit-clicmdflags/master/.dox/01.png)
![](https://raw.githubusercontent.com/codemodify/systemkit-clicmdflags/master/.dox/02.png)
![](https://raw.githubusercontent.com/codemodify/systemkit-clicmdflags/master/.dox/03.png)
![](https://raw.githubusercontent.com/codemodify/systemkit-clicmdflags/master/.dox/04.png)
![](https://raw.githubusercontent.com/codemodify/systemkit-clicmdflags/master/.dox/05.png)
