package argparser

import "fmt"

func isCommand(cmds *[]Command, val string) bool {
	for _, cmd := range *cmds {
		if cmd.long == val || cmd.short == val {
			return true
		}
	}
	return false
}

func (p *Parser) parseString(cmd *Command, args []string) error {
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

func (p *Parser) parseMultiString(cmd *Command, args []string) error {
	for i := 0; i < len(args); i++ {
		if args[i] == cmd.long || args[i] == cmd.short {
			next := i + 1
			if next >= len(args) {
				return fmt.Errorf("multi string parser expects a value after the flag, nothing found.")
			}

			// we only need the FROM to END of the array to parse
			sub := args[next:]
			for j := 0; j < len(sub); j++ {
				if isCommand(p.commands, sub[j]) {
					continue
				}

				*cmd.found = true
				*cmd.result.(*[]string) = append(*cmd.result.(*[]string), sub[j])
			}
			return nil
		}
	}
	return fmt.Errorf("nothing to parse.")
}
