# cli

A flag parse library with support for commands and subcommands.

[![License][License-Image]][License-URL]
[![CircleCI Status][CircleCI-Image]][CircleCI-URL]
[![Coverage Report][Coverage-Image]][Coverage-URL]
[![Go Report Card][GoReportCard-Image]][GoReportCard-URL]
[![CII Best Practices][CII-Image]][CII-URL]
[![GoDoc][GoDoc-Image]][GoDoc-URL]

## Install

```bash
go get github.com/goombaio/cli
```

You can also update an already installed version:

```bash
go get -u github.com/goombaio/cli
```

## Example of use

```go
package main

import (
    "fmt"
    "os"

    "github.com/goombaio/cli"
)

func main() {
    rootCommand := cli.NewCommand("programName")
    rootCommand.Run = func() error {
        if len(rootCommand.Args()) == 0 {
            rootCommand.Usage()
        }

        return nil
    }

    err := rootCommand.Execute()
    if err != nil {
        fmt.Println("ERROR:", err)
        os.Exit(1)
    }
    // Output:
    // usage: programName [-help] <command> <args>
    //
    // Flags:
    //   -h, -help  Show help
    //
    // Use programName [command] -help for more information about a command
}
```

## License

Copyright (c) 2018 Goomba project Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

[License-Image]: https://img.shields.io/badge/License-Apache-blue.svg
[License-URL]: http://opensource.org/licenses/Apache
[CircleCI-Image]: https://circleci.com/gh/goombaio/cli.svg?style=svg
[CircleCI-URL]: https://circleci.com/gh/goombaio/cli
[Coverage-Image]: https://codecov.io/gh/goombaio/cli/branch/master/graph/badge.svg
[Coverage-URL]: https://codecov.io/gh/goombaio/cli
[GoReportCard-Image]: https://goreportcard.com/badge/github.com/goombaio/cli
[GoReportCard-URL]: https://goreportcard.com/report/github.com/goombaio/cli
[CII-Image]: https://bestpractices.coreinfrastructure.org/projects/2224/badge
[CII-URL]: https://bestpractices.coreinfrastructure.org/projects/2224
[GoDoc-Image]: https://godoc.org/github.com/goombaio/cli?status.svg
[GoDoc-URL]: http://godoc.org/github.com/goombaio/cli
