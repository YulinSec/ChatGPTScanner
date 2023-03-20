package controllers

import (
	"go-sec-code/utils"
	"io/ioutil"
	"net/http"

	beego "github.com/beego/beego/v2/server/web"
)

type SSRFVuln1Controller struct {
	beego.Controller
}

type SSRFVuln2Controller struct {
	beego.Controller
}

type SSRFVuln3Controller struct {
	beego.Controller
}

type SSRFSafe1Controller struct {
	beego.Controller
}

func (c *SSRFVuln1Controller) Get() {
	url := c.GetString("url", "http://www.example.com")
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(body)
}

//bypass can be :
//http://LOCALHOST:233
//http://localhost.:233
//http://0:233
//and others
func (c *SSRFVuln2Controller) Get() {
	url := c.GetString("url", "http://www.example.com")
	ssrfFilter := utils.SSRFFilter{}
	blacklists := []string{"localhost", "127.0.0.1"}
	evil := ssrfFilter.DoBlackFilter(url, blacklists)
	if evil == true {
		c.Ctx.ResponseWriter.Write([]byte("evil input"))
	} else {
		res, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}
		c.Ctx.ResponseWriter.Write(body)
	}
}

func (c *SSRFVuln3Controller) Get() {
	url := c.GetString("url", "http://www.example.com")
	ssrfFilter := utils.SSRFFilter{}
	evil := ssrfFilter.DoGogsFilter(url)
	if evil == true {
		c.Ctx.ResponseWriter.Write([]byte("evil input"))
	} else {
		res, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}
		c.Ctx.ResponseWriter.Write(body)
	}
}

func (c *SSRFSafe1Controller) Get() {
	url := c.GetString("url", "http://www.example.com")
	ssrfFilter := utils.SSRFFilter{}
	whitelists := []string{"example.com"}
	evil := ssrfFilter.DoWhiteFilter(url, whitelists)
	if evil == true {
		c.Ctx.ResponseWriter.Write([]byte("evil input"))
	} else {
		res, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}
		c.Ctx.ResponseWriter.Write(body)
	}
}
