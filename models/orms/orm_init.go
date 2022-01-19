package orms

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"quickstart/utils"
)
func mysqlInit() {
	url := utils.GetStringConfig("mysql.datasource.url")
	username := utils.GetStringConfig("mysql.datasource.username")
	pwdEncrtpy := utils.GetStringConfig("mysql.datasource.password")
	//只解密密码 使用des
	pwd, err := utils.Decrypt(pwdEncrtpy, []byte(utils.EncryptKey))
	if err != nil {
		beego.Error(err)
		panic("mysql password decrypt failed, please check!")
	}
	dataSource := fmt.Sprintf("%s:%s@%s", username, pwd, url)
	err = orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		beego.Error(err)
		panic("register driver failed, please check!")
	}
	err = orm.RegisterDataBase("default", "mysql", dataSource)
	if err != nil {
		beego.Error(err)
		panic("register database failed, please check!")
	}
	orm.SetMaxIdleConns("default", 10)
	orm.SetMaxOpenConns("default", 100)

	orm.RegisterModel(new(User))
	orm.RegisterModel(new(Order))
	orm.RegisterModel(new(Account))

	// 检查配置
	if beego.AppConfig.DefaultBool("mysql.table.create.auto.enable", false) {
		beego.Info("auto create table enabled.")
		// 参数2开启自动建表，参数3是否更新表
		err = orm.RunSyncdb("default", true, false)
		if err != nil {
			beego.Error(err)
			panic("open auto create table failed, please check!")
		}
	}
}

func init() {
	mysqlInit()
}