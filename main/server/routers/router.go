package routers

import (
	"SRiR/main/server/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/sendFile", &controllers.SendFileController{})
    beego.Router("/buildFile", &controllers.BuildFileController{})
    beego.Router("/executeFile", &controllers.ExecuteFileController{})
}
