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
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/goombaio/cli"
)

func TestCommand(t *testing.T) {
	os.Args = []string{"programName"}

	rootCommand := cli.NewCommand("programName", "rootCommand Description")
	rootCommand.LongDescription = "rootCommand Long Description"

	err := rootCommand.Execute()
	if err != nil {
		t.Fatalf("Expected no error but got %s", err)
	}
}

func TestCommand_withoutConstructor(t *testing.T) {
	os.Args = []string{"programName"}

	rootCommand := &cli.Command{
		Name:             "programName",
		ShortDescription: "rootCommand Description",
		LongDescription:  "rootCommand Long Description",
		Run: func(c *cli.Command) error {
			c.Usage()

			return nil
		},
	}

	err := cli.Execute(rootCommand)
	if err != nil {
		t.Fatalf("Expected no error but got %s", err)
	}
}

func TestCommand_Execute(t *testing.T) {
	os.Args = []string{"programName"}

	rootCommand := cli.NewCommand("programName", "rootCommand Description")
	rootCommand.LongDescription = "rootCommand Long Description"

	err := rootCommand.Execute()
	if err != nil {
		t.Fatalf("Expected no error but got %s", err)
	}
}

func TestCommand_Name(t *testing.T) {
	os.Args = []string{"programName"}

	rootCommand := cli.NewCommand("programName", "rootCommand Description")
	rootCommand.LongDescription = "rootCommand Long Description"

	if rootCommand.Name != "programName" {
		t.Fatalf("Expected %q but got %q", "programName", rootCommand.Name)
	}

	err := rootCommand.Execute()
	if err != nil {
		t.Fatalf("Expected no error but got %s", err)
	}
}

func TestCommand_ShortDescription(t *testing.T) {
	os.Args = []string{"programName"}

	rootCommand := cli.NewCommand("programName", "rootCommand Description")
	rootCommand.LongDescription = "rootCommand Long Description"

	if rootCommand.ShortDescription != "rootCommand Description" {
		t.Fatalf("Expected %q but got %q", "rootCommand Description", rootCommand.ShortDescription)
	}

	err := rootCommand.Execute()
	if err != nil {
		t.Fatalf("Expected no error but got %s", err)
	}
}

func TestCommand_LongDescription(t *testing.T) {
	os.Args = []string{"programName"}

	rootCommand := cli.NewCommand("programName", "rootCommand Description")
	rootCommand.LongDescription = "rootCommand Long Description"

	if rootCommand.LongDescription != "rootCommand Long Description" {
		t.Fatalf("Expected %q but got %q", "rootCommand Long Description", rootCommand.LongDescription)
	}

	err := rootCommand.Execute()
	if err != nil {
		t.Fatalf("Expected no error but got %s", err)
	}
}

func TestCommand_LongDescription_Unset(t *testing.T) {
	os.Args = []string{"programName"}

	rootCommand := cli.NewCommand("programName", "rootCommand Description")

	if rootCommand.LongDescription != "" {
		t.Fatalf("Expected %q but got %q", "", rootCommand.LongDescription)
	}

	err := rootCommand.Execute()
	if err != nil {
		t.Fatalf("Expected no error but got %s", err)
	}
}

func TestCommand_countCommands_countArguments_countFlags(t *testing.T) {
	os.Args = []string{"programName"}

	rootCommand := cli.NewCommand("programName", "rootCommand Description")
	rootCommand.LongDescription = "rootCommand Long Description"

	err := rootCommand.Execute()
	if err != nil {
		t.Fatalf("Expected no error but got %s", err)
	}

	if len(rootCommand.Commands()) != 0 {
		t.Fatalf("Expected 0 sub-commands but got %d", len(rootCommand.Commands()))
	}

	if len(rootCommand.Arguments()) != 0 {
		t.Fatalf("Expected 0 arguments but got %d", len(rootCommand.Arguments()))
	}

	if len(rootCommand.Flags()) != 1 {
		t.Fatalf("Expected 1 flags but got %d", len(rootCommand.Flags()))
	}
}

func TestCommand_Commands(t *testing.T) {
	os.Args = []string{"programName"}

	rootCommand := cli.NewCommand("programName", "rootCommand Description")
	rootCommand.LongDescription = "rootCommand Long Description"

	subCommand1 := cli.NewCommand("subCommand1", "subCommand1 Description")
	subCommand1.LongDescription = "subCommand1 Long Description"
	rootCommand.AddCommand(subCommand1)

	if len(rootCommand.Commands()) != 1 {
		t.Fatalf("Expected 1 sub-commands but got %d", len(rootCommand.Commands()))
	}
}

