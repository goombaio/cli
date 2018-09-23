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

package cli

import (
	"text/template"
)

const (
	// UsageTemplate is the template being used to render the Usage for
	// any cli.Command that has a flag -h or-help attached to it.
	UsageTemplate = `usage: {{.Name}} [-help] <command> [args]{{if .LongDescription}}

  {{.LongDescription}}{{end}}
{{if .Commands}}
Commands:
{{range .Commands}}  {{.Name}}	{{.ShortDescription}}
{{end}}{{end}}{{if .Flags}}
Flags:
{{range .Flags}}  {{.ShortName}}, {{.LongName}}	{{.Description}}
{{end}}{{end}}
Use {{.Name}} [command] -help for more information about a command.
`
)

// Usage puts out the usage for the command.
//
// It is used when a user provides invalid input or when the flag -h or -help
// is attached in the input.
func (c *Command) Usage() {
	templateData := struct {
		Name            string
		LongDescription string
		Commands        []struct {
			Name             string
			ShortDescription string
		}
		Flags []struct {
			ShortName   string
			LongName    string
			Description string
		}
	}{
		Name:            c.Name,
		LongDescription: c.LongDescription,
	}

	for _, subCommand := range c.commands {
		subc := struct {
			Name             string
			ShortDescription string
		}{
			subCommand.Name,
			subCommand.ShortDescription,
		}
		templateData.Commands = append(templateData.Commands, subc)
	}

	for _, flag := range c.flags {
		subf := struct {
			ShortName   string
			LongName    string
			Description string
		}{
			flag.ShortName,
			flag.LongName,
			flag.Description,
		}
		templateData.Flags = append(templateData.Flags, subf)
	}

	t := template.Must(template.New("usageTemplate").Parse(UsageTemplate))
	_ = t.Execute(c.Output(), templateData)
}
