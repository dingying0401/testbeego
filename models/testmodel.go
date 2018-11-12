package models

import (
"github.com/astaxie/beego/orm"
_ "github.com/go-sql-driver/mysql"
		)

type LoginUser struct {
	Id int
	username string `form："username"`
	password string `form: "password"`
}

type Userinfo struct{
	Id int
	Name string `form:"name,text,姓名：  "`
	Length float32 `form:"length,text,身高： "`
	Weight float64 `form:"weight,text,体重："`
	Briefinfo string `form:"briefinfo,text,简介："`
	Loseweight float64 `form:"loseweight,text,购买体重： "`
	Restweight float64 `form:"restweight,text,剩余体重： "`
	Singleprice float64 `form:"singleprice,text,单价： "`
	Totalprice float64 `form:"totalprice,text,总价： "`

}

func RegisterDB() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(LoginUser), new(Userinfo))
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("mysql", "mysql", "root:dingying@/test?charset=utf8")

}

