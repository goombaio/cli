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
	"io/ioutil"
	"os"
	"testing"

	"github.com/goombaio/cli"
)

func TestCommand(t *testing.T) {
	programName := "programName"
	programShortDescription := "root Command Description"

	rootCommand := cli.NewCommand(programName, programShortDescription)

	os.Args = []string{programName}

	err := rootCommand.Execute()
	if err != nil {
		t.Fatalf("Expected no error but got %s", err)
	}
}

func TestCommand_Name(t *testing.T) {
	programName := "programName"
	rootCommand := cli.NewCommand(programName, "root Command Description")

	if rootCommand.Name() != programName {
		t.Fatalf("Expected command name %s but got %s", programName, rootCommand.Name())
	}

	os.Args = []string{programName}

	err := rootCommand.Execute()
	if err != nil {
		t.Fatalf("Expected no error but got %s", err)
	}
}

func TestCommand_Args_noargs(t *testing.T) {
	programName := "programName"
	programShortDescription := "root Command Description"

	rootCommand := cli.NewCommand(programName, programShortDescription)

	os.Args = []string{programName}

	err := rootCommand.Execute()
	if err != nil {
		t.Fatalf("Expected no error but got %s", err)
	}

	if len(rootCommand.Args()) != 0 {
		t.Fatalf("Expected 0 args but got %d", len(rootCommand.Args()))
	}
}

func TestCommand_Args_args(t *testing.T) {
	programName := "programName"
	programShortDescription := "root Command Description"
	arg1 := "arg1"
	arg2 := "arg2"

	rootCommand := cli.NewCommand(programName, programShortDescription)

	os.Args = []string{programName, arg1, arg2}

	err := rootCommand.Execute()
	if err != nil {
		t.Fatalf("Expected no error but got %s", err)
	}

	if len(rootCommand.Args()) != 2 {
		t.Fatalf("Expected 2 args but got %d", len(rootCommand.Args()))
	}
}

func TestCommand_Arg(t *testing.T) {
	programName := "programName"
	programShortDescription := "root Command Description"
	arg1 := "arg1"
	arg2 := "arg2"

	rootCommand := cli.NewCommand(programName, programShortDescription)

	os.Args = []string{programName, arg1, arg2}

	err := rootCommand.Execute()
	if err != nil {
		t.Fatalf("Expected no error but got %s", err)
	}

	if rootCommand.Arg(0) != arg1 {
		t.Fatalf("Expected arg1 but got %s", rootCommand.Arg(0))
	}

	if rootCommand.Arg(1) != arg2 {
		t.Fatalf("Expected arg2 but got %s", rootCommand.Arg(1))
	}

	if rootCommand.Arg(2) != "" {
		t.Fatalf("Expected blank but got %s", rootCommand.Arg(2))
	}
}

func TestCommand_Output(t *testing.T) {
	programName := "programName"
	programShortDescription := "root Command Description"

	rootCommand := cli.NewCommand(programName, programShortDescription)
	rootCommand.SetOutput(ioutil.Discard)

	output := rootCommand.Output()
	if output != ioutil.Discard {
		t.Fatalf("Expected ioutil.Discard but got %#v", output)
	}
}

func TestCommand_SetOutput(t *testing.T) {
	programName := "programName"
	programShortDescription := "root Command Description"

	rootCommand := cli.NewCommand(programName, programShortDescription)
	rootCommand.Run = func(c *cli.Command) error {
		c.Usage()

		return nil
	}
	buf := new(bytes.Buffer)
	rootCommand.SetOutput(buf)

	os.Args = []string{programName}

	err := rootCommand.Execute()
	if err != nil {
		t.Fatalf("Expected no error but got %s", err)
	}

	expected := fmt.Sprintf("usage: %s [-help] <command> [args]\n", rootCommand.Name())
	expected += fmt.Sprintf("\n")
	expected += fmt.Sprintf("Flags:\n")
	expected += fmt.Sprintf("  -h, -help\tShow help\n")
	if buf.String() != expected {
		t.Fatalf("Expected %q but got %q", expected, buf.String())
	}
}

