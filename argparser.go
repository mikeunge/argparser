package argparser

import (
	"fmt"
	"os"
	"strings"
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

func (p *Parser) Usage(info string) {
	if !p.printHelp {
		return
	}

	maxShort := 0
	maxLong := 0
	for _, c := range *p.commands {
		if len(c.short) > maxShort {
			maxShort = len(c.short)
		}
		if len(c.long) > maxLong {
			maxLong = len(c.long)
		}
	}

	formatHelpTable := func(cmd Command) string {
		// decide what option to display
		opt := ""
		if len(cmd.short) > 0 && len(cmd.short) > 0 {
			opt = fmt.Sprintf("%s %s", cmd.short, cmd.long)
		} else if len(cmd.short) > 0 {
			opt = fmt.Sprintf("%s", cmd.short)
		} else {
			space := strings.Repeat(" ", maxShort)
			opt = fmt.Sprintf("%s %s", space, cmd.long)
		}

		// make sure the option(s) always have the same lenght so it doesn't look so scuffed
		if len(opt) < (maxShort + maxLong + 1) {
			diff := maxShort + maxLong + 1 - len(opt)
			space := strings.Repeat(" ", diff)
			opt = fmt.Sprintf("%s%s", opt, space)
		}

		if cmd.opts.Required {
			return fmt.Sprintf("\t%s\t%s (required)", opt, cmd.opts.Help)
		}
		return fmt.Sprintf("\t%s\t%s", opt, cmd.opts.Help)
	}

	if len(info) > 0 {
		fmt.Println(info)
	}
	fmt.Printf("%s - %s\n\nAvailable options:\n", p.name, p.description)
	for _, cmd := range *p.commands {
		availableOption := formatHelpTable(cmd)
		fmt.Println(availableOption)
	}
}

func (p *Parser) Parse() error {
	args := os.Args
	p.argc = len(args) - 1 // remove one (1) because the first arg is always the app name

	if p.GetArgc() == 0 {
		p.Usage("No arguments provided.")
		return nil
	}

	// check if the user provided --help as flag
	for _, arg := range args {
		if arg == "--help" {
			p.Usage("")
			return nil
		}
	}

	for _, cmd := range *p.commands {
		var err error

		switch cmd.commandType {
		case StringParser:
			err = p.parseString(&cmd, args)
			break
		case MultiStringParser:
			err = p.parseMultiString(&cmd, args)
			break
		case FlagParser:
			err = p.parseFlag(&cmd, args)
			break
		case NumberParser:
			err = p.parseNumber(&cmd, args)
			break
		case MultiNumberParser:
			err = p.parseMultiNumber(&cmd, args)
			break
		default:
			return fmt.Errorf("Command %d is not a known command. Aborting.", cmd.commandType)
		}
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *Parser) String(short string, long string, opts *Options) (*string, *bool) {
	var (
		result string
		found  bool
	)

	command := makeCommand(short, long, &result, &found, StringParser, opts)
	*p.commands = append(*p.commands, command)

	return &result, &found
}

func (p *Parser) MultiString(short string, long string, opts *Options) (*[]string, *bool) {
	var found bool

	result := make([]string, 0)
	command := makeCommand(short, long, &result, &found, MultiStringParser, opts)
	*p.commands = append(*p.commands, command)

	return &result, &found
}

func (p *Parser) Flag(short string, long string, opts *Options) (*bool, *bool) {
	var (
		result bool
		found  bool
	)

	command := makeCommand(short, long, &result, &found, FlagParser, opts)
	*p.commands = append(*p.commands, command)

	return &result, &found
}

func (p *Parser) Number(short string, long string, opts *Options) (*int, *bool) {
	var (
		result int
		found  bool
	)

	command := makeCommand(short, long, &result, &found, NumberParser, opts)
	*p.commands = append(*p.commands, command)

	return &result, &found
}

func (p *Parser) MultiNumber(short string, long string, opts *Options) (*[]int, *bool) {
	var found bool
	result := make([]int, 0)

	command := makeCommand(short, long, &result, &found, MultiNumberParser, opts)
	*p.commands = append(*p.commands, command)

	return &result, &found
}

func makeCommand(short string, long string, result interface{}, found *bool, cmdType CommandType, opts *Options) Command {
	return Command{
		result:      result,
		found:       found,
		commandType: cmdType,
		short:       short,
		long:        long,
		opts:        *opts,
	}
}

func determineUsedFlag(c *Command) string {
	if len(c.long) > 0 {
		return c.long
	} else if len(c.short) > 0 {
		return c.short
	}
	return "No flag"
}

func requiredCheck(c *Command) error {
	if *&c.opts.Required && !*c.found {
		return fmt.Errorf("%s is required but was not provided. Check with --help for more information.", determineUsedFlag(c))
	}
	return nil
}
