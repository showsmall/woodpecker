// Copyright 2021 Woodpecker Authors
// Copyright 2018 Drone.IO Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package model

import (
	"errors"
)

var (
	errEnvironNameInvalid  = errors.New("invalid Environment Variable Name")
	errEnvironValueInvalid = errors.New("invalid Environment Variable Value")
)

// EnvironStore persists environment information to storage.
type EnvironStore interface {
	EnvironList(*Repo) ([]*Environ, error)
}

// Environ represents an environment variable.
type Environ struct {
	Name  string `json:"name"`
	Value string `json:"value,omitempty"`
}

// Validate validates the required fields and formats.
func (e *Environ) Validate() error {
	switch {
	case len(e.Name) == 0:
		return errEnvironNameInvalid
	case len(e.Value) == 0:
		return errEnvironValueInvalid
	default:
		return nil
	}
}

// Copy makes a copy of the environment variable without the value.
func (e *Environ) Copy() *Environ {
	return &Environ{
		Name: e.Name,
	}
}
