package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	_ "quickstart/models/orms"
	_ "quickstart/models/session"
	_ "quickstart/routers"
	"quickstart/utils"
)
// 自定义一个启动的hookfunc
func initMyOwnerHook() error {
	app_name := utils.GetStringConfig("appname")
	host_os := utils.GetStringConfig("hostos")
	logs.Info("app " + app_name + " start on system " + host_os)
	return nil
}

func main() {
	beego.AddAPPStartHook(initMyOwnerHook)
	beego.Run()
}

/*
beego.Run 执行之后，我们看到的效果好像只是监听服务端口这个过程，但是它内部做了很多事情：
1、解析配置文件
beego 会自动解析在 conf 目录下面的配置文件 app.conf，通过修改配置文件相关的属性，我们可以定义：开启的端口，是否开启 session，应用名称等信息。
2、执行用户的 hookfunc
beego 会执行用户注册的 hookfunc，默认的已经存在了注册 mime，用户可以通过函数 AddAPPStartHook 注册自己的启动函数。
3、是否开启 session
会根据上面配置文件的分析之后判断是否开启 session，如果开启的话就初始化全局的 session。
4、是否编译模板
beego 会在启动的时候根据配置把 views 目录下的所有模板进行预编译，然后存在 map 里面，这样可以有效的提高模板运行的效率，无需进行多次编译。
5、是否开启文档功能
根据 EnableDocs 配置判断是否开启内置的文档路由功能
6、是否启动管理模块
beego 目前做了一个很酷的模块，应用内监控模块，会在 8088 端口做一个内部监听，我们可以通过这个端口查询到 QPS、CPU、内存、GC、goroutine、thread 等统计信息。
7、监听服务端口
这是最后一步也就是我们看到的访问 8080 看到的网页端口，内部其实调用了 ListenAndServe，充分利用了 goroutine 的优势
一旦 run 起来之后，我们的服务就监听在两个端口了，一个服务端口 8080 作为对外服务，另一个 8088 端口实行对内监控。
通过这个代码的分析我们了解了 beego 运行起来的过程，以及内部的一些机制。接下来让我们去剥离 Controller 如何来处理逻辑的。
 */