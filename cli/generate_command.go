// Copyright 2018 Josh Lubawy. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/jlubawy/go-cli"
)

var generateOptions = struct {
	Dir string
}{}

var generateCommand = cli.Command{
	Name:             "generate",
	ShortDescription: "generate Go source files for a new CLI tool",
	Description: `Generate Go source files for a new CLI tool. This includes the main.go file and
any *_command.go files to go with it.`,
	ShortUsage: "[program name] <command names...>",
	SetupFlags: func(fs *flag.FlagSet) {
		fs.StringVar(&generateOptions.Dir, "dir", "", "directory to generate files in")
	},
	Run: func(args []string) {
		if len(args) == 0 {
			cli.Fatal("Must provide a program name.\n")
		}

		var dir string
		if generateOptions.Dir == "" {
			cli.Fatal("Must provide an output directory [-dir].\n")
		} else {
			stat, err := os.Stat(generateOptions.Dir)
			if err != nil {
				cli.Fatalf("Error getting directory info: %v\n", err)
			}
			if !stat.IsDir() {
				cli.Fatal("Must specify a directory.\n")
			}
			dir = generateOptions.Dir
		}

		var templData = struct {
			Name         string
			CommandNames []string
		}{
			Name:         args[0],
			CommandNames: args[1:],
		}

		// Create and write the program file (main.go)
		pfn := filepath.Join(dir, "main.go")
		cli.Infof("Creating program file %s\n", pfn)
		f, err := os.OpenFile(pfn, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0664)
		if err != nil {
			cli.Fatalf("Error creating output file: %v\n", err)
		}
		defer f.Close()
		if err := programTempl.Execute(f, &templData); err != nil {
			cli.Fatalf("Error executing program template: %v\n", err)
		}

		// Create and write all command files
		for _, cn := range templData.CommandNames {
			cfn := filepath.Join(dir, fmt.Sprintf("%s_command.go", cn))
			cli.Infof("Creating command file %s\n", cfn)

			f, err := os.OpenFile(cfn, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0644)
			if err != nil {
				cli.Fatalf("Error creating command file: %v\n", err)
			}
			defer f.Close()
			if err := commandTempl.Execute(f, &cn); err != nil {
				cli.Fatalf("Error executing command template: %v\n", err)
			}
		}
	},
}

var programTempl = template.Must(template.New("").Parse(`// TODO: Copyright goes here

package main

import (
    "github.com/jlubawy/go-cli"
)

var program = cli.Program{
    Name:        "{{.Name}}",
    Description: "",
    Commands: []cli.Command{ {{range .CommandNames}}
        {{.}}Command,
{{- end}}
    },
}

func main() { program.RunAndExit() }
`))

var commandTempl = template.Must(template.New("").Parse(`// TODO: Copyright goes here

package main

import (
    "flag"

    "github.com/jlubawy/go-cli"
)

var {{.}}Options = struct {
    // TODO: add any needed options
}{}

var {{.}}Command = cli.Command{
    Name:             "{{.}}",
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
`))
