package argparser

import (
	"os"
	"testing"
)

// TestStringParser
// Should parse the provided string.
func TestStringParser(t *testing.T) {
	name := "argparser string test"
	description := "Testing parsing of string types"

	argument := "001 my list.txt"

	os.Args = append(os.Args, "--list")
	os.Args = append(os.Args, argument)

	parser := NewParser(name, description)
	result, found := parser.String("-l", "--list", &Options{Required: false, Help: "Printing a list."})

	err := parser.Parse()
	if err != nil {
		t.Fatalf("parser.Parse() failed, %s\n", err.Error())
	}

	if !*found {
		t.Fatalf("parser.String() failed. want: true, got: %t\n", *found)
	}

	if *result != argument {
		t.Fatalf("parser.String() failed. want: %s, got: %s\n", argument, *result)
	}

}
