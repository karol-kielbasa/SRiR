package controllers

import (
	"SRiR/main/server/models"
	"bytes"
	"fmt"
	"github.com/astaxie/beego"
	"os/exec"
)

type BuildFileController struct {
	beego.Controller
}

func (c *BuildFileController) Get() {
	fileName := c.GetString("fileName")
	output, err := executeCmd("go", "build", "main/server/filesFromClients/"+fileName)
	if err == nil {
		c.Data["json"] = models.NewResponse("File built successfully", "200")
	} else {
		c.Data["json"] = models.NewResponse(output, "400")
		c.Ctx.Output.SetStatus(400)
	}
	c.ServeJSON()
}

func executeCmd(cmd string, args ...string) (string, error) {
	command := exec.Command(cmd, args...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	command.Stdout = &out
	command.Stderr = &stderr
	err := command.Run()
	if err != nil {
		errstring := fmt.Sprintf(fmt.Sprint(err) + ": " + stderr.String())
		return errstring, err
	}
	return out.String(), nil
}
