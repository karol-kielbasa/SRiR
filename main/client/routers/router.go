package routers

import (
	"SRiR/main/client/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/sendFile", &controllers.SendFileController{})
	beego.Router("/buildFile", &controllers.BuildFileController{})
	beego.Router("/index", &controllers.IndexController{})
	beego.Router("/executeFile", &controllers.ExecuteFileController{})
}
