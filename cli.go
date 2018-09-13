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
	"fmt"
	"io"
	"os"
)

// CLI is the main entry point of a CLI application
type CLI struct {
	ProgramName string
	commands    []*Command
	output      io.Writer
	version     string
	build       string
}

// NewCLI ...
func NewCLI() *CLI {
	cli := &CLI{
		ProgramName: os.Args[0],
		commands:    make([]*Command, 0),
		output:      os.Stderr,
	}

	return cli
}

// AddCommand ...
func (c *CLI) AddCommand(cmd *Command) {
	c.commands = append(c.commands, cmd)
}

// Run ...
func (c *CLI) Run() error {
	/*
		helpPtr := flag.Bool("help", false, "Show help")
		versionPtr := flag.Bool("version", false, "Show version information")

		flag.Parse()

		if *helpPtr {
			c.Usage()
			os.Exit(0)
		}

		if *versionPtr {
			c.ShowVersion(c.version, c.build)
			os.Exit(0)
		}
	*/

	return nil
}

// SetOutput ...
func (c *CLI) SetOutput(output io.Writer) {
	c.output = output
}

// Usage ...
func (c *CLI) Usage() {
	fmt.Fprintf(c.output, "usage: %s [-version] [-help] <command> <args>\n", c.ProgramName)
	fmt.Fprintf(c.output, "\n")
	fmt.Fprintf(c.output, "Flags:\n")
	fmt.Fprintf(c.output, "  --version\tShow version information\n")
	fmt.Fprintf(c.output, "  --help\tShow help\n")
	fmt.Fprintf(c.output, "\n")
	fmt.Fprintf(c.output, "Use %s [command] --help for more information about a command.\n", c.ProgramName)
}

// ShowVersion ...
func (c *CLI) ShowVersion(version string, build string) {
	fmt.Fprintf(c.output, "%s version %s build %s\n", c.ProgramName, version, build)
}
