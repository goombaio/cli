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
	"fmt"
	"io"
)

// Command ...
type Command struct {
	flags            *flag.FlagSet
	commands         []*Command
	ShortDescription string
	Run              func() error
}

// NewCommand ...
func NewCommand(name string, shortDescription string) *Command {
	cmd := &Command{
		commands:         make([]*Command, 0),
		flags:            flag.NewFlagSet(name, flag.ContinueOnError),
		ShortDescription: shortDescription,
		Run:              func() error { return nil },
	}

	return cmd
}

// Name ...
func (c *Command) Name() string {
	return c.flags.Name()
}

// Args ...
func (c *Command) Args() []string {
	return c.flags.Args()
}

// Arg ...
func (c *Command) Arg(id int) string {
	return c.flags.Arg(id)
}

// Output ...
func (c *Command) Output() io.Writer {
	return c.flags.Output()
}

// SetOutput ...
func (c *Command) SetOutput(output io.Writer) {
	c.flags.SetOutput(output)
}

// AddCommand ...
func (c *Command) AddCommand(cmd *Command) {
	c.commands = append(c.commands, cmd)
}

// Execute ...
func (c *Command) Execute() error {
	flag.Usage = c.Usage
	c.flags.Usage = flag.Usage

	flag.Parse()
	c.flags.Parse(flag.Args())

	err := c.Run()

	return err
}

// Usage ...
func (c *Command) Usage() {
	fmt.Fprintf(c.flags.Output(), "usage: %s [-help] <command> [args]\n", c.flags.Name())
	fmt.Fprintf(c.flags.Output(), "\n")
	if len(c.commands) > 0 {
		fmt.Fprintf(c.flags.Output(), "Commands:\n")
		for _, command := range c.commands {
			fmt.Fprintf(c.flags.Output(), "  %s\t%s\n", command.Name(), command.ShortDescription)
		}
		fmt.Fprintf(c.flags.Output(), "\n")
	}
	fmt.Fprintf(c.flags.Output(), "Flags:\n")
	fmt.Fprintf(c.flags.Output(), "  -h, -help\tShow help\n")
	fmt.Fprintf(c.flags.Output(), "\n")
	fmt.Fprintf(c.flags.Output(), "Use %s [command] -help for more information about a command\n", c.flags.Name())
}
