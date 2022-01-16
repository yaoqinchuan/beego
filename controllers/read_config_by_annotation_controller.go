package controllers

import (
	"github.com/astaxie/beego"
	"quickstart/utils"
)

type ReadConfigByAnnotationController struct {
	beego.Controller
}

//增加了 URLMapping 这个函数，这是新增加的函数，用户如果没有进行注册，那么就会通过反射来执行对应的函数，
//如果注册了就会通过 interface 来进行执行函数，性能上面会提升很多。
func (r *ReadConfigByAnnotationController) URLMapping() {
	r.Mapping("GetConfigByAnnotation", r.GetConfigByAnnotation)
}

// @router /config/annotation [get]
func (request *ReadConfigByAnnotationController) GetConfigByAnnotation() {
	configName := request.Input().Get("name")
	// configName = request.GetString("name", "appname")
	configData := utils.GetStringConfig(configName)
	request.Ctx.WriteString("result:" + configData)
}
