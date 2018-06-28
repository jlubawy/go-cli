// TODO: Copyright goes here

package main

import (
    "github.com/jlubawy/go-cli"
)

var program = cli.Program{
    Name:     "example",
    Commands: []cli.Command{ 
        echoCommand,
    },
}

func main() { program.RunAndExit() }
