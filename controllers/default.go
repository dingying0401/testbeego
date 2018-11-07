package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"log"

	)

type MainController struct {
	beego.Controller
}

type Productinfo struct{
	Name string `form:"name,text,姓名：  "`
	Length float32 `form:"length,text,身高： "`
	Weight float64 `form:"weight,text,体重："`
	Briefinfo string `form:"briefinfo,text,简介："`
	Loseweight float64 `form:"loseweight,text,购买体重： "`
	Restweight float64 `form:"restweight,text,剩余体重： "`
	Singleprice float64 `form:"singleprice,text,单价： "`
	Totalprice float64 `form:"totalprice,text,总价： "`

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

func (c *MainController) Evaluate() {
	boyinfo := Productinfo{}
	boyinfo.Name = "Chen Chen"
	boyinfo.Length = 180
	boyinfo.Weight = 170
	boyinfo.Singleprice = 5
	boyinfo.Loseweight, _ = c.GetFloat("loseweight")
	boyinfo.Restweight = boyinfo.Weight - boyinfo.Loseweight
	boyinfo.Totalprice = boyinfo.Singleprice * boyinfo.Loseweight
	boyinfo.Briefinfo = "Is he the most handsome one in the world? No! Once you buy his weight, he will be, hurry up to buy!"
	if (boyinfo.Restweight <= 0) {
		c.Ctx.WriteString("Sorry,Xiao bi Zai has sold out ")

	} else {
		//c.Ctx.WriteString("name: "+boyinfo.Name)
		fmt.Println(&boyinfo)
		c.Data["Form"] = &boyinfo
		//c.Data["json"]=&boyinfo
		//c.ServeJSON()
		c.TplName = "index.tpl"
		return

	}
}
func (c *MainController) Upload(){
		f, h, err := c.GetFile("uploadname")
		if err != nil {
			log.Fatal("getfile err ", err)
		}
		defer f.Close()
		c.SaveToFile("uploadname", "static/upload/" + h.Filename) // 保存位置在 static/upload, 没有文件夹要先创建

	}
func (c *MainController) Productdetail(){
	beego.URLFor("Maincontroller.Productdetail")
	c.TplName="productdetail.tpl"

}









