package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type CRLFSafe1Controller struct {
	beego.Controller
}

func (c *CRLFSafe1Controller) Get() {
	header := c.GetString("header")
	c.Ctx.ResponseWriter.Header().Set("header", header)
	c.Ctx.ResponseWriter.Write([]byte(""))
}
