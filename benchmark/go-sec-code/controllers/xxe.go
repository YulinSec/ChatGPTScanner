package controllers

import (
	"bytes"
	"fmt"
	"io/ioutil"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/beevik/etree"
	"github.com/lestrrat-go/libxml2/parser"
)

type XXEVuln1Controller struct {
	beego.Controller
}

type XXEVuln2Controller struct {
	beego.Controller
}

type XXESafe1Controller struct {
	beego.Controller
}

func (c *XXEVuln1Controller) Get() {
	file, err := ioutil.ReadFile("static/xml/xxe.xml")
	if err != nil {
		panic(err)
	}
	c.Data["xxe"] = string(file)
	c.TplName = "xxe.tpl"
}

func (c *XXEVuln1Controller) Post() {
	file := c.GetString("file")
	p := parser.New(parser.XMLParseNoEnt)
	doc, err := p.ParseReader(bytes.NewReader([]byte(file)))
	if err != nil {
		panic(err)
	}
	defer doc.Free()
	root, err := doc.DocumentElement()
	xxe := root.TextContent()
	c.Data["xxe"] = xxe
	c.TplName = "xxe.tpl"
}

func (c *XXESafe1Controller) Get() {
	file, err := ioutil.ReadFile("static/xml/xxe.xml")
	if err != nil {
		panic(err)
	}
	c.Data["xxe"] = string(file)
	c.TplName = "xxe.tpl"
}

func (c *XXESafe1Controller) Post() {
	file := c.GetString("file")
	err := ioutil.WriteFile("tmp/upload.xml", []byte(file), 0777)
	if err != nil {
		panic(err)
	}
	entityMap := make(map[string]string)
	entityMap["xxe"] = "default xxe value"
	doc := etree.NewDocument()
	doc.ReadSettings.Entity = entityMap
	if err := doc.ReadFromFile("tmp/upload.xml"); err != nil {
		fmt.Println(err)
	}
	xxe := doc.SelectElement("root").Text()
	c.Data["xxe"] = xxe
	c.TplName = "xxe.tpl"
}
