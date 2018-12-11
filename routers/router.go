package routers

import (
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/context"
	"testbeego/controllers"
)

func init() {
	//第一个rootpath参数是访问的路径url，第二个参数是对应的controller，即把请求分发到这个控制器来执行相应的逻辑
	beego.Router("/home", &controllers.MainController{},"*:Evaluate")
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/login/error", &controllers.LoginController{},"*:Errorpage")
	beego.Router("/register",&controllers.RegisterController{},"get:Get;post:Register")

	//beego.Router("/home/detail",&controllers.MainController{},"*:Productdetail")
	beego.Router("/home/product",&controllers.ProductController{},"*:ListUser")
	beego.Router("/home/productinfo",&controllers.ProductController{},"*:SearchProduct")

	beego.Router("/home/shop",&controllers.OrderController{},"*:HandleOrder")
	//beego.Router("/user/order",&controllers.OrderController{},"*:SearchOrder")
	//beego.Router("/user/order/detail",&controllers.OrderController{},"*:CheckOrder")
	//beego.Router("/user/order/delete",&controllers.OrderController{},"*:DeleteOrder")

	//beego.Router("/admin/product/delete",&controllers.AdminController{},"*:DeleteProduct")
	//beego.Router("/admin/product/update",&controllers.AdminController{},"*:UpdateProduct")


}

