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

// Flag implements a command line flag
type Flag struct {
	ShortName   string
	LongName    string
	Description string
	Value       string
	Parsed      bool
}

// IsFlag checks if an string is a flag or not.
//
// It will be a flag if it has the format:
// * -flag=value
// * --flag
// * -f=value
// * -f
func IsFlag(str string) bool {
	return ((len(str) >= 3 && str[1] == '-') || (len(str) >= 2 && str[0] == '-' && str[1] != '-'))
}
