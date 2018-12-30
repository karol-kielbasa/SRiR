package main

import (
	_ "SRiR/main/server/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
)

func main() {
	conf, err := config.NewConfig("json", "main/client/conf/app.json")
	if err != nil {
		panic(err)
	}
	port := conf.String("server_port")
	ip := conf.String("server_ip")
	beego.Run(ip + port)
}
