package argparser

import (
	"fmt"
	"testing"
)

// TestParseString
// Should retrieve the parsed element.
func TestParseString(t *testing.T) {
	var (
		parser Parser
		result string
		found  bool
	)

	want := "my-test-string"
	args := []string{"--test", want}

	cmd := Command{
		result: &result,
		found:  &found,
		long:   "--test",
		short:  "-t",
	}

	if err := parser.parseString(&cmd, args); err != nil {
		t.Fatalf("parser.parseString() failed,  %s\n", err)
	}

	if result != want {
		t.Fatalf("parser.parseString() failed. want: %s, got: %s\n", want, result)
	}

	if !found {
		t.Fatalf("parser.parseString() failed. want: true, got: %t\n", found)
	}
}

// TestFailParseString
// Parsing should fail do to abscence of parsable arguments.
func TestFailParseString(t *testing.T) {
	var (
		parser Parser
		result string
		found  bool
	)

	args := []string{"--test"}
	cmd := Command{
		result: &result,
		found:  &found,
		long:   "--test",
		short:  "-t",
	}

	var err error
	if err = parser.parseString(&cmd, args); err == nil {
		t.Fatalf("parser.parseString() should fail, but didn't, %s\n", result)
	}

	want := "string parser expects a value after the flag, nothing found."
	if err.Error() != want {
		t.Fatalf("parser.parseString() failed with wrong message, want: %s got: %s\n", want, err.Error())
	}
}

// TestParseMultiString
// Should parse multiple strings.
func TestParseMultiString(t *testing.T) {
	name := "argparser init test"
	description := "Testing the parser init + arg count functionality"

	parser := NewParser(name, description)
	result := make([]string, 0)
	want1 := "my-test-string"
	want2 := "my-second-test-string"
	args := []string{"--test", want1, want2}

	var found bool
	cmd := Command{
		result: &result,
		found:  &found,
		long:   "--test",
		short:  "-t",
	}

	*parser.commands = append(*parser.commands, cmd)
	if err := parser.parseMultiString(&cmd, args); err != nil {
		t.Fatalf("parser.parseMultiString() failed,  %s\n", err)
	}

	if !found {
		t.Fatalf("parser.parseMultiString() failed. want: true, got: %t\n", found)
	}

	for i := 0; i < len(result); i++ {
		if result[i] == want1 {
			continue
		}
		if result[i] == want2 {
			continue
		}
		t.Fatalf("parser.parseMultiString() failed. want: %s | %s, got: %s | %s\n", want1, want2, result[0], result[1])
	}
}

// TestParseFlag
// Should parse a flag.
func TestParseFlag(t *testing.T) {
	var (
		parser Parser
		result bool
		found  bool
	)

	cmd := Command{
		result: &result,
		found:  &found,
		long:   "--test",
		short:  "-t",
	}
	args := []string{"--test"}

	if err := parser.parseFlag(&cmd, args); err != nil {
		t.Fatalf("parser.parseFlag() should not fail, %s\n", err.Error())
	}

	if !found {
		t.Fatalf("parser.parseFlag() failed. want: true, got: %t\n", found)
	}

	if !result {
		t.Fatalf("parser.parseFlag() failed. want: true, got: %t\n", result)
	}
}

// TestParseNumber
// Should parse a number.
func TestParseNumber(t *testing.T) {
	var (
		parser Parser
		result int
		found  bool
	)
	want := 454

	cmd := Command{
		result: &result,
		found:  &found,
		long:   "--test",
		short:  "-t",
	}
	args := []string{"--test", fmt.Sprintf("%d", want)}

	if err := parser.parseNumber(&cmd, args); err != nil {
		t.Fatalf("parser.parseNumber() should not fail, %s\n", err.Error())
	}

	if !found {
		t.Fatalf("parser.parseNumber() failed. want: true, got: %t\n", found)
	}

	if result != want {
		t.Fatalf("parser.parseNumber() failed. want: %d, got: %d\n", want, result)
	}
}

// TestFailParseNumber
// Should fail parseing a string as number.
func TestFailParseNumber(t *testing.T) {
	var (
		parser Parser
		result int
		found  bool
	)

	cmd := Command{
		result: &result,
		found:  &found,
		long:   "--test",
		short:  "-t",
	}
	args := []string{"--test", "abc"}

	var err error
	if err = parser.parseNumber(&cmd, args); err == nil {
		t.Fatalf("parser.parseNumber() should fail, but didn't, %d\n", result)
	}

	want := "argument is not of type int."
	if err.Error() != want {
		t.Fatalf("parser.parseNumber() failed with wrong message, want: %s got: %s\n", want, err.Error())
	}
}

// TestParseMultiNumber
// Should parse an array of numbers.
func TestParseMultiNumber(t *testing.T) {
	name := "argparser init test"
	description := "Testing the parser init + arg count functionality"
	want1 := 454
	want2 := 12312

	parser := NewParser(name, description)
	args := []string{"--test", fmt.Sprintf("%d", want1), fmt.Sprintf("%d", want2)}
	result := make([]int, 0)

	var found bool
	cmd := Command{
		result: &result,
		found:  &found,
		long:   "--test",
		short:  "-t",
	}

	*parser.commands = append(*parser.commands, cmd)
	if err := parser.parseMultiNumber(&cmd, args); err != nil {
		t.Fatalf("parser.parseMultiNumber() failed,  %s\n", err)
	}

	if !found {
		t.Fatalf("parser.parseMultiNumber() failed. want: true, got: %t\n", found)
	}

	for i := 0; i < len(result); i++ {
		if result[i] == want1 {
			continue
		}
		if result[i] == want2 {
			continue
		}
		t.Fatalf("parser.parseMultiNumber() failed. want: %d | %d, got: %d | %d\n", want1, want2, result[0], result[1])
	}
}
