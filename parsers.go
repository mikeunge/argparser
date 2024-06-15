package argparser

import (
	"fmt"
	"strconv"
)

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
			next := i + 1
			if next >= len(args) {
				return fmt.Errorf("string parser expects a value after the flag, nothing found.")
			}
			*cmd.found = true
			*cmd.result.(*string) = args[next]
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
					return nil
				}

				*cmd.found = true
				*cmd.result.(*[]string) = append(*cmd.result.(*[]string), sub[j])
			}
			return nil
		}
	}
	return fmt.Errorf("nothing to parse.")
}

func (p *Parser) parseFlag(cmd *Command, args []string) error {
	for i := 0; i < len(args); i++ {
		if args[i] == cmd.long || args[i] == cmd.short {
			*cmd.found = true
			*cmd.result.(*bool) = true
			break
		}
	}
	return nil
}

func (p *Parser) parseNumber(cmd *Command, args []string) error {
	for i := 0; i < len(args); i++ {
		if args[i] == cmd.long || args[i] == cmd.short {
			next := i + 1
			if next >= len(args) {
				return fmt.Errorf("number parser expects a value after the flag, nothing found.")
			}

			var err error
			if *cmd.result.(*int), err = strconv.Atoi(args[next]); err != nil {
				return fmt.Errorf("argument is not of type int.")
			}
			*cmd.found = true
			return nil
		}
	}
	return fmt.Errorf("nothing to parse.")
}

func (p *Parser) parseMultiNumber(cmd *Command, args []string) error {
	for i := 0; i < len(args); i++ {
		if args[i] == cmd.long || args[i] == cmd.short {
			next := i + 1
			if next >= len(args) {
				return fmt.Errorf("number parser expects a value after the flag, nothing found.")
			}

			sub := args[next:]
			for j := 0; j < len(sub); j++ {
				if isCommand(p.commands, sub[j]) {
					return nil
				}

				num, err := strconv.Atoi(sub[j])
				if err != nil {
					return fmt.Errorf("argument is not of type int.")
				}

				*cmd.result.(*[]int) = append(*cmd.result.(*[]int), num)
				*cmd.found = true
			}
			return nil
		}
	}
	return fmt.Errorf("nothing to parse.")
}
