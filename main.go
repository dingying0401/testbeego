package main

import (
	"github.com/astaxie/beego/context"
	"strings"
	_ "testbeego/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

)

var FilterUser = func(ctx *context.Context) {
	_, ok := ctx.Input.Session("authResult").(string)
	if !ok && !strings.Contains(ctx.Request.RequestURI, "/login"){
		ctx.Redirect(302, "/login")
	}
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//orm.RegisterDataBase("default", "mysql", "root:dingying@/test?charset=utf8")
	orm.RegisterDataBase("default", "mysql", "root:dingying@tcp(10.71.200.21:3306)/test?charset=utf8")

}

func main() {
	//beego.SetStaticPath("/static","public")
	orm.Debug = true
	//orm.RunSyncdb("default",true, true)
	//open session
	//beego.BConfig.WebConfig.Session.SessionOn = true
	//filter
	//beego.InsertFilter("/home/?:id",beego.BeforeRouter,FilterUser)
	beego.InsertFilter("/user/?:id", beego.BeforeRouter, FilterUser)
	beego.InsertFilter("/login/?:id", beego.BeforeRouter, FilterUser)
	beego.InsertFilter("/admin/?:id",beego.BeforeRouter,FilterUser)

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


