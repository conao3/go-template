package main

import (
	"fmt"
	"os"

	flags "github.com/jessevdk/go-flags"
)

type Options struct {
	Verbose []bool `short:"v" long:"verbose" description:"Show verbose debug information"`
}

func main() {
	var opts Options
	fmt.Println("Hello World")
	args, err := flags.Parse(&opts)
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(args)
	fmt.Println(opts)
}
