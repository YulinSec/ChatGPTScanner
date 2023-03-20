package controllers

import (
	"io/ioutil"

	beego "github.com/beego/beego/v2/server/web"
)

type FaviconController struct {
	beego.Controller
}

func (c *FaviconController) Get() {
	icon, err := ioutil.ReadFile("favicon.ico")
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(icon)
}
