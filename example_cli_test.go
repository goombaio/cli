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

package cli_test

import (
	"fmt"
	"os"

	"github.com/goombaio/cli"
	"github.com/goombaio/log"
)

func ExampleCommand() {
	os.Args = []string{"programName"}

	rootCommand := cli.NewCommand("programName", "rootCommand Short Description")
	rootCommand.LongDescription = "rootCommand Long Description"
	rootCommand.Run = func(c *cli.Command) error {
		fmt.Fprintf(c.Output(), "Run %s\n", c.Name)

		return nil
	}
	rootCommand.SetLogger(log.NewFmtLogger(os.Stderr))

	err := rootCommand.Execute()
	if err != nil {
		_ = rootCommand.Logger().Log("ERROR:", err)
		os.Exit(1)
	}
	// Output:
	// Run programName
}

func ExampleCommand_Usage() {
	os.Args = []string{"programName", "-help"}

	rootCommand := cli.NewCommand("programName", "rootCommand Short Description")
	rootCommand.LongDescription = "rootCommand Long Description"
	rootCommand.Run = func(c *cli.Command) error {
		fmt.Fprintf(c.Output(), "Run %s\n", c.Name)

		return nil
	}
	rootCommand.SetLogger(log.NewFmtLogger(os.Stderr))

	err := rootCommand.Execute()
	if err != nil {
		_ = rootCommand.Logger().Log("ERROR:", err)
		os.Exit(1)
	}
	// Output:
	// usage: programName [-help] <command> [args]
	//
	//   rootCommand Long Description
	//
	// Flags:
	//   -h, -help	Show help message
	//
	// Use programName [command] -help for more information about a command.
}

func ExampleCommand_subCommand() {
	os.Args = []string{"programName", "subCommand1"}

	rootCommand := cli.NewCommand("programName", "rootCommand Description")
	rootCommand.LongDescription = "rootCommand Long Description"
	rootCommand.Run = func(c *cli.Command) error {
		fmt.Printf("Running %s\n", c.Name)

		return nil
	}
	rootCommand.SetLogger(log.NewFmtLogger(os.Stderr))

	subCommand1 := cli.NewCommand("subCommand1", "subCommand1 Description")
	subCommand1.LongDescription = "subCommand1 Long Description"
	subCommand1.Run = func(c *cli.Command) error {
		fmt.Printf("Running %s\n", c.Name)

		return nil
	}
	rootCommand.AddCommand(subCommand1)

	err := rootCommand.Execute()
	if err != nil {
		_ = rootCommand.Logger().Log("ERROR:", err)
		os.Exit(1)
	}
	// Output:
	// Running subCommand1
}

func ExampleCommand_subCommand_usage() {
	os.Args = []string{"programName", "subCommand1", "-help"}

	rootCommand := cli.NewCommand("programName", "rootCommand Description")
	rootCommand.LongDescription = "rootCommand Long Description"
	rootCommand.Run = func(c *cli.Command) error {
		fmt.Printf("Running %s\n", c.Name)

		return nil
	}
	rootCommand.SetLogger(log.NewFmtLogger(os.Stderr))

	subCommand1 := cli.NewCommand("subCommand1", "subCommand1 Description")
	subCommand1.LongDescription = "subCommand1 Long Description"
	subCommand1.Run = func(c *cli.Command) error {
		fmt.Printf("Running %s\n", c.Name)

		return nil
	}
	rootCommand.AddCommand(subCommand1)

	err := rootCommand.Execute()
	if err != nil {
		_ = rootCommand.Logger().Log("ERROR:", err)
		os.Exit(1)
	}
	// Output:
	// usage: subCommand1 [-help] <command> [args]
	//
	//   subCommand1 Long Description
	//
	// Flags:
	//   -h, -help	Show help message
	//
	// Use subCommand1 [command] -help for more information about a command.
}
