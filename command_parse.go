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

// ParseCommands ...
func (c *Command) ParseCommands(args []string) *Command {
	cmd := c

	// Do not parse if there is no subcommands
	if len(cmd.commands) == 0 {
		return cmd
	}

	offsetArgs := 1

	for _, arg := range args {
		candidate := ""

		if !IsFlag(arg) {
			candidate = arg
		} else {
			offsetArgs++
		}

		if candidate == "" {
			continue
		}

		for _, command := range cmd.Commands() {
			if command.Name == candidate {
				command.arguments = cmd.arguments[offsetArgs:]
				cmd = command
			}
		}

		offsetArgs++
	}

	return cmd
}

// ParseFlags ...
func (c *Command) ParseFlags(args []string) ([]string, error) {
	var pflags []string
	for _, arg := range args {
		switch {
		// A flag without a value, or with an `=` separated value
		case IsFlag(arg):
			pflags = append(pflags, arg)
			continue
		}
	}
	return pflags, nil
}
