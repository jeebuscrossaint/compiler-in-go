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
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "Fatal error encountered: wrong count of input files")
		fmt.Fprintln(os.Stderr, "compilation killed.")
		fmt.Fprintln(os.Stderr, "Usage:")
		fmt.Fprintln(os.Stderr, "diddl <input> -o <output>")
		os.Exit(ExitFailure)
	} else {
		file, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to open file:", err)
			os.Exit(ExitFailure)
		}
		defer file.Close()

		fmt.Println("Booty")
		os.Exit(ExitSuccess)
		// time 33:34
	}
}
