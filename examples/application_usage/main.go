package main

import (
	"fmt"
	"os"

	"github.com/mikeunge/argparser"
)

const (
	APP_NAME = "My app"
	APP_DESC = "Example of how the parser could be used in a real app."
)

var (
	debugInfo  = false
	configPath = "./config.yml"
)

func dbgPrint(msg string) {
	if debugInfo {
		fmt.Printf("DEBUG: %s\n", msg)
	}
}

func main() {
	results := make(map[string]interface{}, 0)
	found := make(map[string]*bool, 0)

	parser := argparser.NewParser(APP_NAME, APP_DESC)
	_, found["debug"] = parser.Flag("", "--debug", &argparser.Options{Required: false, Help: "Print debug information."})
	results["config"], found["config"] = parser.String("-c", "--config", &argparser.Options{Required: false, Help: "Path to config file."})
	results["names"], found["names"] = parser.MultiString("-n", "--names", &argparser.Options{Required: false, Help: "A list of names to print."})

	if err := parser.Parse(); err != nil {
		fmt.Printf("Parser returned with error: %s\n", err.Error())
		parser.PrintHelp()
		os.Exit(1)
	}

	if *found["debug"] {
		debugInfo = true
	}

	if *found["config"] {
		configPath = *results["config"].(*string)
		dbgPrint(fmt.Sprintf("Loading config from: %s", *results["config"].(*string)))
	}

	if !*found["names"] {
		fmt.Println("\nNo names to print, aborting.")
		os.Exit(0)
	}

	for _, name := range *results["names"].(*[]string) {
		fmt.Println(name)
	}
}
