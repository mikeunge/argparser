package main

import (
	"fmt"
	"os"

	"github.com/mikeunge/argparser"
)

const (
	APP_NAME = "StringParser"
	APP_DESC = "String parser example."
)

func main() {
	parser := argparser.NewParser(APP_NAME, APP_DESC)
	result, found := parser.String("-e", "--echo", &argparser.Options{Required: false, Help: "Echo's whatever you provide."})

	if err := parser.Parse(); err != nil {
		parser.Usage(err.Error())
		os.Exit(1)
	}

	if !*found {
		fmt.Println("\nNo argument provided for --echo, goodbye!")
		os.Exit(0)
	}
	fmt.Printf("%s\n", *result)
}
