package routers

import (
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/context"
	"testbeego/controllers"
)

func init() {
	/*RESTful，用户执行Get就执行Get方法,Post就执行Post方法
	第一个rootpath参数是访问的路径url，
	第二个参数是对应的controller（&packagename.Controller），即把请求分发到这个控制器来执行相应的逻辑
	第三个参数是具体路由到的某个方法，*代表所有/任意请求方法都经过某个自定义方法
	*/

	/*登录、注册、错误页*/
	//登陆入口
	beego.Router("/login", &controllers.LoginController{})
	//登录错误
	beego.Router("/login/error", &controllers.LoginController{}, "*:Errorpage")
	//注册
	beego.Router("/register", &controllers.RegisterController{}, "get:Get;post:Register")
	//注册失败
	beego.Router("/register/error", &controllers.RegisterController{}, "*:RegistErrorpage")

	/*商品页面：商品搜索*/
	//展示所有的商品列表
	beego.Router("/home/product", &controllers.ProductController{}, "*:ListUser")
	//根据商品名称（id）搜索某个商品
	beego.Router("/home/productinfo", &controllers.ProductController{}, "*:SearchProduct")

	/*订单处理：购买商品&增加订单记录，查询所有订单列表，查询单个订单列表，删除历史订单记录，用户更新个人信息*/
	//购买商品且生成订单记录
	beego.Router("/home/shop", &controllers.OrderController{}, "*:HandleOrder")
	//查看当前用户所有的历史订单记录
	beego.Router("/user/order", &controllers.OrderController{}, "*:SearchOrder")
	beego.Router("/user/order/detail", &controllers.OrderController{}, "*:CheckOrder")
	beego.Router("/user/order/delete", &controllers.OrderController{}, "*:DeleteOrder")
	beego.Router("/user/info", &controllers.UserController{}, "*:UpdateUserInfo")

	/*管理员操作*/
	//删除某个商品
	beego.Router("/admin/product/delete", &controllers.AdminController{}, "*:DeleteProduct")
	//更新商品信息
	beego.Router("/admin/product/update", &controllers.AdminController{}, "*:UpdateProduct")
	//折扣计算
	beego.Router("/shopcart", &controllers.ProductController{}, "*:CouponCalculate")

	/*vip会员操作*/
	//购买vip会员
	beego.Router("/vip", &controllers.UserController{}, "*:VipOrder")

	//beego.Router("/home", &controllers.MainController{},"*:Evaluate")
	//商品详情
	//beego.Router("/home/detail",&controllers.MainController{},"*:Productdetail")
}
