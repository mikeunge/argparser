# argparser

> This is a small library that makes it easier for you to parse command line arguments.

This project is created because of the need for a good, easy to use and straight forward arg parser in go.

It is heavily inspired by **akamensky**'s [argparse library](https://github.com/akamensky/argparse) that I've used heavily in all my cli projects.
The only reason I wanted to create my own was because of some quirks I find kind of annoying, like the usage of _positional arguments_ and _lists_.

### Installation

To install and use _argparser_ run:

`$ go get github.com/mikeunge/argparser`

### Usage

**argparser** has a nice caviat to it, if a flag is detected, it parses **TIL** another flag is encountered.
This means, you can do something like: `$ mycli --foo some strings --bar also some strings`, this will get parsed as `foo=[some, string]` & `bar=[also, some, string]`.
It might seem obvious but this is an issue I had with **akamensky**'s parser and optionals weren't an option.
