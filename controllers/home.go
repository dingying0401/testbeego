package controllers

import (
	"github.com/astaxie/beego"
	"log"
	"testbeego/models"
)

type MainController struct {
	beego.Controller
}



func (c *MainController) Prepare() {

}
func (c *MainController) Evaluate() {
	boyinfo := &models.Userinfo{}
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
