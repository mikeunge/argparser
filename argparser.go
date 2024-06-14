package argparser

import (
	"os"
)

type CommandType int

const (
	StringParser      CommandType = 0
	MultiStringParser             = 1
	FlagParser                    = 2
	NumberParser                  = 3
	MultiNumberParser             = 4
)

type Command struct {
	result      interface{}
	found       *bool
	short       string
	long        string
	opts        Options
	commandType CommandType
}

type Options struct {
	Required bool
	Help     string
	Default  interface{}
}

type Parser struct {
	name        string
	description string
	commands    *[]Command
	argc        int
	printHelp   bool
}

func NewParser(name string, description string) Parser {
	cmds := make([]Command, 0)
	return Parser{
		name:        name,
		description: description,
		printHelp:   true,
		commands:    &cmds,
	}
}

func (p *Parser) DisableHelp() {
	p.printHelp = false
}

func (p *Parser) GetArgc() int {
	return p.argc
}

func (p *Parser) Parse() error {
	args := os.Args
	p.argc = len(args)

	if p.GetArgc() == 0 {
		return nil
	} else if len(*p.commands) == 0 {
		return nil
	}

	for _, cmd := range *p.commands {
		switch cmd.commandType {
		case StringParser:
			return p.parseString(&cmd, args)
		case MultiStringParser:
			return p.parseMultiString(&cmd, args)
		case FlagParser:
			return p.parseFlag(&cmd, args)
		case NumberParser:
			return p.parseNumber(&cmd, args)
		case MultiNumberParser:
			return p.parseMultiNumber(&cmd, args)
		default:
			return nil
		}
	}

	return nil
}

func (p *Parser) String(short string, long string, opts *Options) (*string, *bool) {
	var (
		result string
		found  bool
	)
	cmd := Command{
		result:      &result,
		found:       &found,
		commandType: StringParser,
		short:       short,
		long:        long,
		opts:        *opts,
	}
	*p.commands = append(*p.commands, cmd)

	return &result, &found
}

func (p *Parser) MultiString(short string, long string, opts *Options) (*[]string, *bool) {
	var found bool

	result := make([]string, 0)
	cmd := Command{
		result:      &result,
		found:       &found,
		commandType: MultiStringParser,
		short:       short,
		long:        long,
		opts:        *opts,
	}
	*p.commands = append(*p.commands, cmd)

	return &result, &found
}

func (p *Parser) Flag(short string, long string, opts *Options) (*bool, *bool) {
	var (
		result bool
		found  bool
	)

	cmd := Command{
		result:      &result,
		found:       &found,
		commandType: FlagParser,
		short:       short,
		long:        long,
		opts:        *opts,
	}
	*p.commands = append(*p.commands, cmd)

	return &result, &found
}

func (p *Parser) Number(short string, long string, opts *Options) (*int, *bool) {
	var (
		result int
		found  bool
	)

	cmd := Command{
		result:      &result,
		found:       &found,
		commandType: NumberParser,
		short:       short,
		long:        long,
		opts:        *opts,
	}
	*p.commands = append(*p.commands, cmd)

	return &result, &found
}

func (p *Parser) MultiNumber(short string, long string, opts *Options) (*[]int, *bool) {
	var (
		result []int
		found  bool
	)

	cmd := Command{
		result:      &result,
		found:       &found,
		commandType: MultiNumberParser,
		short:       short,
		long:        long,
		opts:        *opts,
	}
	*p.commands = append(*p.commands, cmd)

	return &result, &found
}
