package controllers

import (
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

type Logininfo struct {
	username string `form："username"`
	password string `form: "password"`
}

func (c *LoginController) Prepare() {

}

func (c *LoginController) Get() {
	//maincontroller里面会有很多方法，这里重新定义Get()方法
	//我们可以通过各种方式获取数据，然后赋值到 this.Data 中，
	// 这是一个用来存储输出数据的 map，可以赋值任意类型的值，这里我们只是简单举例输出两个字符串
	//c.Data["Website"] = "xiaobizai.com"
	//c.Data["Email"] = "xiaobizai@gmail.com"
	//views里面的页面地址名称
	//最后一个就是需要去渲染的模板，this.TplName 就是需要渲染的模板，这里指定了 index.tpl，
	// 如果用户不设置该参数，那么默认会去到模板目录的 Controller/<方法名>.tpl 查找，例如上面的方法会去 maincontroller/get.tpl (文件、文件夹必须小写)。
	c.TplName = "login/login.tpl"
	//当然也可以不使用模版，直接用 this.Ctx.WriteString 输出字符串，如：
	//c.Ctx.WriteString("hhaha")

}
func (c *LoginController) Post() {
	userinfo := Logininfo{}
	userinfo.username = c.GetString("username")
	userinfo.password = c.GetString("password")
	if (userinfo.username == "dingying" && userinfo.password == "dingying") {
		c.Redirect("/home", 301)
	} else {
		c.Redirect("/error", 404)
	}
	return

	/*userinfo := Logininfo{}
	userinfo.username = c.GetString("username")
	userinfo.password = c.GetString("password")
	if userinfo.username == "" || userinfo.password == "" {
		c.Ctx.WriteString("please input username or password!")
		return
	} else {
		if (userinfo.username != "xiaobizai" || userinfo.password != "biebibi") {
			c.Ctx.WriteString("username or password is invalid!")
		} else {
			c.Ctx.WriteString("welcome to xiaobizai's world!")
		}
		return
	}*/

}
