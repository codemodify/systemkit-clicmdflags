package tests

import (
	"fmt"
	"os"
	"path/filepath"

	clicmdflags "github.com/codemodify/systemkit-clicmdflags"
)

// ~~~~ root ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~

type rootCmdFlags struct {
	RootCmdFlag1 bool `flagName:"rootCmdFlag1" flagDefault:"false" flagDescription:"rootCmdFlag1 description"`
	RootCmdFlag2 bool `flagName:"rootCmdFlag2" flagDefault:"false" flagDescription:"rootCmdFlag2 description"`
	RootCmdFlag3 bool `flagName:"rootCmdFlag3" flagDefault:"false" flagDescription:"rootCmdFlag3 description"`
	RootCmdFlag4 bool `flagName:"rootCmdFlag4" flagDefault:"false" flagDescription:"rootCmdFlag4 description"`
}

var rootCmd = &clicmdflags.Command{
	Name:        filepath.Base(os.Args[0]),
	Flags:       &rootCmdFlags{},
	Description: "This is `rootCmd` description",
	Handler: func(thisCmd *clicmdflags.Command) {
		fmt.Println("EXEC `rootCmd`")
		clicmdflags.DEBUGDumpCommandFlags(thisCmd)
	},
}

// ~~~~ one  ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~

type oneCmdFlags struct {
	OneCmdFlag1 bool `flagName:"oneCmdFlag1" flagDefault:"false" flagDescription:"oneCmdFlag1 description"`
	OneCmdFlag2 bool `flagName:"oneCmdFlag2" flagDefault:"false" flagDescription:"oneCmdFlag2 description"`
	OneCmdFlag3 bool `flagName:"oneCmdFlag3" flagDefault:"false" flagDescription:"oneCmdFlag3 description"`
	OneCmdFlag4 bool `flagName:"oneCmdFlag4" flagDefault:"false" flagDescription:"oneCmdFlag4 description"`
}

var oneCmd = &clicmdflags.Command{
	Name:        "oneCmd",
	Flags:       oneCmdFlags{},
	Description: "This is `oneCmd` description",
	Handler: func(thisCmd *clicmdflags.Command) {
		fmt.Println("EXEC `oneCmd`")
		clicmdflags.DEBUGDumpCommandFlags(thisCmd)
	},
}

// ~~~~ two  ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~

type twoCmdFlags struct {
	TwoCmdFlag1 bool `flagName:"twoCmdFlag1" flagDefault:"false" flagDescription:"twoCmdFlag1 description"`
	TwoCmdFlag2 bool `flagName:"twoCmdFlag2" flagDefault:"false" flagDescription:"twoCmdFlag2 description"`
	TwoCmdFlag3 bool `flagName:"twoCmdFlag3" flagDefault:"false" flagDescription:"twoCmdFlag3 description"`
	TwoCmdFlag4 bool `flagName:"twoCmdFlag4" flagDefault:"false" flagDescription:"twoCmdFlag4 description"`
}

var twoCmd = &clicmdflags.Command{
	Name:        "twoCmd",
	Flags:       twoCmdFlags{},
	Description: "This is `twoCmd` description",
	Handler: func(thisCmd *clicmdflags.Command) {
		fmt.Println("EXEC `twoCmd`")
		clicmdflags.DEBUGDumpCommandFlags(thisCmd)
	},
}

// ~~~~ three ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~

type threeCmdFlags struct {
	ThreeCmdFlag1 bool `flagName:"threeCmdFlag1" flagDefault:"false" flagDescription:"threeCmdFlag1 description"`
	ThreeCmdFlag2 bool `flagName:"threeCmdFlag2" flagDefault:"false" flagDescription:"threeCmdFlag2 description"`
	ThreeCmdFlag3 bool `flagName:"threeCmdFlag3" flagDefault:"false" flagDescription:"threeCmdFlag3 description"`
	ThreeCmdFlag4 bool `flagName:"threeCmdFlag4" flagDefault:"false" flagDescription:"threeCmdFlag4 description"`
}

var threeCmd = &clicmdflags.Command{
	Name:        "threeCmd",
	Flags:       threeCmdFlags{},
	Description: "This is `threeCmd` description",
	Handler: func(thisCmd *clicmdflags.Command) {
		fmt.Println("EXEC `threeCmd`")
		clicmdflags.DEBUGDumpCommandFlags(thisCmd)
	},
}

// ~~~~ four ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~

type fourCmdFlags struct {
	FourCmdFlags1 bool `flagName:"fourCmdFlags1" flagRequired:"true" flagDescription:"fourCmdFlags1 description"`
	FourCmdFlags2 bool `flagName:"fourCmdFlags2" flagDefault:"false" flagDescription:"fourCmdFlags2 description"`
	FourCmdFlags3 bool `flagName:"fourCmdFlags3" flagRequired:"true" flagDescription:"fourCmdFlags3 description"`
	FourCmdFlags4 bool `flagName:"fourCmdFlags4" flagDefault:"false" flagDescription:"fourCmdFlags4 description"`
}

var fourCmd = &clicmdflags.Command{
	Name:        "fourCmd",
	Flags:       fourCmdFlags{},
	Description: "This is `fourCmd` description",
	Handler: func(thisCmd *clicmdflags.Command) {
		fmt.Println("EXEC `fourCmd`")
		clicmdflags.DEBUGDumpCommandFlags(thisCmd)
	},
}
