package controllers

import (
	"SRiR/main/server/models"
	"github.com/astaxie/beego"
	"io/ioutil"
	"os"
	"strings"
)

type SendFileController struct {
	beego.Controller
}

func (c *SendFileController) Post() {
	file, err := c.Ctx.Request.MultipartForm.File["file"][0].Open()
	fileNamePath := c.Ctx.Request.MultipartForm.File["file"][0].Filename
	result := strings.Split(fileNamePath, "/")
	fileName := result[len(result)-1]
	if err == nil {
		contentFromFile, err := ioutil.ReadAll(file)
		if err == nil {
			f, err := os.Create("main/server/filesFromClients/" + fileName)
			if err == nil {
				_, err = f.Write(contentFromFile)
				defer f.Close()
				c.Data["json"] = models.NewResponse("File saved", "200")
			} else {
				c.Data["json"] = models.NewResponse("Something went wrong", "400")
				c.Ctx.Output.SetStatus(400)
			}
		} else {
			c.Data["json"] = models.NewResponse("Something went wrong", "400")
			c.Ctx.Output.SetStatus(400)
		}
	} else {
		c.Data["json"] = models.NewResponse("Something went wrong", "400")
		c.Ctx.Output.SetStatus(400)
	}
	c.ServeJSON()
}
