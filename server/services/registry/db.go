// Copyright 2023 Woodpecker Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package registry

import (
	"go.woodpecker-ci.org/woodpecker/v2/server/model"
)

type db struct {
	store model.RegistryStore
}

// New returns a new local registry service.
func NewDB(store model.RegistryStore) Service {
	return &db{store}
}

func (d *db) RegistryFind(repo *model.Repo, name string) (*model.Registry, error) {
	return d.store.RegistryFind(repo, name)
}

func (d *db) RegistryList(repo *model.Repo, p *model.ListOptions) ([]*model.Registry, error) {
	return d.store.RegistryList(repo, p)
}

func (d *db) RegistryCreate(_ *model.Repo, in *model.Registry) error {
	return d.store.RegistryCreate(in)
}

func (d *db) RegistryUpdate(_ *model.Repo, in *model.Registry) error {
	return d.store.RegistryUpdate(in)
}

func (d *db) RegistryDelete(repo *model.Repo, addr string) error {
	return d.store.RegistryDelete(repo, addr)
}
