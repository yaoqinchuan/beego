package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	"io/ioutil"
	"quickstart/common"
	"quickstart/models/orms"
	"strconv"
	"strings"
)

type UserController struct {
	beego.Controller
}

// 前置过滤器
func beforeFilter(c *context.Context) {
	beego.Info("before filter printf all params.")
	for key, value := range c.Request.URL.Query() {
		paramInfo := fmt.Sprintf("%s %s", key, value)
		beego.Info(paramInfo)
	}
}

// 后置过滤器
func afterFilter(c *context.Context) {
	beego.Info("after filter printf all params.")
}

func getProccess(u *UserController) {
	defer func() {
		if err := recover(); err != nil {
			u.Data["json"] = common.DefaultFailedRestResult(err)
		}
		u.ServeJSON()
	}()
	ctx := u.Ctx
	params := ctx.Request.URL.Query()
	if params == nil {
		u.Data["json"] = common.DefaultFailedRestResult("params is empty.")
	} else {
		for key, value := range params {
			switch strings.ToLower(key) {
			case "name":
				user, err := orms.FindUserByName(value[0])
				if err != nil {
					panic(err)
				}
				u.Data["json"] = user
				break
			case "idnum":
				user, err := orms.FindUserByIdNum(value[0])
				if err != nil {
					panic(err)
				}
				u.Data["json"] = user
				break
			case "id":
				id, err := strconv.ParseInt(value[0], 10, 64)
				if err != nil {
					panic(err)
				}
				user, err := orms.FindUserById(id)
				if err != nil {
					panic(err)
				}
				u.Data["json"] = user
				break
			default:
				panic("invalid query field!")
			}
			break
		}
	}
}
func postProcess(u *UserController) {
	defer func() {
		if err := recover(); err != nil {
			u.Data["json"] = common.DefaultFailedRestResult(err)
		}
		u.ServeJSON()
	}()
	o := orm.NewOrm()
	user := &orms.User{}
	b, err := ioutil.ReadAll(u.Ctx.Request.Body)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(b, user)
	line, err := o.InsertOrUpdate(user, "id", "id=id+1")
	if err != nil {
		panic(err)
	} else {
		u.Data["json"] = common.SuccessRestResult(line)
	}

}

func deleteProcess(u *UserController) {
	defer func() {
		if err := recover(); err != nil {
			u.Data["json"] = common.DefaultFailedRestResult(err)
		}
		u.ServeJSON()
	}()
	ctx := u.Ctx
	params := ctx.Request.URL.Query()
	if params == nil {
		u.Data["json"] = common.DefaultFailedRestResult("params is empty.")
	} else {
		for key, value := range params {
			switch strings.ToLower(key) {
			case "id":
				id, err := strconv.ParseInt(value[0], 10, 64)
				if err != nil {
					panic(err)
				}
				orms.DeleteUserById(id)
				u.Data["json"] = common.SuccessRestResult("delete success!")
				break
			default:
				u.Data["json"] = common.DefaultFailedRestResult("delete by user id" + value[0] + " failed!")
			}
			break
		}
	}
}

// @router /v1/user [get]
func (u *UserController) Get() {
	getProccess(u)
}

// @router /v1/user [post]
func (u *UserController) Post() {
	postProcess(u)
}

// @router /v1/user [delete]
func (u *UserController) Delete() {
	deleteProcess(u)
}

func initNameSpace() {
	ns := beego.NewNamespace("/v1/user", beego.NSCond(
		func(ctx *context.Context) bool {
			if ctx.Request.URL.Query().Get("token") == "" || ctx.Request.URL.Query().Get("token") == "user_token" {
				beego.Error("token is invalid.")
				return false
			}
			return true
		}),
		beego.NSBefore(beforeFilter),
		beego.NSRouter("", &UserController{}),
		beego.NSAfter(afterFilter),
	)
	beego.AddNamespace(ns)
}

func (u *UserController) InitNameSpaces() {
	initNameSpace()
}
