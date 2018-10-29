package controllers

import (
	"github.com/astaxie/beego"

	"fmt"
)

type MainController struct {
	beego.Controller
}

type Productinfo struct{
	name string `form name`
	length float32 `form length`
	weight float64 `form weight`
	briefinfo string `form briefinfo`
	loseweight float64 `form loseweight`
	restweight float64 `form restweight`
	singleprice float64 `form -`
	totalprice float64 `form -`

}

func (c *MainController) Prepare() {

}

/*func (c *MainController) Get() {
	//maincontroller里面会有很多方法，这里重新定义Get()方法
	//我们可以通过各种方式获取数据，然后赋值到 this.Data 中，
	// 这是一个用来存储输出数据的 map，可以赋值任意类型的值，这里我们只是简单举例输出两个字符串
	c.Data["Website"] = "xiaobizai.com"
	c.Data["Email"] = "xiaobizai@gmail.com"
	//views里面的页面地址名称
	//最后一个就是需要去渲染的模板，this.TplName 就是需要渲染的模板，这里指定了 index.tpl，
	// 如果用户不设置该参数，那么默认会去到模板目录的 Controller/<方法名>.tpl 查找，例如上面的方法会去 maincontroller/get.tpl (文件、文件夹必须小写)。
	c.TplName = "index.tpl"
	//当然也可以不使用模版，直接用 this.Ctx.WriteString 输出字符串，如：
	//c.Ctx.WriteString("hhaha")
}
*/

func (c *MainController) Buying() {
	boyinfo:=Productinfo{}
	boyinfo.name = "Xiao Bi Zai"
	boyinfo.length = 180
	boyinfo.weight = 170
	boyinfo.singleprice = 5
	boyinfo.loseweight,_ = c.GetFloat("loseweight")
	boyinfo.restweight = boyinfo.weight - boyinfo.loseweight
	boyinfo.totalprice = boyinfo.singleprice * boyinfo.loseweight
	boyinfo.briefinfo ="Is he the most handsome one in the world? No! Once you buy his weight, he will be, hurry up to buy!"
	if (boyinfo.restweight == 0) {
		c.Ctx.WriteString("Sorry,Xiao bi Zai has sold out ")

	}else{
		fmt.Println(&boyinfo)
		//c.Data["json"]=&boyinfo
		//c.ServeJSON()
	}
	c.TplName = "index.tpl"
	return

}


