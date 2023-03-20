package controllers

import (
	"go-sec-code/utils"
	"html/template"
	"io/ioutil"

	beego "github.com/beego/beego/v2/server/web"
)

type XSSVuln1Controller struct {
	beego.Controller
}

type XSSVuln2Controller struct {
	beego.Controller
}

type XSSVuln3Controller struct {
	beego.Controller
}

type XSSVuln4Controller struct {
	beego.Controller
}

type XSSSafe1Controller struct {
	beego.Controller
}

type XSSSafe2Controller struct {
	beego.Controller
}

func (c *XSSVuln1Controller) Get() {
	xss := c.GetString("xss", "hello")
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html")
	c.Ctx.ResponseWriter.Write([]byte(xss))
}

func (c *XSSVuln2Controller) Get() {
	xss := c.GetSession("xss")
	if xss == nil {
		xss = "hello"
	}
	c.Data["xss"] = template.HTML(xss.(string))
	c.TplName = "xss.tpl"
}

func (c *XSSVuln3Controller) Get() {
	file, err := ioutil.ReadFile("static/xss/poc.svg")
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(file)
}

func (c *XSSVuln4Controller) Get() {
	file, err := ioutil.ReadFile("static/xss/poc.pdf")
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Header().Set("Content-Security-Policy", `default-src 'self';`)
	c.Ctx.ResponseWriter.Write(file)
}

func (c *XSSVuln2Controller) Post() {
	xss := c.GetString("xss", "hello")
	c.SetSession("xss", xss)
	c.Data["xss"] = template.HTML(xss)
	c.TplName = "xss.tpl"
}

func (c *XSSSafe1Controller) Get() {
	xss := c.GetString("xss", "hello")
	xssFilter := utils.XSSFilter{}
	xss = xssFilter.DoFilter(xss)
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html")
	c.Ctx.ResponseWriter.Write([]byte(xss))
}

func (c *XSSSafe2Controller) Get() {
	file, err := ioutil.ReadFile("static/xss/poc.svg")
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Header().Set("Content-Security-Policy", `default-src 'self';`)
	c.Ctx.ResponseWriter.Write(file)
}
