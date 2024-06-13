package argparser

import (
	"os"
)

type CommandType int

const (
	StringParser      CommandType = 0
	MultiStringParser             = 1
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
		if cmd.commandType == StringParser {
			if err := p.parseString(cmd, args); err != nil {
				return err
			}
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

func (p *Parser) MultiString(short string, long string, opts *Options) (value []string, found bool) {
	return []string{}, false
}
