package main

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/dgrijalva/jwt-go"
	"testbeego/controllers"
	_ "testbeego/routers"
)



var FilterUser = func(ctx *context.Context) {

	var this *controllers.LoginController
	this =new(controllers.LoginController)
	usernames := this.GetString("username")
	token, e := this.ParseToken()
	if e != nil {
		beego.Error(e)
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		errors.New("errPermission")
		return
	}
	var user string = claims["username"].(string)
	if (user != usernames){
		ctx.Redirect(302, "/login")
	}
/*
	_, ok := ctx.Input.Session("authResult").(string)
	if !ok && !strings.Contains(ctx.Request.RequestURI, "/login"){
		ctx.Redirect(302, "/login")

	}
*/
}

func init() {
	/*
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//orm.RegisterDataBase("default", "mysql", "root:dingying@/test?charset=utf8")
	orm.RegisterDataBase("default", "mysql", "root:dingying@tcp(10.71.225.15:3306)/test?charset=utf8")
*/
}

func main() {
	//beego.SetStaticPath("/static","public")
	//orm.Debug = true
	//orm.RunSyncdb("default",true, true)
	//open session
	beego.BConfig.WebConfig.Session.SessionOn = true
	//filter
	beego.InsertFilter("/home/*",beego.BeforeRouter,FilterUser)
	beego.InsertFilter("/user/*", beego.BeforeRouter, FilterUser)
	beego.InsertFilter("/login/?:id", beego.BeforeRouter, FilterUser)
	beego.InsertFilter("/admin/*",beego.BeforeRouter,FilterUser)




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


