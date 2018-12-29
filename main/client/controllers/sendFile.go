package controllers

import (
	"SRiR/main/client/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"io/ioutil"
	"os"
)

type SendFileController struct {
	beego.Controller
}

func (c *SendFileController) Post() {
	file, err := c.Ctx.Request.MultipartForm.File["file"][0].Open()
	fileName := c.Ctx.Request.MultipartForm.File["file"][0].Filename
	if err == nil {
		contentFromFile, err := ioutil.ReadAll(file)
		if err == nil {
			f, err := os.Create("main/client/fileToSend/" + fileName)
			if err == nil {
				f.Write(contentFromFile)
				defer f.Close()
				c.Data["json"] = models.NewResponse("File saved","200")
			}
		} else {
			c.Data["json"] = models.NewResponse("Something went wrong","400")
			c.Ctx.Output.SetStatus(400)
		}
	} else {
		c.Data["json"] = models.NewResponse("Something went wrong","400")
		c.Ctx.Output.SetStatus(400)
	}
	req := httplib.Post("http://localhost:8080/sendFile")
	req.PostFile("file", "main/client/fileToSend/" + fileName)
	res, err := req.Response()
	if err == nil {
		bodyBytes, _ := ioutil.ReadAll(res.Body)
		var response * models.Response
		err = json.Unmarshal(bodyBytes, &response)
		c.Data["json"] = response
	} else {
		fmt.Println("error")
	}
	c.ServeJSON()
}


