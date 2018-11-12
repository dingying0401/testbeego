package main

import (
			_ "testbeego/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"fmt"
		models2	"testbeego/models"
)



func init() {
     models2.RegisterDB()

}

/*func main() {
	//beego.SetStaticPath("/static","public")
	/*beego.SetStaticPath("/asd","conf")*/
	//beego.SetStaticPath("/asd", "static/img")
	//o := orm.NewOrm()

	//userlogin := new(LoginUser)
	//fmt.Println(o.Insert(userlogin))
	/*
	1.app.conf
	2.hookfunc (函数：AddAPPStartHook)
	3.session or not
	4.是否编译模板：beego会在启动时候根据配置把views目录下的模板进行预编译，然后存在map里面，这样可以有效提高模板运行的效率，无须多次编译
	5.是否开启文档功能：根据 EnableDocs 配置判断是否开启内置的文档路由功能
	6.是否启动管理模块：EnableAdmin
	7.监听服务端口：ListenAndServ
	 */


func main() {
	//beego.SetStaticPath("/static","public")
	/*beego.SetStaticPath("/asd","conf")*/
	//beego.SetStaticPath("/asd", "static/img")
	beego.Run()
	o := orm.NewOrm()
	userlogin := new(models2.LoginUser)
	fmt.Println(o.Insert(userlogin))
	/*
	1.app.conf
	2.hookfunc (函数：AddAPPStartHook)
	3.session or not
	4.是否编译模板：beego会在启动时候根据配置把views目录下的模板进行预编译，然后存在map里面，这样可以有效提高模板运行的效率，无须多次编译
	5.是否开启文档功能：根据 EnableDocs 配置判断是否开启内置的文档路由功能
	6.是否启动管理模块：EnableAdmin
	7.监听服务端口：ListenAndServ
	 */
}

