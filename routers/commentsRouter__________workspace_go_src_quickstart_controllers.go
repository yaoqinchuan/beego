package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["quickstart/controllers:ReadConfigByAnnotationController"] = append(beego.GlobalControllerRouter["quickstart/controllers:ReadConfigByAnnotationController"],
        beego.ControllerComments{
            Method: "GetConfigByAnnotation",
            Router: "/config/annotation",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["quickstart/controllers:UserController"] = append(beego.GlobalControllerRouter["quickstart/controllers:UserController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/v1/user",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["quickstart/controllers:UserController"] = append(beego.GlobalControllerRouter["quickstart/controllers:UserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/v1/user",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["quickstart/controllers:UserController"] = append(beego.GlobalControllerRouter["quickstart/controllers:UserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/v1/user",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
