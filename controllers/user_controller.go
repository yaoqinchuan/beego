package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	"io/ioutil"
	"quickstart/common"
	"strconv"
	"strings"
)

type Sex int

// Enum the Database driver
const (
	_      Sex = iota // int enum type
	Male              // 男
	Female            // 女
)

type User struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Age   int32  `json:"age"`
	Sex   Sex    `json:"sex"`
	IdNum string `json:"id_num"`
	Job   string `json:"job"`
}

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
	beego.Info("afteer filter printf all params.")
}

func findByName(name string) (User, error) {
	o := orm.NewOrm()
	user := User{}
	// 使用QuerySeter来组织查询是一种方式
	num, err := o.QueryTable("user").Filter("name", name).Distinct().All(&user)
	if err != nil {
		beego.Error(fmt.Sprintf("error occurs: %s", err))
		panic(fmt.Sprintf("error occurs: %s", err))
	}
	if num > 1 {
		beego.Error(fmt.Sprintf("name duplicate is database, name: %s", name))
		panic(fmt.Sprintf("name duplicate is database, name: %s", name))
	} else if num == 0 {
		panic(fmt.Sprintf("user is empty, name: %s", name))
	}
	return user, nil
}

func findByIdNum(idNum string) (User, error) {
	o := orm.NewOrm()
	user := User{IdNum: idNum}
	// 默认使用主键，没有主键只能指定列了，同时映射到结构体
	err := o.Read(&user, "id_num")
	if err != nil {
		beego.Error(fmt.Sprintf("error occurs: %s", err))
		panic(fmt.Sprintf("error occurs: %s", err))
	}
	return user, nil
}

func findById(id int64) (User, error) {
	o := orm.NewOrm()

	var user = User{}

	/*
		用于一次 prepare 多次 exec，以提高批量执行的速度，但是对于类型转换不太友好
			p, err := o.Raw("select * from user where id = ?").Prepare()
			res, err :=  p.Exec(id)
	*/

	// 使用原生的语句，同时映射到结构体，随后可以使用RowsToMap等将结构体做映射到特定结构
	err := o.Raw("select * from user where id = ?", id).QueryRow(&user)

	if err != nil {
		beego.Error(fmt.Sprintf("error occurs: %s", err))
		panic(fmt.Sprintf("error occurs: %s", err))
	}
	return user, nil
}

func deleteById(id int64) {

	o := orm.NewOrm()
	// 使用原生的sql
	_, err := o.Raw("delete from user where id = ?", id).Exec()
	if err != nil {
		beego.Error(fmt.Sprintf("error occurs: %s", err))
		panic(fmt.Sprintf("error occurs: %s", err))
	}
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
				user, err := findByName(value[0])
				if err != nil {
					panic(err)
				}
				u.Data["json"] = user
				break
			case "idnum":
				user, err := findByIdNum(value[0])
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
				user, err := findById(id)
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
	user := &User{}
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
				deleteById(id)
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
