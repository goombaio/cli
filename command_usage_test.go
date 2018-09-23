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
	"github.com/goombaio/log"
)

func TestCommand_Usage(t *testing.T) {
	os.Args = []string{"programName"}

	rootCommand := cli.NewCommand("programName", "rootCommand Description")
	rootCommand.LongDescription = "rootCommand Long Description"
	rootCommand.Run = func(c *cli.Command) error {
		c.Usage()

		return nil
	}
	rootCommand.SetLogger(log.NewFmtLogger(os.Stderr))

	buf := new(bytes.Buffer)
	rootCommand.SetOutput(buf)

	subCommand1 := cli.NewCommand("subCommand1", "subCommand1 Description")
	subCommand1.LongDescription = "subCommand1 Long Description"
	subCommand1.Run = func(c *cli.Command) error {
		c.Usage()

		return nil
	}
	rootCommand.AddCommand(subCommand1)

	err := rootCommand.Execute()
	if err != nil {
		t.Fatalf("Expected no error but got %s", err)
	}

	expected := fmt.Sprintf("usage: %s [-help] <command> [args]\n", rootCommand.Name)
	expected += fmt.Sprintf("\n")
	expected += fmt.Sprintf("  %s\n", rootCommand.LongDescription)
	expected += fmt.Sprintf("\n")
	expected += fmt.Sprintf("Commands:\n")
	for _, command := range rootCommand.Commands() {
		expected += fmt.Sprintf("  %s	%s\n", command.Name, command.ShortDescription)
	}
	expected += fmt.Sprintf("\n")
	expected += fmt.Sprintf("Flags:\n")
	for _, flag := range rootCommand.Flags() {
		expected += fmt.Sprintf("  %s, %s	%s\n", flag.ShortName, flag.LongName, flag.Description)
	}
	expected += fmt.Sprintf("\n")
	expected += fmt.Sprintf("Use %s [command] -help for more information about a command.\n", rootCommand.Name)
	if buf.String() != expected {
		t.Fatalf("Expected %q but got %q", expected, buf.String())
	}
}
