package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"quickstart/utils"
)

type Info struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

type ReadConfigController struct {
	beego.Controller
}

func auth(ctx *context.Context) {
	fmt.Println("namespace filter activated.")
}

func (request *ReadConfigController) Get() {
	configName := request.Input().Get("name")
	configData := utils.GetStringConfig(configName)
	info := Info{
		"200",
		"ok",
		configData,
	}
	request.Data["json"] = info
	beego.Informational("result is", configData)
	//返回JSon类型的数据
	request.ServeJSON()
}

/*
我们应用中经常会遇到这样的情况，在 Prepare 阶段进行判断，如果用户认证不通过，就输出一段信息，然后直接中止进程，
之后的 Post、Get 之类的不再执行，那么如何终止呢？可以使用 StopRun 来终止执行逻辑，可以在任意的地方执行。
调用 StopRun 之后，如果你还定义了 Finish 函数就不会再执行，如果需要释放资源，
那么请自己在调用 StopRun 之前手工调用 Finish 函数。
*/

func (request *ReadConfigController) Prepare() {
	//request.ServeJSONP()
	// 可以停止启动
	//request.StopRun()
	fmt.Println("prepare to start server")
}

func (request *ReadConfigController) UpdateConfig() {

	// 获取所有的params
	params := request.Ctx.Input.Params()
	for key, value := range params {
		utils.UpdateStringConfig(key, value)
	}

	for key, _ := range params {
		value := utils.GetStringConfig(key)
		fmt.Println("key:" + key + " value: " + value)
	}

	request.Ctx.WriteString("configure updated.")
}
