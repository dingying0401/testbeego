package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	)

type Userinfo struct{
	Id int
	Name string `form:"name,text,姓名：  "`
	Length float32 `form:"length,text,身高： "`
	Weight float64 `form:"weight,text,体重："`
	Briefinfo string `form:"briefinfo,text,简介："`
	Singleprice float64 `form:"singleprice,text,单价： "`
	Totalprice float64 `form:"totalprice,text,总价： "`
}

type Orderinfo struct {
	Id int
	Loseweight float64 `form:"loseweight,text,购买体重： "`
	Restweight float64 `form:"restweight,text,剩余体重： "`
}

func RegisterDB4UserInfo() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Userinfo),new(Orderinfo))
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:dingying@/test?charset=utf8")

}
/*

func BoyInfo(info *Userinfo) *Userinfo {
	o := orm.NewOrm()
	boyinfo := Userinfo{Id: 1}
	name := o.Read(&boyinfo)
	fmt.Println(o.Insert(boyinfo))

}
*/