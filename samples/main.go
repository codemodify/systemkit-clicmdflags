package main

import (
	"log"
	"os"

	"github.com/codemodify/systemkit-clicmdflags/samples/cmdflags"
)

func main() {
	os.Args = append(os.Args, "add")
	if err := cmdflags.Execute(); err != nil {
		log.Fatal(err)
	}
}
