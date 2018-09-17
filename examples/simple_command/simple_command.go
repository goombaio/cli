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

package main

import (
	"fmt"
	"os"

	"github.com/goombaio/cli"
)

func main() {
	rootCommand := cli.NewCommand("program", "rootCommand Short Description")
	rootCommand.Run = func(c *cli.Command) error {
		c.Usage()

		return nil
	}

	subCommand1 := cli.NewCommand("subCommand1", "subCommand1 Short Description")
	subCommand1.LongDescription = "subCommand1 Long Description"
	subCommand1.Run = func(c *cli.Command) error {
		c.Usage()

		return nil
	}
	rootCommand.AddCommand(subCommand1)

	subCommand2 := cli.NewCommand("subCommand2", "subCommand2 Short Description")
	subCommand2.LongDescription = "subCommand2 Long Description"
	subCommand2.Run = func(c *cli.Command) error {
		c.Usage()

		return nil
	}
	rootCommand.AddCommand(subCommand2)

	err := rootCommand.Execute()
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(2)
	}
}
