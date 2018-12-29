package controllers

import (
	"SRiR/main/server/models"
	"github.com/astaxie/beego"
)

type ExecuteFileController struct {
	beego.Controller
}

func (c *ExecuteFileController) Get() {
	fileName := c.GetString("fileName")
	output, err := executeCmd("go", "run", "main/server/filesFromClients/"+fileName)
	if err == nil {
		c.Data["json"] = models.NewResponse(output, "200")
	} else {
		c.Data["json"] = models.NewResponse(output, "400")
		c.Ctx.Output.SetStatus(400)
	}
	c.ServeJSON()
}
