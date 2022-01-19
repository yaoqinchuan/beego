package orms

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
)

type Account struct {
	Id       int64  `orm:"column(id);auto" description:"账户主键"`
	Name     string `orm:"column(name);size(20)" description:"账户名称"`
	Password string `orm:"column(password);size(40)" description:"账户密码"`
}

func (account *Account) TableName() string {
	return "account"
}

func FindAccountByName(name string) (Account, error) {
	qb, _ := orm.NewQueryBuilder("go")
	qb.Select("*").From("account").Where("name=?")
	sql := qb.String()
	fmt.Println(sql)
	o := orm.NewOrm()
	account := Account{}

	// 使用QuerySeter来组织查询是一种方式
	num, err := o.QueryTable("account").Filter("name", name).Distinct().All(&account)

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
	return account, nil
}

func FindAccountById(id int64) (Account, error) {
	o := orm.NewOrm()

	var account = Account{}

	err := o.Raw("select * from user where id = ?", id).QueryRow(&account)

	if err != nil {
		beego.Error(fmt.Sprintf("error occurs: %s", err))
		panic(fmt.Sprintf("error occurs: %s", err))
	}
	return account, nil
}

func DeleteAccountById(id int64) {

	o := orm.NewOrm()
	// 使用原生的sql
	_, err := o.Raw("delete from account where id = ?", id).Exec()
	if err != nil {
		beego.Error(fmt.Sprintf("error occurs: %s", err))
		panic(fmt.Sprintf("error occurs: %s", err))
	}
}

func (account Account) ValidateUser() error {
	o := orm.NewOrm() //获得用于操作数据库的orm
	var accountInDb Account
	_ = o.Raw("select * from account where name = ? and password = ?", account.Name, account.Password).QueryRow(&accountInDb)
	if accountInDb.Name == "" {
		return errors.New("用户名或密码错误！")
	}
	return nil
}
