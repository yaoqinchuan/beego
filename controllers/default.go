package controllers

import (
	"fmt"
	_ "fmt"
	"github.com/astaxie/beego"
	"quickstart/models/orms"
	_ "quickstart/utils"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	sess := this.GetSession("name")    //判断此次会话的session是否已经存在
	if sess == nil{
		this.Redirect("/login",301)    //跳转到登录逻辑
	} else {
		account := sess.(orms.Account)
		this.Data["user"] = account.Name    //用于向前端页面传送数据
		this.Data["pass"] = account.Password
		this.TplName = "succeed.html"    //渲染succeed.html页面
	}
}

func (c *MainController) Login() {
	c.TplName = "login.html"
}

func (c *MainController) LoginCheck() {
	var account orms.Account
	inputs := c.Input()
	account.Name = inputs.Get("name")
	account.Password = inputs.Get("password")
	err := account.ValidateUser()
	if err == nil {
		c.SetSession("name", account)
		c.Redirect("/",301)
	} else {
		fmt.Println(err)
		c.Data["info"] = err
		c.TplName = "error.html"
	}
}

func (c *MainController) Out() {
	var account orms.Account
	inputs := c.Input()
	account.Name = inputs.Get("name")
	c.DelSession("name")
	c.Redirect("/",301)
}