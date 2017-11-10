package main

import (
	"fmt"
	"os"

	flags "github.com/jessevdk/go-flags"
)

func main() {
	fmt.Printf("Hello, world.\n")

	args, err := flags.ParseArgs(&opts, os.Args)
}
