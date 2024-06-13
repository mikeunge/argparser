package argparser

import "fmt"

func (p *Parser) parseString(cmd Command, args []string) error {
	// found := false
	for i := 0; i < len(args); i++ {
		if args[i] == cmd.long || args[i] == cmd.short {
			if i+1 >= len(args) {
				return fmt.Errorf("String parser expects a value after the flag, nothing found.")
			}
			fmt.Println(args[i+1])
		}
	}
	return nil
}
