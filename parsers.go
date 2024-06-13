package argparser

import "fmt"

func (p *Parser) parseString(cmd Command, args []string) error {
	// found := false
	for i := 0; i < len(args); i++ {
		if args[i] == cmd.long || args[i] == cmd.short {
			if i+1 >= len(args) {
				return fmt.Errorf("string parser expects a value after the flag, nothing found.")
			}
			*cmd.found = true
			*cmd.result.(*string) = args[i+1]
			return nil
		}
	}
	return fmt.Errorf("nothing to parse.")
}
