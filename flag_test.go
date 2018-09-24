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
	"testing"

	"github.com/goombaio/cli"
)

func TestFlag(t *testing.T) {
	_ = &cli.Flag{
		ShortName:   "-f",
		LongName:    "-flag",
		Description: "Flag Description",
		Value:       "false",
		Parsed:      false,
	}
}

func TestFlag_IsFlag_short_novalue(t *testing.T) {
	isFlag := cli.IsFlag("-h")
	if !isFlag {
		t.Fatalf("-h expected to be a flag but it does not")
	}
}

func TestFlag_IsFlag_long_novalue(t *testing.T) {
	isFlag := cli.IsFlag("-help")
	if !isFlag {
		t.Fatalf("-help expected to be a flag but it does not")
	}
}

func TestFlag_IsFlag_short_value(t *testing.T) {
	isFlag := cli.IsFlag("-f=bar")
	if !isFlag {
		t.Fatalf("-f=bar expected to be a flag but it does not")
	}
}

func TestFlag_IsFlag_long_value(t *testing.T) {
	isFlag := cli.IsFlag("-foo=bar")
	if !isFlag {
		t.Fatalf("-foo=bar expected to be a flag but it does not")
	}
}
