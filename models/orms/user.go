package orms

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
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

/* 总结一下，查询的几种方式
1、使用QueryTable进行查询；
2、使用QueryBuilder生成查询语句；
3、使用基础CRUD操作；
4、使用原生Raw查询；
*/
func FindUserByName(name string) (User, error) {
	qb, _ := orm.NewQueryBuilder("go")
	qb.Select("*").From("user").Where("name=?")
	sql := qb.String()
	fmt.Println(sql)

	o := orm.NewOrm()
	user := User{}

	// 使用QuerySeter来组织查询是一种方式
	num, err := o.QueryTable("user").Filter("name", name).Distinct().All(&user)

	// 也可以使用QueryBuilder 结合raw
	//o.Raw(sql, name).QueryRow(&user)

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

func FindUserByIdNum(idNum string) (User, error) {
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

func FindUserById(id int64) (User, error) {
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

func DeleteUserById(id int64) {

	o := orm.NewOrm()
	// 使用原生的sql
	_, err := o.Raw("delete from user where id = ?", id).Exec()
	if err != nil {
		beego.Error(fmt.Sprintf("error occurs: %s", err))
		panic(fmt.Sprintf("error occurs: %s", err))
	}
}
