package orms

import (
	"time"
)

type Order struct {
	Id         int64     `orm:"column(id);auto" description:"订单主键"`
	UserId     int       `orm:"column(user_id);index" description:"创建订单的人"`
	GoodName   string    `orm:"column(good_name);size(60)" description:"商品名称"`
	Price      float64   `orm:"column(price);digits(12);decimals(4);default(0)" description:"商品价格"`
	OrderNum   string    `orm:"column(order_num);size(28);unique" description:"订单单号"`
	CreateTime time.Time `orm:"column(create_time);auto_now_add;type(datetime)" description:"创建时间"`
	EndTime    time.Time `orm:"column(end_time);auto_now;type(datetime)" description:"订单结束时间"`
	Status     string    `orm:"column(status);size(10)" description:"订单状态"`
}

func (u *Order) TableName() string {
	return "user_order"
}

