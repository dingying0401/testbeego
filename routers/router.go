package routers

import (
	"testbeego/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//第一个rootpath参数是访问的路径url，第二个参数是对应的controller，即把请求分发到这个控制器来执行相应的逻辑
    beego.Router("/login", &controllers.MainController{})

}
