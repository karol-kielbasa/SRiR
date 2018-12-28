package routers

import (
	"SRiR/main/server/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/file", &controllers.MainController{})
}
