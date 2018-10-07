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
	"io/ioutil"
	"os"
	"testing"

	"github.com/goombaio/cli"
)

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

	rootCommand.SetOutput(ioutil.Discard)

	err := cli.Execute(rootCommand)
	if err != nil {
		t.Fatalf("Expected no error but got %s", err)
	}
}

func TestCommand_withoutConstructor_countCommands_countArguments_countFlags(t *testing.T) {
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

	rootCommand.SetOutput(ioutil.Discard)

	err := cli.Execute(rootCommand)
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