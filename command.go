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
	"text/template"
)

const (
	// UsageTemplate ...
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

// Command ...
type Command struct {
	ShortDescription string
	LongDescription  string

	commands []*Command
	flags    *flag.FlagSet

	Run func(c *Command) error
}

// NewCommand ...
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

// Usage ...
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
		fmt.Printf("%#v", subCommand)
		cc := struct {
			Name             string
			ShortDescription string
		}{
			"subCommand.Name()",
			subCommand.ShortDescription,
		}
		templateData.Commands = append(templateData.Commands, cc)
	}

	t := template.Must(template.New("usageTemplate").Parse(UsageTemplate))
	_ = t.Execute(c.flags.Output(), templateData)
}
