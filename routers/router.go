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

package routers

import (
	"github.com/beego/beego"

	"github.com/casbin/casvisor/controllers"
)

func init() {
	initAPI()
}

func initAPI() {
	ns :=
		beego.NewNamespace("/api",
			beego.NSInclude(
				&controllers.ApiController{},
			),
		)
	beego.AddNamespace(ns)

	beego.Router("/api/signin", &controllers.ApiController{}, "POST:Signin")
	beego.Router("/api/signout", &controllers.ApiController{}, "POST:Signout")
	beego.Router("/api/get-account", &controllers.ApiController{}, "GET:GetAccount")

	beego.Router("/api/get-global-datasets", &controllers.ApiController{}, "GET:GetGlobalDatasets")
	beego.Router("/api/get-datasets", &controllers.ApiController{}, "GET:GetDatasets")
	beego.Router("/api/get-dataset", &controllers.ApiController{}, "GET:GetDataset")
	beego.Router("/api/update-dataset", &controllers.ApiController{}, "POST:UpdateDataset")
	beego.Router("/api/add-dataset", &controllers.ApiController{}, "POST:AddDataset")
	beego.Router("/api/delete-dataset", &controllers.ApiController{}, "POST:DeleteDataset")
}
