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

// Execute executes the root command.
//
// Execute uses the command arguments and run through the command tree finding
// appropriate matches for commands and then corresponding flags.
func Execute(c *Command) error {
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

	// In other case run the command action.
	err := cmd.Run(cmd)

	return err
}
