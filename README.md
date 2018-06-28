# jlubawy/go-cli

[![GoDoc](https://godoc.org/github.com/jlubawy/go-cli?status.svg)](https://godoc.org/github.com/jlubawy/go-cli)
[![Build Status](https://travis-ci.org/jlubawy/go-cli.svg?branch=master)](https://travis-ci.org/jlubawy/go-cli)

Package cli provides a simple way of creating new command-line interface (CLI)
programs. It is built using nothing but standard packages and is based on the
behavior of the 'go' command-line tool (with some minor changes).

## Example

The following is a simple program that echoes text back to stdout.

```go
package main

import (
    "flag"
    "fmt"
    "strings"

    "github.com/jlubawy/go-cli"
)

type EchoOptions struct {
    Prefix string
}

var echoOptions EchoOptions

var program = cli.Program{
    Name:        "example",
    Description: "Example is a program demonstrating how to use the cli package.",
    Commands: []cli.Command{
        {
            Name:             "echo",
            ShortDescription: "echo the provided text",
            Description:      "Echo the provided text.",
            ShortUsage:       "[text to echo]",
            SetupFlags: func(fs *flag.FlagSet) {
                fs.StringVar(&echoOptions.Prefix, "prefix", "", "prefix to be added to the text")
            },
            Run: func(args []string) {
                var text string
                if len(args) >= 1 {
                    text = strings.Join(args, " ")
                }
                if echoOptions.Prefix == "" {
                    fmt.Println(text)
                } else {
                    fmt.Printf("%s: %s\n", echoOptions.Prefix, text)
                }
            },
        },
    },
}

func main() { program.RunAndExit() }

```

If run with no arguments:

    $ ./example
    Example is a program demonstrating how to use the cli package.

    Usage:

        example command [options]

    The commands are:

        echo       echo the provided text

    Use "example help [command]" for more Information about a command.

Help for a particular sub-command:

    $ ./example help echo
    Usage: example echo [text to echo]

    Echo the provided text.

    Options:

        -prefix     prefix to be added to the text (default=)

If run normally:

    $ ./example echo -prefix="Some prefix" Some random text
    Some prefix: Some random text
