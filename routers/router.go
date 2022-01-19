package routers

import (
	"github.com/astaxie/beego"
	"quickstart/controllers"
)
/*
// 可以前置检查token
func filter(ctx *context.Context) {
	token := ctx.Request.URL.Query().Get("token")
	if token == "" {
		panic("token is empty,please check!")
	}
}
*/

// import 就会执行
func init() {
	defer func() {
		if err := recover(); err != nil {
			beego.Error(err)
		}
	}()

	// 添加前置拦截器
	// beego.InsertFilter("/*", beego.BeforeRouter, filter)

	// 常见的做法，执行controller的默认GET等
	beego.Router("/", &controllers.MainController{})
	// 测试session
	beego.Router("/login", &controllers.MainController{},"get:Login;post:LoginCheck")

	beego.Router("/out", &controllers.MainController{},"post:Out")

	// 覆盖之后默认方法就用不了，有点蠢
	beego.Router("/v1/config", &controllers.ReadConfigController{}, "get:Get;post:UpdateConfig")

	// 使用注解的方式进行controller生成
	beego.Include(&controllers.ReadConfigByAnnotationController{})

	// 使用namespace进行路由注册
	userController := controllers.UserController{}
	userController.InitNameSpaces()
}
