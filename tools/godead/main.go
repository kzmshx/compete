package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s </path/to/main.go>\n", os.Args[0])
		os.Exit(1)
	}

	Run(os.Args[1])
}
