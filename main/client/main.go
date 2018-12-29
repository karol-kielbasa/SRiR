package main

import (
	_ "SRiR/main/client/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.LoadAppConfig("conf","main/client/conf")
	beego.BConfig.WebConfig.ViewsPath = "main/client/views"
	beego.Run()
}

