package tests

import (
	"fmt"
	"os"
	"path/filepath"

	clicmdflags "github.com/codemodify/systemkit-clicmdflags"
)

// ~~~~ root ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~

type rootCmdFlags struct {
	RootCmdFlag1 bool `flagName:"rootCmdFlag1" flagDescription:"rootCmdFlag1 description" flagDefault:"false"`
	RootCmdFlag2 bool `flagName:"rootCmdFlag2" flagDescription:"rootCmdFlag2 description" flagDefault:"false"`
	RootCmdFlag3 bool `flagName:"rootCmdFlag3" flagDescription:"rootCmdFlag3 description" flagDefault:"false"`
	RootCmdFlag4 bool `flagName:"rootCmdFlag4" flagDescription:"rootCmdFlag4 description" flagDefault:"false"`
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
	OneCmdFlag1 bool `flagName:"oneCmdFlag1" flagDescription:"oneCmdFlag1 description" flagDefault:"false"`
	OneCmdFlag2 bool `flagName:"oneCmdFlag2" flagDescription:"oneCmdFlag2 description" flagDefault:"false"`
	OneCmdFlag3 bool `flagName:"oneCmdFlag3" flagDescription:"oneCmdFlag3 description" flagDefault:"false"`
	OneCmdFlag4 bool `flagName:"oneCmdFlag4" flagDescription:"oneCmdFlag4 description" flagDefault:"false"`
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
	TwoCmdFlag1 bool `flagName:"twoCmdFlag1" flagDescription:"twoCmdFlag1 description" flagDefault:"false"`
	TwoCmdFlag2 bool `flagName:"twoCmdFlag2" flagDescription:"twoCmdFlag2 description" flagDefault:"false"`
	TwoCmdFlag3 bool `flagName:"twoCmdFlag3" flagDescription:"twoCmdFlag3 description" flagDefault:"false"`
	TwoCmdFlag4 bool `flagName:"twoCmdFlag4" flagDescription:"twoCmdFlag4 description" flagDefault:"false"`
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
	ThreeCmdFlag1 bool `flagName:"threeCmdFlag1" flagDescription:"threeCmdFlag1 description" flagDefault:"false"`
	ThreeCmdFlag2 bool `flagName:"threeCmdFlag2" flagDescription:"threeCmdFlag2 description" flagDefault:"false"`
	ThreeCmdFlag3 bool `flagName:"threeCmdFlag3" flagDescription:"threeCmdFlag3 description" flagDefault:"false"`
	ThreeCmdFlag4 bool `flagName:"threeCmdFlag4" flagDescription:"threeCmdFlag4 description" flagDefault:"false"`
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