func TestCommand_Command(t *testing.T) {
	os.Args = []string{"programName"}

	rootCommand := cli.NewCommand("programName", "rootCommand Description")
	rootCommand.LongDescription = "rootCommand Long Description"

	subCommand1 := cli.NewCommand("subCommand1", "subCommand1 Description")
	subCommand1.LongDescription = "subCommand1 Long Description"
	rootCommand.AddCommand(subCommand1)

	if rootCommand.Command(0).Name != "subCommand1" {
		t.Fatalf("Expected subCommand1 but got %s", rootCommand.Command(0).Name)
	}

	if rootCommand.Command(0).ShortDescription != "subCommand1 Description" {
		t.Fatalf("Expected subCommand1 Description but got %s", rootCommand.Command(0).ShortDescription)
	}

	if rootCommand.Command(0).LongDescription != "subCommand1 Long Description" {
		t.Fatalf("Expected subCommand1 Long Description but got %s", rootCommand.Command(0).LongDescription)
	}
}

func TestCommand_Arguments(t *testing.T) {
	os.Args = []string{"programName", "argument1"}

	rootCommand := cli.NewCommand("programName", "rootCommand Description")
	rootCommand.LongDescription = "rootCommand Long Description"

	subCommand1 := cli.NewCommand("subCommand1", "subCommand1 Description")
	subCommand1.LongDescription = "subCommand1 Long Description"
	rootCommand.AddCommand(subCommand1)

	if len(rootCommand.Arguments()) != 1 {
		t.Fatalf("Expected 1 arguments but got %d", len(rootCommand.Arguments()))
	}
}

func TestCommand_Argument(t *testing.T) {
	os.Args = []string{"programName", "argument1"}

	rootCommand := cli.NewCommand("programName", "rootCommand Description")
	rootCommand.LongDescription = "rootCommand Long Description"

	subCommand1 := cli.NewCommand("subCommand1", "subCommand1 Description")
	subCommand1.LongDescription = "subCommand1 Long Description"
	rootCommand.AddCommand(subCommand1)

	if rootCommand.Argument(0) != "argument1" {
		t.Fatalf("Expected argument1 arguments but got %d", len(rootCommand.Argument(0)))
	}
}

func TestCommand_Flags(t *testing.T) {
	os.Args = []string{"programName"}

	rootCommand := cli.NewCommand("programName", "rootCommand Description")
	rootCommand.LongDescription = "rootCommand Long Description"

	subCommand1 := cli.NewCommand("subCommand1", "subCommand1 Description")
	subCommand1.LongDescription = "subCommand1 Long Description"
	rootCommand.AddCommand(subCommand1)

	if len(rootCommand.Flags()) != 1 {
		t.Fatalf("Expected 1 flags but got %d", len(rootCommand.Flags()))
	}
}

func TestCommand_Flag(t *testing.T) {
	os.Args = []string{"programName"}

	rootCommand := cli.NewCommand("programName", "rootCommand Description")
	rootCommand.LongDescription = "rootCommand Long Description"

	subCommand1 := cli.NewCommand("subCommand1", "subCommand1 Description")
	subCommand1.LongDescription = "subCommand1 Long Description"
	rootCommand.AddCommand(subCommand1)

	if rootCommand.Flag(0).ShortName != "-h" {
		t.Fatalf("Expected -h but got %s", rootCommand.Flag(0).ShortName)
	}

	if rootCommand.Flag(0).LongName != "-help" {
		t.Fatalf("Expected -help but got %s", rootCommand.Flag(0).LongName)
	}

	if rootCommand.Flag(0).Description != "Show help message" {
		t.Fatalf("Expected Show help message but got %s", rootCommand.Flag(0).Description)
	}

	if rootCommand.Flag(0).Value != "false" {
		t.Fatalf("Expected false but got %s", rootCommand.Flag(0).Value)
	}
}

func TestCommand_SetOutput(t *testing.T) {
	os.Args = []string{"programName"}

	rootCommand := cli.NewCommand("programName", "rootCommand Description")
	rootCommand.LongDescription = "rootCommand Long Description"
	rootCommand.Run = func(c *cli.Command) error {
		fmt.Fprintf(c.Output(), "Run %s\n", c.Name)

		return nil
	}
	buf := new(bytes.Buffer)
	rootCommand.SetOutput(buf)

	err := rootCommand.Execute()
	if err != nil {
		t.Fatalf("Expected no error but got %s", err)
	}

	expected := fmt.Sprintf("Run programName\n")
	if buf.String() != expected {
		t.Fatalf("Expected %q but got %q", expected, buf.String())
	}
}

func TestCommand_Output(t *testing.T) {
	os.Args = []string{"programName"}

	rootCommand := cli.NewCommand("programName", "rootCommand Description")
	rootCommand.LongDescription = "rootCommand Long Description"
	rootCommand.Run = func(c *cli.Command) error {
		fmt.Fprintf(c.Output(), "Run %s\n", c.Name)

		return nil
	}
	buf := new(bytes.Buffer)
	rootCommand.SetOutput(buf)

	output := rootCommand.Output()

	if output != buf {
		t.Fatalf("Expected %#v but got %#v", buf, output)
	}

	err := rootCommand.Execute()
	if err != nil {
		t.Fatalf("Expected no error but got %s", err)
	}

	expected := fmt.Sprintf("Run programName\n")
	if buf.String() != expected {
		t.Fatalf("Expected %q but got %q", expected, buf.String())
	}
}
