package main

import (
	_ "go-sec-code/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}

