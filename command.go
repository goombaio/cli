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
	"io"
	"os"
	"text/template"
)

const (
	// UsageTemplate is the template being used to render the Usage for
	// any cli.Command that has a flag -h or-help attached to it.
	UsageTemplate = `usage: {{.Name}} [-help] <command> [args]{{if .LongDescription}}

  {{.LongDescription}}{{end}}
{{if .Commands}}
Commands:
{{range .Commands}}  {{.Name}}	{{.ShortDescription}}
{{end}}{{end}}{{if .Flags}}
Flags:
{{range .Flags}}  -{{.ShortName}}, -{{.LongName}}	{{.Description}}
{{end}}{{end}}
Use {{.Name}} [command] -help for more information about a command.
`
)

// Command type implements a command or subcommand.
//
// A command is just that, a command for your application.
// E.g.  'go run ...' - 'run' is the command.
type Command struct {
	// Name is the unique name of the command
	Name string

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

	// arguments are the list of arguments that a command have associated with
	// it.
	arguments []string

	// flags are the list of flags that a command have associated with it.
	flags []*Flag

	// Run is the actual work that the command will do when it is invoked.
	Run func(c *Command) error

	// output is where (an io.Writer) the reults will be printed
	output io.Writer
}

// NewCommand creates a new Command.
//
// Cli requires you to define the name and description as part of your command
// definition to ensure usability.
func NewCommand(name string, shortDescription string) *Command {
	cmd := &Command{
		Name:             name,
		ShortDescription: shortDescription,
		LongDescription:  "",

		commands:  make([]*Command, 0),
		arguments: make([]string, 0),
		flags:     make([]*Flag, 0),

		Run: func(c *Command) error { return nil },

		output: os.Stdout,
	}

	return cmd
}

// Commands returns the list of sub-commands of this command.
func (c *Command) Commands() []*Command {
	return c.commands
}

// Command returns a *Command that represents an sub-command  of this command given a
// numerical index.
func (c *Command) Command(id int) *Command {
	return c.commands[id]
}

// Arguments returns the list of arguments of this command.
func (c *Command) Arguments() []string {
	return c.arguments
}

// Argument returns an string that represents an argument of this command given a
// numerical index.
func (c *Command) Argument(id int) string {
	return c.arguments[id]
}

// Flags returns the list of flags of this command.
func (c *Command) Flags() []*Flag {
	return c.flags
}

// Flag returns an string that represents an argument of this command given a
// numerical index.
func (c *Command) Flag(id int) *Flag {
	return c.flags[id]
}

// Output retuns the destination for usage and error messages of this command.
//
// By default a Command uses os.Stderr as output.
func (c *Command) Output() io.Writer {
	return c.output
}

// SetOutput sets the destination for usage and error messages.
//
// If output is nil, os.Stderr is used.
func (c *Command) SetOutput(output io.Writer) {
	c.output = output
}

// AddCommand adds a subCommand to this Command.
func (c *Command) AddCommand(cmd *Command) {
	c.commands = append(c.commands, cmd)
}

// parseCommands parses a list of string arguments and builds a list of
// Commands from it.
func (c *Command) parseCommands(args []string) (commands []*Command, err error) {
	return commands, nil
}

// parseArguments parses a list of string arguments and builds a list of
// Arguments from it.
func (c *Command) parseArguments(args []string) (arguments []string, err error) {
	return arguments, nil
}

// parseArguments parses a list of string arguments and builds a list of
// Arguments from it.
func (c *Command) parseFlags(args []string) (flags []*Flag, err error) {
	return flags, nil
}

// Execute executes the command.
//
// Execute uses the args (os.Args[1:] by default) and run through the command
// tree finding appropriate matches for commands and then corresponding flags.
func (c *Command) Execute() error {
	osArgs := os.Args[1:]

	commands, err := c.parseCommands(osArgs)
	if err != nil {
		return err
	}
	c.commands = commands

	arguments, err := c.parseArguments(osArgs)
	if err != nil {
		return err
	}
	c.arguments = arguments

	flags, err := c.parseFlags(osArgs)
	if err != nil {
		return err
	}
	c.flags = flags

	err = c.Run(c)

	return err
}

// Usage puts out the usage for the command.
//
// It is used when a user provides invalid input or when the flag -h or -help
// is attached in the input.
func (c *Command) Usage() {
	templateData := struct {
		Name            string
		LongDescription string
		Commands        []struct {
			Name             string
			ShortDescription string
		}
		Flags []struct {
			ShortName   string
			LongName    string
			Description string
		}
	}{
		Name:            c.Name,
		LongDescription: c.LongDescription,
	}

	for _, subCommand := range c.commands {
		subc := struct {
			Name             string
			ShortDescription string
		}{
			subCommand.Name,
			subCommand.ShortDescription,
		}
		templateData.Commands = append(templateData.Commands, subc)
	}

	for _, flag := range c.flags {
		subf := struct {
			ShortName   string
			LongName    string
			Description string
		}{
			flag.ShortName,
			flag.LongName,
			flag.Description,
		}
		templateData.Flags = append(templateData.Flags, subf)
	}

	t := template.Must(template.New("usageTemplate").Parse(UsageTemplate))
	_ = t.Execute(os.Stderr, templateData)
}
