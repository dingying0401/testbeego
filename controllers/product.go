package controllers

import (
	"testbeego/models"
		"github.com/astaxie/beego"
	"fmt"
)
type ProductController struct {
	beego.Controller
}

func (c *ProductController) ListUser(){
	k:=models.BoyInfo()
	c.Data["json"] = &k
	c.ServeJSON()
	//c.Data["Form"] = &k
	//c.TplName ="index.tpl"

}

func (c *ProductController) SearchProduct(){
	pid,err:=c.GetInt("pid")
	if(err != nil){
		println("获取不到商品id")
	}
	price,productinfo,has:=models.SearchProduct(pid)
	if has == true{
		fmt.Println(price)
		c.Data["json"] = &productinfo
		c.ServeJSON()
	}else{
		fmt.Println("商品不存在")
		c.Ctx.WriteString("您要查询的商品不存在")
	}

}










