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
	"html/template"
	"io"
)

const (
	// UsageTemplate ...
	UsageTemplate = `{{if .LongDescription}}
{{.LongDescription}}

{{end}}usage: {{.ProgramName}} [-help] <command> [args]
{{if .Commands}}
Commands:
{{range .Commands}}  {{.Name}}	{{.Description}}{{end}}
{{end}}
Flags:
  -h, -help	Show help

Use {{.ProgramName}} [command] -help for more information about a command
`
)

// Command ...
type Command struct {
	ShortDescription string
	LongDescription  string

	commands []*Command
	flags    *flag.FlagSet

	Run func() error
}

// NewCommand ...
func NewCommand(name string, shortDescription string) *Command {
	cmd := &Command{
		ShortDescription: shortDescription,
		LongDescription:  "",

		commands: make([]*Command, 0),
		flags:    flag.NewFlagSet(name, flag.ContinueOnError),

		Run: func() error { return nil },
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
	err := c.flags.Parse(flag.Args())
	if err != nil {
		return err
	}

	err = c.Run()

	return err
}

// Usage ...
func (c *Command) Usage() {
	templateData := struct {
		ProgramName     string
		LongDescription string
		Commands        []struct {
			Name        string
			Description string
		}
	}{
		ProgramName:     c.flags.Name(),
		LongDescription: c.LongDescription,
	}
	for _, command := range c.commands {
		c := struct {
			Name        string
			Description string
		}{
			command.Name(),
			command.ShortDescription,
		}
		templateData.Commands = append(templateData.Commands, c)
	}
	t := template.Must(template.New("usage").Parse(UsageTemplate))
	_ = t.Execute(c.flags.Output(), templateData)
}
