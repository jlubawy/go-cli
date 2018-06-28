// TODO: Copyright goes here

package main

import (
    "flag"

    "github.com/jlubawy/go-cli"
)

var echoOptions = struct {
    // TODO: add any needed options
}{}

var echoCommand = cli.Command{
    Name:             "echo",
    ShortDescription: "",
    Description:      "",
    ShortUsage:       "",
    SetupFlags: func(fs *flag.FlagSet) {
        // TODO: Setup any flags
    },
    Run: func(args []string) {
        // TODO: Program goes here
    },
}
