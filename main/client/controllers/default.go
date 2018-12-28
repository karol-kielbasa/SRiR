package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Post() {

	data := c.Ctx.Request.Body.Read;
	fmt.Print(data)
	c.ServeJSON()
}
