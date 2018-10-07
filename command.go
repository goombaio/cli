// Copyright © 2018, Goomba project Authors. All rights reserved.
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

	"github.com/goombaio/log"
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

	// logger is the log.Logger being used
	logger log.Logger
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
		arguments: os.Args[1:],
		flags:     make([]*Flag, 0),

		Run: func(c *Command) error { return nil },

		output: os.Stdout,

		logger: log.NewNoopLogger(),
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

// Flag returns a cli.Flag that represents a Flag of this command given a
// numerical index.
func (c *Command) Flag(id int) *Flag {
	return c.flags[id]
}

// FlagName returns a cli.Flag that represents a Flag of this command given a
// string name.
func (c *Command) FlagName(name string) *Flag {
	for _, flag := range c.Flags() {
		if flag.ShortName == name {
			return flag
		}
		if flag.LongName == name {
			return flag
		}
	}

	return nil
}

// Output return the destination for usage and error messages of this command.
//
// By default a Command uses os.Stdout as output.
func (c *Command) Output() io.Writer {
	if c.output == nil {
		c.output = os.Stdout
	}
	return c.output
}

// SetOutput sets the destination for usage and error messages.
func (c *Command) SetOutput(output io.Writer) {
	c.output = output
}

// Logger returns the current log.Logger for this Command.
func (c *Command) Logger() log.Logger {
	return c.logger
}

// SetLogger sets the log.Logger to be used.
func (c *Command) SetLogger(logger log.Logger) {
	c.logger = logger
}

// AddCommand adds a subCommand to this Command.
func (c *Command) AddCommand(cmd *Command) {
	// Setup command default flag set
	c.setupDefaultFlags()

	cmd.SetOutput(c.Output())
	cmd.SetLogger(c.Logger())

	c.commands = append(c.commands, cmd)
}

// AddFlag adds a flag to this Command.
func (c *Command) AddFlag(flag *Flag) {
	c.flags = append(c.flags, flag)
}

// execute executes the command.
//
// Execute uses the command arguments and run through the command tree finding
// appropriate matches for commands and then corresponding flags.
func (c *Command) execute() error {
	// Setup command default flag set
	c.setupDefaultFlags()

	c.SetOutput(c.Output())
	c.SetLogger(c.Logger())

	// Parse commands ans subcommands from the cli, routing to the command it
	// Will be selected for execution.
	cmd := c.ParseCommands(c.Arguments())

	// Parses flags and arguments for the selected command for execution.
	cmd = cmd.ParseFlags(c.Arguments())

	// If the special flags '-h', or '-help' are present on the current
	// parsed flags execute the Usage() method for the command.
	for _, flag := range cmd.Flags() {
		if flag.ShortName == "-h" || flag.LongName == "-help" {
			if flag.Parsed {
				cmd.Usage()
				return nil
			}
		}
	}

	// Run the command action if it is runnable.
	if cmd.Run != nil {
		err := cmd.Run(cmd)
		if err != nil {
			return err
		}
	}

	return nil
}

// setupDefaultFlags ...
func (c *Command) setupDefaultFlags() {
	// help Flag
	helpFlag := &Flag{
		ShortName:   "-h",
		LongName:    "-help",
		Description: "Show help message",
		Value:       "false",
	}
	c.flags = append(c.flags, helpFlag)
}
