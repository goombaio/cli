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

package cli_test

import (
	"fmt"
	"os"

	"github.com/goombaio/cli"
)

func ExampleCommand_useCase1() {
	os.Args = []string{"program", "command", "subcommand", "argument1"}

	rootCommand := cli.NewCommand("program", "program Description")
	rootCommand.SetOutput(os.Stdout)
	rootCommand.LongDescription = "program Long Description"
	rootCommand.Run = func(c *cli.Command) error {
		c.Usage()

		return nil
	}

	commandCommand := cli.NewCommand("command", "command Description")
	commandCommand.LongDescription = "command Long Description"
	commandCommand.Run = func(c *cli.Command) error {
		c.Usage()

		return nil
	}
	rootCommand.AddCommand(commandCommand)

	subcommandCommand := cli.NewCommand("subcommand", "subcommand Description")
	subcommandCommand.LongDescription = "subcommand Long Description"
	subcommandCommand.Run = func(c *cli.Command) error {
		fmt.Fprintf(c.Output(), "Running %s Arguments: %s\n", c.Name, c.Arguments())

		return nil
	}
	commandCommand.AddCommand(subcommandCommand)

	err := cli.Execute(rootCommand)
	if err != nil {
		_ = rootCommand.Logger().Log("ERROR:", err)
		os.Exit(1)
	}
	// Output:
	// Running subcommand Arguments: [argument1]
}