func TestCommand_AddCommand(t *testing.T) {
	programName := "programName"
	programShortDescription := "root Command Description"

	rootCommand := cli.NewCommand(programName, programShortDescription)
	rootCommand.Run = func(c *cli.Command) error {
		c.Usage()

		return nil
	}
	buf := new(bytes.Buffer)
	rootCommand.SetOutput(buf)

	subCommandName := "subCommand"
	subCommand := cli.NewCommand(subCommandName, "Sub Command Description")

	rootCommand.AddCommand(subCommand)

	os.Args = []string{programName}

	err := rootCommand.Execute()
	if err != nil {
		t.Fatalf("Expected no error but got %s", err)
	}

	expected := fmt.Sprintf("usage: %s [-help] <command> [args]\n", rootCommand.Name())
	expected += fmt.Sprintf("\n")
	expected += fmt.Sprintf("Commands:\n")
	expected += fmt.Sprintf("  subCommand\tSub Command Description\n")
	expected += fmt.Sprintf("\n")
	expected += fmt.Sprintf("Flags:\n")
	expected += fmt.Sprintf("  -h, -help\tShow help\n")
	expected += fmt.Sprintf("\n")
	expected += fmt.Sprintf("Use %s [command] -help for more information about a command.\n", rootCommand.Name())
	if buf.String() != expected {
		t.Fatalf("Expected %q but got %q", expected, buf.String())
	}
}

func TestCommand_Execute(t *testing.T) {
	programName := "programName"
	rootCommand := cli.NewCommand(programName, "root Command Description")
	rootCommand.SetOutput(ioutil.Discard)

	os.Args = []string{programName}

	err := rootCommand.Execute()
	if err != nil {
		t.Fatalf("Expected no error but got %s", err)
	}
}

func TestCommand_Execute_noflags_noargs(t *testing.T) {
	programName := "programName"
	programShortDescription := "root Command Description"

	rootCommand := cli.NewCommand(programName, programShortDescription)
	rootCommand.Run = func(c *cli.Command) error {
		c.Usage()

		return nil
	}
	buf := new(bytes.Buffer)
	rootCommand.SetOutput(buf)

	os.Args = []string{programName}

	err := rootCommand.Execute()
	if err != nil {
		t.Fatalf("Expected no error but got %s", err)
	}

	expected := fmt.Sprintf("usage: %s [-help] <command> [args]\n", rootCommand.Name())
	expected += fmt.Sprintf("\n")
	expected += fmt.Sprintf("Flags:\n")
	expected += fmt.Sprintf("  -h, -help\tShow help\n")
	if buf.String() != expected {
		t.Fatalf("Expected %q but got %q", expected, buf.String())
	}
}

func TestCommand_Execute_noflags_args(t *testing.T) {
	programName := "programName"
	programShortDescription := "root Command Description"
	arg1 := "command"

	rootCommand := cli.NewCommand(programName, programShortDescription)

	os.Args = []string{programName, arg1}

	err := rootCommand.Execute()
	if err != nil {
		t.Fatalf("Expected no error but got %s", err)
	}

	if len(rootCommand.Args()) != 1 {
		t.Fatalf("Expected 1 args but got %d", len(rootCommand.Args()))
	}

	if rootCommand.Arg(0) != arg1 {
		t.Fatalf("Expected command but got %s", rootCommand.Arg(0))
	}
}

func TestCommand_Execute_subcommand(t *testing.T) {
	programName := "programName"
	programShortDescription := "root Command Description"

	rootCommand := cli.NewCommand(programName, programShortDescription)

	subCommandName := "subCommand"
	subCommand := cli.NewCommand(subCommandName, "subCommand description")
	subCommand.LongDescription = "subCommand is the long description about this command"
	subCommand.Run = func(c *cli.Command) error {
		c.Usage()

		return nil
	}
	buf := new(bytes.Buffer)
	subCommand.SetOutput(buf)

	rootCommand.AddCommand(subCommand)

	os.Args = []string{programName, "subCommand"}

	err := rootCommand.Execute()
	if err != nil {
		t.Fatalf("Expected no error but got %s", err)
	}

	expected := fmt.Sprintf("usage: %s [-help] <command> [args]\n", subCommand.Name())
	expected += fmt.Sprintf("\n")
	expected += fmt.Sprintf("  %s is the long description about this command.\n", subCommand.Name())
	expected += fmt.Sprintf("\n")
	expected += fmt.Sprintf("Flags:\n")
	expected += fmt.Sprintf("  -h, -help\tShow help\n")
	if buf.String() != expected {
		t.Fatalf("Expected %q but got %q", expected, buf.String())
	}
}
