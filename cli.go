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
	"os"
)

// CLI is the main entry point of a CLI application
type CLI struct {
	ProgramName string
	Commands    []*CommandInterface
}

// NewCLI ...
func NewCLI() *CLI {
	cli := &CLI{
		ProgramName: os.Args[0],
		Commands:    make([]*CommandInterface, 0),
	}

	return cli
}

// AddCommand ...
func (c *CLI) AddCommand(cmd *CommandInterface) {
	c.Commands = append(c.Commands, cmd)
}

// Run ...
func (c *CLI) Run() error {

	return nil
}
