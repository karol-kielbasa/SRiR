package routers

import (
	"SRiR/main/client/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/file", &controllers.MainController{})
}
