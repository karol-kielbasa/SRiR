package controllers

import (
	"SRiR/main/server/models"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"os"
	"strings"
)

type ExecuteFileController struct {
	beego.Controller
}

func (c *ExecuteFileController) Get() {
	fileName := c.GetString("fileName")
	output, err := executeCmd("go", "run", "main/server/filesFromClients/"+fileName)
	if err == nil {
		c.Data["json"] = models.NewResponse(output, "200")
		file, err := os.Open("main/server/filesFromClients/" + fileName)
		if err == nil {
			contentFromFile, err := ioutil.ReadAll(file)
			if err == nil {
				fileName2 := strings.Replace(fileName, ".", "-", -1)
				f, err := os.Create("main/server/filesForReports/" + fileName2 + ".txt")
				if err == nil {
					_, err = f.Write(contentFromFile)
					defer f.Close()
				} else {
					fmt.Print("Error writing content to .txt file")
				}
			} else {
				fmt.Print("Error creating .txt file")
			}
		} else {
			fmt.Print("Error opening file")
		}
	} else {
		c.Data["json"] = models.NewResponse(output, "400")
		c.Ctx.Output.SetStatus(400)
	}
	c.ServeJSON()
}
