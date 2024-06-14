package argparser

import (
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
		short:  "--test",
		long:   "-t",
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
		short:  "--test",
		long:   "-t",
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
// Should parse multiple strings..
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
		short:  "--test",
		long:   "-t",
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
