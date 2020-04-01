package main

import (
	"log"

	"github.com/codemodify/systemkit-clicmdflags/samples/cmdflags"
)

func main() {
	if err := cmdflags.Execute(); err != nil {
		log.Fatal(err)
	}
}
