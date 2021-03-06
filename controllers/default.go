package controllers

import (
	"github.com/astaxie/beego"
	"quickstart/utils"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *MainController) ShowVersion() {
	version := utils.GetStringConfig("version")
	c.Ctx.Output.Body([]byte(version))
}
