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

/*
Package cli implements a commander providing a command-subcommand interface to
create powerful modern CLI tools.

In addition to providing an interface, cli also provides a controller to
organize your application code.

Example

	package cli_test

	import (
		"fmt"
		"os"

		"github.com/goombaio/cli"
	)

	func ExampleCommand() {
		programName := "programName"
		rootCommand := cli.NewCommand(programName, "rootCommand Short Description")
		rootCommand.Run = func(c *cli.Command) error {
			c.Usage()

			return nil
		}
		rootCommand.SetOutput(os.Stdout)

		os.Args = []string{programName}

		err := rootCommand.Execute()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		// Output:
		// usage: programName [-help] <command> [args]
		//
		// Flags:
		//   -h, -help	Show help
	}
*/
package cli
