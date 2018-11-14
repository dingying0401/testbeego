package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"fmt"
)

type LoginUser struct {
	Id int
	Username string `form："username"`
	Password string `form: "password"`
}


func RegisterDB() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(LoginUser))
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:dingying@/test?charset=utf8")
}

func Task (data *LoginUser) string{
	o := orm.NewOrm()
	var i int
	for i=0; i<= data.Id ;i++{
	user := LoginUser{Id:i}
	test := o.Read(&user)
		if test != nil {
			return test
		}
	return ""
	}
}

