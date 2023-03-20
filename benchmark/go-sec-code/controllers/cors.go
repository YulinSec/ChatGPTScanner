package controllers

import (
	"encoding/json"
	"go-sec-code/utils"

	beego "github.com/beego/beego/v2/server/web"
)

type CorsVuln1Controller struct {
	beego.Controller
}

type CorsVuln2Controller struct {
	beego.Controller
}

type CorsSafe1Controller struct {
	beego.Controller
}

func (c *CorsVuln1Controller) Get() {
	origin := c.Ctx.Request.Header.Get("Origin")
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", origin)
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, X-Extra-Header, Content-Type, Accept, Authorization")
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Credentials", "true")
	jsonp := make(map[string]interface{})
	jsonp["username"] = "admin"
	jsonp["password"] = "admin@123"
	data, err := json.Marshal(jsonp)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(data)
}

func (c *CorsVuln2Controller) Get() {
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, X-Extra-Header, Content-Type, Accept, Authorization")
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Credentials", "true")
	jsonp := make(map[string]interface{})
	jsonp["username"] = "admin"
	jsonp["password"] = "admin@123"
	data, err := json.Marshal(jsonp)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(data)
}

func (c *CorsSafe1Controller) Get() {
	origin := c.Ctx.Request.Header.Get("origin")
	whitelists := []string{"localhost:233", "example.com"}
	corsFilter := utils.CorsFilter{}
	if origin != "" && corsFilter.DoFilter(origin, whitelists) {
		c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", origin)
		c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, X-Extra-Header, Content-Type, Accept, Authorization")
		c.Ctx.ResponseWriter.Header().Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Credentials", "true")
	}
	jsonp := make(map[string]interface{})
	jsonp["username"] = "admin"
	jsonp["password"] = "admin@123"
	data, err := json.Marshal(jsonp)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(data)
}
