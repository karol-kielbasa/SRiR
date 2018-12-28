package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	sendFile()
	c.ServeJSON()
}

func sendFile() {
	req := httplib.Post("http://localhost:8080/file")
	req.PostFile("file", "main/hello.txt")
	res, err := req.Response()
	if err == nil {
		if res.StatusCode == 200 {
			fmt.Println("file sent")
		}
	} else {
		fmt.Println("error")
	}
}
