// Copyright 2023 The casbin Authors. All Rights Reserved.
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

package controllers

import (
	"encoding/json"

	"github.com/casbin/caswire/object"
)

func (c *ApiController) GetGlobalDatasets() {
	c.Data["json"] = object.GetGlobalDatasets()
	c.ServeJSON()
}

func (c *ApiController) GetDatasets() {
	owner := c.Input().Get("owner")

	c.Data["json"] = object.GetDatasets(owner)
	c.ServeJSON()
}

func (c *ApiController) GetDataset() {
	id := c.Input().Get("id")

	c.Data["json"] = object.GetDataset(id)
	c.ServeJSON()
}

func (c *ApiController) UpdateDataset() {
	id := c.Input().Get("id")

	var dataset object.Dataset
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &dataset)
	if err != nil {
		panic(err)
	}

	c.Data["json"] = object.UpdateDataset(id, &dataset)
	c.ServeJSON()
}

func (c *ApiController) AddDataset() {
	var dataset object.Dataset
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &dataset)
	if err != nil {
		panic(err)
	}

	c.Data["json"] = object.AddDataset(&dataset)
	c.ServeJSON()
}

func (c *ApiController) DeleteDataset() {
	var dataset object.Dataset
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &dataset)
	if err != nil {
		panic(err)
	}

	c.Data["json"] = object.DeleteDataset(&dataset)
	c.ServeJSON()
}
