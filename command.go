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
)

// Command ...
type Command struct {
	Name     string
	Commands []*Command
	output   io.Writer
}

// NewCommand ...
func NewCommand(name string) *Command {
	cmd := &Command{
		Name:     name,
		Commands: make([]*Command, 0),
		output:   os.Stderr,
	}

	return cmd
}

// SetOutput ...
func (c *Command) SetOutput(output io.Writer) {
	c.output = output
}

// AddSubCommand ...
func (c *Command) AddSubCommand(cmd *Command) {
	c.Commands = append(c.Commands, cmd)
}

// Run ...
func (c *Command) Run() error {

	return nil
}
