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
	"os"
)

// CLI is the main entry point of a CLI application
type CLI struct {
	Name     string
	flags    *flag.FlagSet
	args     []string
	commands []*Command
}

// NewCLI ...
func NewCLI() *CLI {
	cli := &CLI{
		Name:     os.Args[0],
		commands: make([]*Command, 0),
		flags:    flag.NewFlagSet(os.Args[0], flag.ExitOnError),
		args:     make([]string, 0),
	}

	flag.Usage = cli.Usage
	cli.flags.Usage = cli.Usage

	flag.Parse()

	cli.args = flag.Args()

	return cli
}

// AddCommand ...
func (c *CLI) AddCommand(cmd *Command) {
	c.commands = append(c.commands, cmd)
}

// Run ...
func (c *CLI) Run() error {
	if len(c.args) == 0 {
		c.Usage()
		return nil
	}

	return nil
}

// SetOutput ...
func (c *CLI) SetOutput(output io.Writer) {
	c.flags.SetOutput(output)
}

// Usage ...
func (c *CLI) Usage() {
	fmt.Fprintf(c.flags.Output(), "usage: %s [-version] [-help] <command> <args>\n", c.Name)
	fmt.Fprintf(c.flags.Output(), "\n")
	fmt.Fprintf(c.flags.Output(), "Flags:\n")
	// fmt.Fprintf(c.output, "  -version\tShow version information\n")
	fmt.Fprintf(c.flags.Output(), "  -help\tShow help\n")
	fmt.Fprintf(c.flags.Output(), "\n")
	fmt.Fprintf(c.flags.Output(), "Use %s [command] -help for more information about a command\n", c.Name)
}

// ShowVersion ...
func (c *CLI) ShowVersion(version string, build string) {
	fmt.Fprintf(c.flags.Output(), "%s version %s build %s\n", c.Name, version, build)
}
