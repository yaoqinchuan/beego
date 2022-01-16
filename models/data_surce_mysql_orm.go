package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"quickstart/controllers"
	"quickstart/utils"
)

func init(){
	url := utils.GetStringConfig("mysql.datasource.url")
	username := utils.GetStringConfig("mysql.datasource.username")
	pwdEncrtpy := utils.GetStringConfig("mysql.datasource.password")
	//只解密密码 使用des
	pwd, err:= utils.Decrypt(pwdEncrtpy, []byte(utils.EncryptKey))
	if err != nil {
		beego.Error(err)
		panic("mysql password decrypt failed, please check!")
	}
	dataSource := fmt.Sprintf("%s:%s@%s", username, pwd, url)
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dataSource)
	orm.RegisterModel(new(controllers.User))
	orm.SetMaxIdleConns("default", 10)
	orm.SetMaxOpenConns("default", 100)
}