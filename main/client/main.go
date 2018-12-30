package main

import (
	_ "SRiR/main/client/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
)

func main() {
	conf, err := config.NewConfig("json", "main/client/conf/app.json")
	if err != nil {
		panic(err)
	}
	port := conf.String("client_port")
	clientIp := conf.String("client_ip")
	beego.BConfig.WebConfig.ViewsPath = "main/client/views"
	beego.Run(clientIp + port)
}
