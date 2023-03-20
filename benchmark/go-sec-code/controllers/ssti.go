package controllers

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"os"

	"github.com/Masterminds/sprig"
	beego "github.com/beego/beego/v2/server/web"
)

type SSTIVuln1Controller struct {
	beego.Controller
}

type SSTISafe1Controller struct {
	beego.Controller
}

func (c *SSTIVuln1Controller) Get() {
	os.Setenv("go-sec-code-secret-key", "b81024f158eefcf60792ae9df9524f82")
	usertemplate := c.GetString("template", "please send your template")
	t := template.New("ssti").Funcs(sprig.FuncMap())
	t, _ = t.Parse(usertemplate)
	buff := bytes.Buffer{}
	err := t.Execute(&buff, struct{}{})
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(&buff)
	if err != nil {
		panic(err)
	}
	c.Data["usertemplate"] = string(data)
	c.TplName = "ssti.tpl"
}

func (c *SSTISafe1Controller) Get() {
	usertemplate := c.GetString("template", "please send your template")
	c.Data["usertemplate"] = usertemplate
	c.TplName = "ssti.tpl"
}
