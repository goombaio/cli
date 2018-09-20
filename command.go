// Copyright 2018, Goomba project Authors. All rights reserved.
//
// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with this
// work for additional information regarding copyright ownership.  The ASF
// licenses this file to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.  See the
// License for the specific language governing permissions and limitations
// under the License.

package cli

import (
	"flag"
	"io"
	"os"
	"text/template"
)

const (
	// UsageTemplate is the template being used to render the Usage for
	// any cli.Command that has a flag -h or-help attached to it.
	UsageTemplate = `usage: {{.CommandName}} [-help] <command> [args]{{if .LongDescription}}

  {{.LongDescription}}.{{end}}
{{if .Commands}}
Commands:
{{range .Commands}}  {{.Name}}	{{.ShortDescription}}
{{end}}{{end}}
Flags:
  -h, -help	Show help{{if .Commands}}

Use {{.CommandName}} [command] -help for more information about a command.{{end}}
`
)

// Command type implements a command or subcommand.
//
// A command is just that, a command for your application.
// E.g.  'go run ...' - 'run' is the command.
type Command struct {
	// ShortDescription is the message shown in the usage output when using the
	// flag -h or --help.
	ShortDescription string

	// LongDescription is the long message shown when the flag -h or -help is ç
	// used in combination with the command '<this-command> -h' or
	// '<this-command> -help' output.
	LongDescription string

	// commands are the list of subcommands that a command have associated with
	// it.
	commands []*Command

	// commands are the list of flags that a command have associated with it.
	flags *flag.FlagSet

	// Run is the actual work that the command will do when it is invoked.
	Run func(c *Command) error
}

// NewCommand creates a new Command.
//
// Cli requires you to define the name and description as part of your command
// definition to ensure usability.
func NewCommand(name string, shortDescription string) *Command {
	cmd := &Command{
		ShortDescription: shortDescription,
		LongDescription:  "",

		commands: make([]*Command, 0),
		flags:    flag.NewFlagSet(name, flag.ContinueOnError),

		Run: func(c *Command) error { return nil },
	}
	cmd.flags.SetOutput(os.Stderr)

	return cmd
}

// Name returns the name of the command.
func (c *Command) Name() string {
	return c.flags.Name()
}

// Args returns the list of arguments of this command.
//
// Arguments are represented as a list of strings.
func (c *Command) Args() []string {
	return c.flags.Args()
}

// Arg returns an string that represents an argument of this command given a
// numerical index.
func (c *Command) Arg(id int) string {
	return c.flags.Arg(id)
}

// Output retuns the destination for usage and error messages of this command.
//
// By default a Command uses os.Stderr as output.
func (c *Command) Output() io.Writer {
	return c.flags.Output()
}

// SetOutput sets the destination for usage and error messages.
//
// If output is nil, os.Stderr is used.
func (c *Command) SetOutput(output io.Writer) {
	c.flags.SetOutput(output)
}

// AddCommand adds a subCommand to this Command.
func (c *Command) AddCommand(cmd *Command) {
	c.commands = append(c.commands, cmd)
}

// Execute executes the command.
//
// Execute uses the args (os.Args[1:] by default) and run through the command
// tree finding appropriate matches for commands and then corresponding flags.
func (c *Command) Execute() error {
	// By default rootCommand (level 0)
	cmd := c

	// Find subCommand
	if len(os.Args) > 1 {

		// subCommand level 1
		for _, subCommand := range c.commands {
			if subCommand.Name() == os.Args[1] {
				cmd = subCommand
			}
		}

	}

	flag.Parse()
	flag.Usage = cmd.Usage

	err := cmd.flags.Parse(flag.Args())
	if err != nil {
		return err
	}
	cmd.flags.Usage = cmd.Usage

	err = cmd.Run(cmd)

	return err
}

// Usage puts out the usage for the command.
//
// It is used when a user provides invalid input or when the flag -h or -help
// is attached in the input.
func (c *Command) Usage() {
	templateData := struct {
		CommandName     string
		LongDescription string
		Commands        []struct {
			Name             string
			ShortDescription string
		}
	}{
		CommandName:     c.flags.Name(),
		LongDescription: c.LongDescription,
	}

	for _, subCommand := range c.commands {
		subc := struct {
			Name             string
			ShortDescription string
		}{
			subCommand.Name(),
			subCommand.ShortDescription,
		}
		templateData.Commands = append(templateData.Commands, subc)
	}

	t := template.Must(template.New("usageTemplate").Parse(UsageTemplate))
	_ = t.Execute(c.flags.Output(), templateData)
}
