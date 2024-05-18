package main

import (
	"fmt"
	"os"
)

// I NEED MY C++ MACROS BRUHHHH

const (
	ExitSuccess = 0
	ExitFailure = 1
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Fatal error encountered: no input files")
		fmt.Fprintln(os.Stderr, "compilation killed.")
		fmt.Fprintln(os.Stderr, "Usage:")
		fmt.Fprintln(os.Stderr, "diddl <input> -o <output>")
		os.Exit(ExitFailure)
	} else {
		fmt.Println("Booty")
		os.Exit(ExitSuccess)
	}
}
