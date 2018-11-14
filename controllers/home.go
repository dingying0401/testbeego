package controllers

import (
	"github.com/astaxie/beego"
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
func (c *MainController) Evaluate() {
	boyinfo := Productinfo{}
	boyinfo.Name = "chen chen"
	boyinfo.Weight = 1800
	boyinfo.Briefinfo ="xiao bi zai"
	boyinfo.Loseweight, _ = c.GetFloat("loseweight")
	boyinfo.Restweight = boyinfo.Weight - boyinfo.Loseweight
	boyinfo.Totalprice = boyinfo.Singleprice * boyinfo.Loseweight
	if (boyinfo.Restweight <= 0) {
		c.Ctx.WriteString("Sorry,Xiao bi Zai has sold out ")

	} else {
		//c.Ctx.WriteString("name: "+boyinfo.Name)
		//fmt.Println(&list)
		//c.Data["Form"] = &boyinfo
		c.Data["json"] = &boyinfo
		c.ServeJSON()
		//c.TplName = "index.tpl"
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
