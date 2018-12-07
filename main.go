package main

import (
	_ "testbeego/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:dingying@/test?charset=utf8")
	//orm.RegisterDataBase("default", "mysql", "root:dingying@tcp(10.71.200.74:3306)/test?charset=utf8")
}

func main() {
	//beego.SetStaticPath("/static","public")
	/*beego.SetStaticPath("/asd","conf")*/
	//beego.SetStaticPath("/asd", "static/img")
	orm.Debug = true
	//orm.RunSyncdb("default",true, true)
	/*o := orm.NewOrm()
	userlogin := new(models2.LoginUser)
	userlogin.Username = "chenlaogoubi"
	userlogin.Password ="chenchenlaogoubi"
	fmt.Println(o.Insert(userlogin))*/
	beego.Run()
	/*
	1.app.conf
	2.hookfunc (函数：AddAPPStartHook)
	3.session or not
	4.是否编译模板：beego会在启动时候根据配置把views目录下的模板进行预编译，然后存在map里面，这样可以有效提高模板运行的效率，无须多次编译
	5.是否开启文档功能：根据 EnableDocs 配置判断是否开启内置的文档路由功能
	6.是否启动管理模块：EnableAdmin
	7.监听服务端口：ListenAndServer
	 */
}

