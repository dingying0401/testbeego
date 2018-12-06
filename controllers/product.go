package controllers

import (
	"testbeego/models"
	"fmt"
	"github.com/astaxie/beego"
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
	id,err:=c.GetInt("id")
	if(err != nil){
		println(err)
	}
	pname,productinfo:=models.SearchProduct(id)
	fmt.Println(pname)
	c.Data["json"] = &productinfo
	c.ServeJSON()
}







