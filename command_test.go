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
	"io/ioutil"
	"testing"

	"github.com/goombaio/cli"
)

// ExampleRoot_Test ...
func TestCommand(t *testing.T) {
	command1 := cli.NewCommand("command1")

	expected := "command1"
	if command1.Name != expected {
		t.Fatalf("Expected command name %s but got %s", expected, command1.Name)
	}
}

func TestCommand_AddSubCommand(t *testing.T) {
	command1 := cli.NewCommand("command1")

	subcommand1 := cli.NewCommand("subcommand_name")
	command1.AddSubCommand(subcommand1)
}

func TestCommand_Run(t *testing.T) {
	command1 := cli.NewCommand("command1")

	command1.SetOutput(ioutil.Discard)

	err := command1.Run()
	if err != nil {
		t.Fatalf("Expected no error but got %s", err)
	}
}
