package main

import (
	"fmt"
	"os"

	"github.com/mikeunge/argparser"
)

const (
	APP_NAME = "MultiStringParser"
	APP_DESC = "Multi string parser example."
)

func main() {
	parser := argparser.NewParser(APP_NAME, APP_DESC)
	result, found := parser.MultiString("-l", "--list", &argparser.Options{Required: false, Help: "Echo's whatever you provide."})

	if err := parser.Parse(); err != nil {
		parser.Usage(err.Error())
		os.Exit(1)
	}

	if !*found {
		fmt.Println("No argument provided for --list, goodbye!")
		os.Exit(0)
	}

	for _, res := range *result {
		fmt.Println(res)
	}
}
