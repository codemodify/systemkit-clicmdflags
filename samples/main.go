package main

import (
	"log"

	"github.com/codemodify/systemkit-clicmdflags/samples/cmdflags"
)

func main() {
	// os.Args = append(os.Args, "help", "info")
	if err := cmdflags.Execute(); err != nil {
		log.Fatal(err)
	}
}
