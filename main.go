package main

import (
	"fmt"
	"os"

	"github.com/blinktag/elepher/repl"
	"github.com/blinktag/elepher/version"
)

func main() {
	fmt.Printf("Go Hypertext Preprocessor. Version %s\n",
		version.Version)
	repl.Start(os.Stdin, os.Stdout)
}
