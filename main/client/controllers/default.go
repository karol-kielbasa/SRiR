package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Post() {
	file, err := c.Ctx.Request.MultipartForm.File["file"][0].Open()
	if err == nil {
		contentFromFile, err := ioutil.ReadAll(file)
		if err == nil {
			fmt.Println(string(contentFromFile))
		} else {
			c.Ctx.Output.SetStatus(400)
		}
	} else {
		c.Ctx.Output.SetStatus(400)
	}
	c.ServeJSON()
}
