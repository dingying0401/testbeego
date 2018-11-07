package routers

import (
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/context"
	"testbeego/controllers"
)

func init() {
	//第一个rootpath参数是访问的路径url，第二个参数是对应的controller，即把请求分发到这个控制器来执行相应的逻辑
	beego.Router("/home", &controllers.MainController{},"*:Evaluate")
	beego.Router("/login?:id", &controllers.LoginController{})
	beego.Router("/error", &controllers.LoginController{},"*:Errorpage")
	beego.Router("/detail",&controllers.MainController{},"*:Productdetail")
	//beego.Router("/product", &controllers.SaleController{},"*:Buying")
	/*beego.Router("/login?:id", &controllers.LoginController{},"post:Post")

	beego.Get("/login/test",func(ctx *context.Context){
		ctx.Output.Body([]byte("hello world"))
	})

	beego.Post("/alice",func(ctx *context.Context){
		ctx.Output.Body([]byte("bob"))
		//ctx.Output.Header([]byte("dingying","dingying"))
	})*/
}
