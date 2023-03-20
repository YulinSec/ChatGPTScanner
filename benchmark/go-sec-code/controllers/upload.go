package controllers

import (
	"fmt"
	"go-sec-code/utils"
	"time"

	beego "github.com/beego/beego/v2/server/web"
)

type FileUploadVuln1Controller struct {
	beego.Controller
}

type FileUploadSafe1Controller struct {
	beego.Controller
}

func (c *FileUploadVuln1Controller) Get() {
	c.TplName = "fileUpload.tpl"
}

func (c *FileUploadVuln1Controller) Post() {
	userid := c.GetString("userid")
	_, h, err := c.GetFile("file")
	if err != nil {
		panic(err)
	}
	savePath := "static/upload/" + userid + fmt.Sprint(time.Now().Unix()) + h.Filename
	c.SaveToFile("file", savePath)
	c.Data["savePath"] = savePath
	c.TplName = "fileUpload.tpl"
}

func (c *FileUploadSafe1Controller) Get() {
	c.TplName = "fileUpload.tpl"
}

func (c *FileUploadSafe1Controller) Post() {
	userid := c.GetString("userid")
	fileUploadFilter := utils.FileUploadFilter{}
	evil := fileUploadFilter.DoFilter(userid)
	if evil == true {
		c.Ctx.ResponseWriter.Write([]byte("evil input"))
		return
	}
	_, h, err := c.GetFile("file")
	if err != nil {
		panic(err)
	}
	savePath := "static/upload/" + userid + fmt.Sprint(time.Now().Unix()) + h.Filename
	c.SaveToFile("file", savePath)
	c.Data["savePath"] = savePath
	c.TplName = "fileUpload.tpl"
}
