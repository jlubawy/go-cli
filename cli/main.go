// Copyright 2018 Josh Lubawy. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/jlubawy/go-cli"
)

var program = cli.Program{
	Name:        "cli",
	Description: "Cli is a tool for creating new command-line interface (CLI) programs.",
	Commands: []cli.Command{
		generateCommand,
	},
}

func main() { program.RunAndExit() }
