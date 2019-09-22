/*
 * Copyright 2012-2019 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package example

import (
	"github.com/didi/go-spring/spring-web"
	"github.com/didi/go-spring/spring-core"
	"github.com/didi/go-spring/spring-jdbc"
	"github.com/go-spring/go-spring-boot/spring-boot"
)

func init() {
	SpringBoot.RegisterModule(func(springContext SpringCore.SpringContext) {
		springContext.RegisterBean(new(Controller))
	})
}

type Controller struct {
	JdbcTemplate SpringJdbc.JdbcTemplate `autowire:""`
}

func (controller *Controller) InitController(c SpringWeb.WebContainer) {
	c.GET("/test", controller.Test)
}

func (controller *Controller) Test(context *SpringWeb.SpringWebContext) interface{} {

	var r struct {
		Id         string `column:"id"`
		GroupId    string `column:"groupId"`
		DeviceId   string `column:"deviceId"`
		DeviceName string `column:"deviceName"`
	}

	controller.JdbcTemplate.QueryRow(&r, "select * from camera")
	return r
}
