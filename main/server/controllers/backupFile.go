package controllers

import (
	"SRiR/main/server/models"
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
)

type BackupController struct {
	beego.Controller
}

func (c *BackupController) Get() {

	dir, err := ioutil.ReadDir("./converter/workspace")
	for _, d := range dir {
		os.RemoveAll(path.Join([]string{"converter/workspace", d.Name()}...))
	}

	fileName := c.GetString("fileName")
	fileName2 := strings.Replace(fileName, ".", "-", -1)

	from, err := os.Open("./main/server/filesForReports/" + fileName2 + ".txt")
	if err != nil {
		log.Fatal(err)
	}
	defer from.Close()

	to, err := os.OpenFile("./converter/workspace/"+fileName2+".txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer to.Close()

	_, err = io.Copy(to, from)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := os.Stat("report.txt"); !os.IsNotExist(err) {
		os.Remove("report.txt")
	}

	cmd_output, err := exec.Command("java", "-jar", "project.jar", fileName2+".txt").Output()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("combined out:\n%s\n", string(cmd_output))

	for true {
		if _, err := os.Stat("report.txt"); !os.IsNotExist(err) {
			break
		}
	}

	c.Data["json"] = models.NewResponse("Report file created, you can find it in project folder ", "200")

	c.ServeJSON()

}
