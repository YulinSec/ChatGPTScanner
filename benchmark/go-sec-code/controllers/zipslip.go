package controllers

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	beego "github.com/beego/beego/v2/server/web"
)

type ZipSlipVuln1Controller struct {
	beego.Controller
}

func (c *ZipSlipVuln1Controller) Get() {
	c.TplName = "fileUpload.tpl"
}

func (c *ZipSlipVuln1Controller) Post() {
	_, h, err := c.GetFile("file")
	if err != nil {
		panic(err)
	}
	timestamp := fmt.Sprint(time.Now().Unix())
	savePath := "static/upload/" + timestamp + h.Filename
	c.SaveToFile("file", savePath)
	unzipPath := "static/unzip/" + timestamp + h.Filename
	r, err := zip.OpenReader(savePath)
	if err != nil {
		panic(err)
	}
	for _, f := range r.File {
		fpath := filepath.Join(unzipPath, f.Name)
		if f.FileInfo().IsDir() {
			// Make Folder
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		// Make File
		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			panic(err)
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			panic(err)
		}

		rc, err := f.Open()
		if err != nil {
			panic(err)
		}

		_, err = io.Copy(outFile, rc)

		// Close the file without defer to close before next iteration of loop
		outFile.Close()
		rc.Close()

		if err != nil {
			panic(err)
		}
	}
	c.Data["savePath"] = unzipPath
	c.TplName = "fileUpload.tpl"
}
