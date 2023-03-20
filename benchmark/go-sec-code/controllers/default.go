package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "index.tpl"
}

func (c *MainController) Post() {
	foo := c.GetString("foo")
	c.Ctx.ResponseWriter.Write([]byte(foo))
}
