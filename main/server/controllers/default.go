package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"io/ioutil"
	"log"
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
	bt, err := ioutil.ReadFile("main/hello.txt")
	if err != nil {
		log.Fatal("read file err:", err)
	}
	req.Body(bt)
	req.DoRequest()
}
