package argparser

import (
	"os"
	"testing"
)

// TestInitParser
// Should initialize the parser.
func TestInitParser(t *testing.T) {
	name := "argparser init test"
	description := "Testing the parser init + arg count functionality"

	parser := NewParser(name, description)

	os.Args = append(os.Args, "--help")
	err := parser.Parse()
	if err != nil {
		t.Fatalf("parser.parse() failed, %s\n", err.Error())
	}

	// Why 4? We only append "--help"?
	// go test adds 3 extra arguments:
	// 1. temp build path
	// 2. panic on exit
	// 3. max timeout
	// 4. finally our "--help" flag :)

	want := 4
	argc := parser.GetArgc()
	if argc != want {
		t.Fatalf("parser.getArgc() failed. want: %d, got: %d\n", want, argc)
	}
}

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
