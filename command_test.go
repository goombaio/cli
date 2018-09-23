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
	programName := "programName"
	programShortDescription := "rootCommand Description"

	rootCommand := cli.NewCommand(programName, programShortDescription)

	os.Args = []string{programName}

	err := rootCommand.Execute()
	if err != nil {
		t.Fatalf("Expected no error but got %s", err)
	}
}

func TestCommand_Execute(t *testing.T) {
	programName := "programName"
	programShortDescription := "rootCommand Description"
	programLongDescription := "rootCommand Long Description"
	rootCommand := cli.NewCommand(programName, programShortDescription)
	rootCommand.LongDescription = programLongDescription

	os.Args = []string{programName}

	err := rootCommand.Execute()
	if err != nil {
		t.Fatalf("Expected no error but got %s", err)
	}
}

func TestCommand_Name(t *testing.T) {
	programName := "programName"
	programShortDescription := "rootCommand Description"
	programLongDescription := "rootCommand Long Description"
	rootCommand := cli.NewCommand(programName, programShortDescription)
	rootCommand.LongDescription = programLongDescription

	if rootCommand.Name != programName {
		t.Fatalf("Expected %q but got %q", programName, rootCommand.Name)
	}

	os.Args = []string{programName}

	err := rootCommand.Execute()
	if err != nil {
		t.Fatalf("Expected no error but got %s", err)
	}
}

func TestCommand_ShortDescription(t *testing.T) {
	programName := "programName"
	programShortDescription := "rootCommand Description"
	programLongDescription := "rootCommand Long Description"
	rootCommand := cli.NewCommand(programName, programShortDescription)
	rootCommand.LongDescription = programLongDescription

	if rootCommand.ShortDescription != programShortDescription {
		t.Fatalf("Expected %q but got %q", rootCommand.ShortDescription, programShortDescription)
	}

	os.Args = []string{programName}

	err := rootCommand.Execute()
	if err != nil {
		t.Fatalf("Expected no error but got %s", err)
	}
}

func TestCommand_LongDescription(t *testing.T) {
	programName := "programName"
	programShortDescription := "rootCommand Description"
	programLongDescription := "rootCommand Long Description"
	rootCommand := cli.NewCommand(programName, programShortDescription)
	rootCommand.LongDescription = programLongDescription

	if rootCommand.LongDescription != programLongDescription {
		t.Fatalf("Expected %q but got %q", rootCommand.LongDescription, programLongDescription)
	}

	os.Args = []string{programName}

	err := rootCommand.Execute()
	if err != nil {
		t.Fatalf("Expected no error but got %s", err)
	}
}

func TestCommand_LongDescription_Unset(t *testing.T) {
	programName := "programName"
	programShortDescription := "rootCommand Description"
	rootCommand := cli.NewCommand(programName, programShortDescription)

	if rootCommand.LongDescription != "" {
		t.Fatalf("Expected %q but got %q", rootCommand.LongDescription, "")
	}

	os.Args = []string{programName}

	err := rootCommand.Execute()
	if err != nil {
		t.Fatalf("Expected no error but got %s", err)
	}
}

func TestCommand_countCommands_countArguments_countFlags(t *testing.T) {
	programName := "programName"
	programShortDescription := "rootCommand Description"
	programLongDescription := "rootCommand Long Description"
	rootCommand := cli.NewCommand(programName, programShortDescription)
	rootCommand.LongDescription = programLongDescription

	os.Args = []string{programName}

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

	if len(rootCommand.Flags()) != 0 {
		t.Fatalf("Expected 0 sub-commands but got %d", len(rootCommand.Flags()))
	}
}

func TestCommand_Output(t *testing.T) {
	programName := "programName"
	programShortDescription := "rootCommand Description"
	programLongDescription := "rootCommand Long Description"
	rootCommand := cli.NewCommand(programName, programShortDescription)
	rootCommand.LongDescription = programLongDescription

	output := rootCommand.Output()
	if output != os.Stdout {
		t.Fatalf("Expected ioutil.Discard but got %#v", output)
	}
}

func TestCommand_SetOutput(t *testing.T) {
	programName := "programName"
	programShortDescription := "rootCommand Description"
	programLongDescription := "rootCommand Long Description"
	rootCommand := cli.NewCommand(programName, programShortDescription)
	rootCommand.LongDescription = programLongDescription
	rootCommand.Run = func(c *cli.Command) error {
		fmt.Fprintf(c.Output(), "Run %s\n", c.Name)

		return nil
	}
	buf := new(bytes.Buffer)
	rootCommand.SetOutput(buf)

	os.Args = []string{programName}

	err := rootCommand.Execute()
	if err != nil {
		t.Fatalf("Expected no error but got %s", err)
	}

	expected := fmt.Sprintf("Run programName\n")
	if buf.String() != expected {
		t.Fatalf("Expected %q but got %q", expected, buf.String())
	}
}
