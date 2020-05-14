package tests

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"

	clicmdflags "github.com/codemodify/systemkit-clicmdflags"
)

type cmdFlags struct {
	Name        string `flagName:"name"        flagRequired:"true"  flagDescription:"Service name"`
	Description string `flagName:"description" flagRequired:"false" flagDescription:"Service description"`
	Executable  string `flagName:"executable"  flagRequired:"true"  flagDescription:"Service executable"`
	Args        string `flagName:"args"        flagRequired:"false" flagDescription:"Executable args"`
	JSON        bool   `flagName:"json"        flagDefault:"false"  flagDescription:"Enables JSON output"`
	Verbose     bool   `flagName:"verbose"     flagDefault:"false"  flagDescription:"Enables verbose output"`
}

func Test00_SingleCommand(t *testing.T) {
	os.Args = append(os.Args, "help")

	var cmd = &clicmdflags.Command{
		Name:        filepath.Base(os.Args[0]),
		Description: "Create a system service",
		Examples: []string{
			filepath.Base(os.Args[0]) + " -json",
			filepath.Base(os.Args[0]) + " -json true",
		},
		Flags: cmdFlags{},
		Handler: func(command *clicmdflags.Command) {
			fmt.Println("Handler()")
		},
	}

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
