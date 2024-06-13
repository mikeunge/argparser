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
	if err := parser.parseString(cmd, args); err != nil {
		t.Fatalf("parser.parseString() failed,  %s\n", err)
	}

	if result != want {
		t.Fatalf("parser.parseString() failed. want: %s, got: %s\n", want, result)
	}

}
