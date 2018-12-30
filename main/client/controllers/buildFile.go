package controllers

import (
	"SRiR/main/client/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"io/ioutil"
)

type BuildFileController struct {
	beego.Controller
}

func (c *BuildFileController) Get() {
	fileName := c.GetString("fileName")
	req := httplib.Get("http://"+ getServerPortAndIp()+ "/buildFile")
	req.Param("fileName", fileName)
	res, err := req.Response()

	if err == nil {
		bodyBytes, _ := ioutil.ReadAll(res.Body)
		var response *models.Response
		err = json.Unmarshal(bodyBytes, &response)
		c.Data["json"] = response
	} else {
		c.Data["json"] = models.NewResponse("Something went wrong", "400")
		fmt.Println("error")
	}
	c.ServeJSON()
}
