package argparser

import (
	// "fmt"
	"os"
)

type CommandType int

const (
	StringParser      CommandType = 0
	MultiStringParser             = 1
)

type Command struct {
	short  string
	long   string
	opts   Options
	ctype  CommandType
	result interface{}
	found  *bool
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
	return Parser{
		name:        name,
		description: description,
	}
}

func (p *Parser) GetArgc() int {
	return p.argc
}

func (p *Parser) Parse() error {
	args := os.Args
	p.argc = len(args)

	// no args, no parse - simple
	if p.GetArgc() == 0 {
		return nil
	}

	if len(*p.commands) == 0 {
		return nil
	}

	/*
		for _, cmd := range *p.commands {
			if cmd.ctype == StringParser {
				fmt.Println(cmd.short)
			}
		}
	*/

	return nil
}

func (p *Parser) String(short string, long string, opts *Options) (result *string, found *bool) {
	cmd := Command{
		result: result,
		found:  found,
		ctype:  StringParser,
		short:  short,
		long:   long,
		opts:   *opts,
	}
	*p.commands = append(*p.commands, cmd)

	return result, found
}

func (p *Parser) MultiString(short string, long string, opts *Options) (value []string, found bool) {
	return []string{}, false
}
