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
	"os"
	"testing"

	"github.com/goombaio/cli"
)

func TestCommand_ParseCommands(t *testing.T) {
	os.Args = []string{"programName"}

	rootCommand := cli.NewCommand("programName", "rootCommand Description")
	rootCommand.LongDescription = "rootCommand Long Description"

	subCommand1 := cli.NewCommand("subCommand1", "subCommand1 Description")
	subCommand1.LongDescription = "subCommand1 Long Description"
	rootCommand.AddCommand(subCommand1)

	_ = rootCommand.ParseCommands(os.Args)
}

func TestCommand_ParseCommands_withArguments(t *testing.T) {
	os.Args = []string{"programName", "argument1"}

	rootCommand := cli.NewCommand("programName", "rootCommand Description")
	rootCommand.LongDescription = "rootCommand Long Description"

	subCommand1 := cli.NewCommand("subCommand1", "subCommand1 Description")
	subCommand1.LongDescription = "subCommand1 Long Description"
	rootCommand.AddCommand(subCommand1)

	_ = rootCommand.ParseCommands(os.Args)
}

func TestCommand_ParseCommands_withArguments_withSubCommands(t *testing.T) {
	os.Args = []string{"programName", "argument1", "subCommand1", "argument2"}

	rootCommand := cli.NewCommand("programName", "rootCommand Description")
	rootCommand.LongDescription = "rootCommand Long Description"

	subCommand1 := cli.NewCommand("subCommand1", "subCommand1 Description")
	subCommand1.LongDescription = "subCommand1 Long Description"
	rootCommand.AddCommand(subCommand1)

	_ = rootCommand.ParseCommands(os.Args)
}

func TestCommand_ParseCommands_withFlags(t *testing.T) {
	os.Args = []string{"programName", "-flag1", "-flag2"}

	rootCommand := cli.NewCommand("programName", "rootCommand Description")
	rootCommand.LongDescription = "rootCommand Long Description"

	subCommand1 := cli.NewCommand("subCommand1", "subCommand1 Description")
	subCommand1.LongDescription = "subCommand1 Long Description"
	rootCommand.AddCommand(subCommand1)

	_ = rootCommand.ParseCommands(os.Args)
}

func TestCommand_ParseCommands_withFlags_withSubCommands(t *testing.T) {
	os.Args = []string{"programName", "-flag1", "subCommand1", "-flag2"}

	rootCommand := cli.NewCommand("programName", "rootCommand Description")
	rootCommand.LongDescription = "rootCommand Long Description"

	subCommand1 := cli.NewCommand("subCommand1", "subCommand1 Description")
	subCommand1.LongDescription = "subCommand1 Long Description"
	rootCommand.AddCommand(subCommand1)

	_ = rootCommand.ParseCommands(os.Args)
}

func TestCommand_ParseFlags_shortFlag(t *testing.T) {
	os.Args = []string{"programName", "-h"}

	rootCommand := cli.NewCommand("programName", "rootCommand Description")
	rootCommand.LongDescription = "rootCommand Long Description"

	subCommand1 := cli.NewCommand("subCommand1", "subCommand1 Description")
	subCommand1.LongDescription = "subCommand1 Long Description"
	rootCommand.AddCommand(subCommand1)

	cmd := rootCommand.ParseCommands(os.Args)
	_ = cmd.ParseFlags(os.Args)
}

func TestCommand_ParseFlags_longFlag(t *testing.T) {
	os.Args = []string{"programName", "-help"}

	rootCommand := cli.NewCommand("programName", "rootCommand Description")
	rootCommand.LongDescription = "rootCommand Long Description"

	subCommand1 := cli.NewCommand("subCommand1", "subCommand1 Description")
	subCommand1.LongDescription = "subCommand1 Long Description"
	rootCommand.AddCommand(subCommand1)

	cmd := rootCommand.ParseCommands(os.Args)
	_ = cmd.ParseFlags(os.Args)
}
