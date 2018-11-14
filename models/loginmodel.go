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

func Task (data *LoginUser) error{
	o := orm.NewOrm()
	user := LoginUser{Id:1}
	test := o.Read(&user)
		if test != nil {
			fmt.Printf("---------->"+user.Username, ":"+user.Password)
			fmt.Println("testifithasbeenprinted",test)
		} else{
			fmt.Println("testifithasbeenprinted",test)
			return test
	}
	return test
}

