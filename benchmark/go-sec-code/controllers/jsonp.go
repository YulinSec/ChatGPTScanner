package controllers

import (
	"encoding/json"
	"go-sec-code/utils"

	beego "github.com/beego/beego/v2/server/web"
)

type JsonpVuln1Controller struct {
	beego.Controller
}

type JsonpVuln2Controller struct {
	beego.Controller
}

type JsonpSafe1Controller struct {
	beego.Controller
}

func (c *JsonpVuln1Controller) Get() {
	callback := c.GetString("callback")
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "application/javascript")
	jsonp := make(map[string]interface{})
	jsonp["username"] = "admin"
	jsonp["password"] = "admin@123"
	data, err := json.Marshal(jsonp)
	output := callback + "(" + string(data) + ")"
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write([]byte(output))
}

func (c *JsonpVuln2Controller) Get() {
	callback := c.GetString("callback")
	referer := c.Ctx.Request.Header.Get("referer")
	jsonpFilter := utils.JsonpFilter{}
	whitelists := []string{"localhost:233", "example.com"}
	if referer == "" || jsonpFilter.DoFilter(referer, whitelists) {
		c.Ctx.ResponseWriter.Header().Set("Content-Type", "application/javascript")
		jsonp := make(map[string]interface{})
		jsonp["username"] = "admin"
		jsonp["password"] = "admin@123"
		data, err := json.Marshal(jsonp)
		output := callback + "(" + string(data) + ")"
		if err != nil {
			panic(err)
		}
		c.Ctx.ResponseWriter.Write([]byte(output))
	} else {
		c.Ctx.ResponseWriter.Write([]byte("evil input"))
	}
}

func (c *JsonpSafe1Controller) Get() {
	callback := c.GetString("callback")
	referer := c.Ctx.Request.Header.Get("referer")
	jsonpFilter := utils.JsonpFilter{}
	whitelists := []string{"localhost:233", "example.com"}
	if referer != "" && jsonpFilter.DoFilter(referer, whitelists) {
		c.Ctx.ResponseWriter.Header().Set("Content-Type", "application/javascript")
		jsonp := make(map[string]interface{})
		jsonp["username"] = "admin"
		jsonp["password"] = "admin@123"
		data, err := json.Marshal(jsonp)
		output := callback + "(" + string(data) + ")"
		if err != nil {
			panic(err)
		}
		c.Ctx.ResponseWriter.Write([]byte(output))
	} else {
		c.Ctx.ResponseWriter.Write([]byte("evil input"))
	}
}
