package tests

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"

	clicmdflags "github.com/codemodify/systemkit-clicmdflags"
)

type cmdFlags struct {
	Name        string `flagName:"name"        flagRequired:"false"  flagDescription:"Service name"`
	Description string `flagName:"description" flagRequired:"false" flagDescription:"Service description"`
	Executable  string `flagName:"executable"  flagRequired:"false"  flagDescription:"Service executable"`
	Args        string `flagName:"args"        flagRequired:"false" flagDescription:"Executable args"`
	JSON        bool   `flagName:"json"        flagDefault:"false"  flagDescription:"Enables JSON output"`
	Verbose     bool   `flagName:"verbose"     flagDefault:"false"  flagDescription:"Enables verbose output"`
}

func Test00_SingleCommand(t *testing.T) {
	os.Args = append(os.Args, strings.Split("", " ")...)

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

	cmd.AddCommand(&clicmdflags.Command{
		Name:        "scan",
		Description: "Scan",
		Handler: func(command *clicmdflags.Command) {
			fmt.Println("ScanHandler()")
		},
	})

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}

	a, _ := json.Marshal(cmd.Flags)
	fmt.Println(string(a))
}
