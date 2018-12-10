package controllers

import (
	"github.com/astaxie/beego"
	)

type AdminController struct {
	beego.Controller
}
/*
func (c *AdminController) DeleteProduct(){
	product_id, _ := c.GetInt("pid")
	err := models.DeleteProduct(product_id)
	if err ==nil {
		c.Ctx.WriteString("删除商品成功")
	}else{
		c.Ctx.WriteString("删除商品失败")
	}
}
*/
/*
func (c *AdminController) UpdateProduct(){
	pid,_ := c.GetInt("pid")
	pname := c.GetString("pname")
	pweight, _ := c.GetFloat("pweight")
	plength, _ := c.GetFloat("plength")
	pinfo := c.GetString("pinfo")
	singleprice,_ := c.GetFloat("singleprice")
	result := models.UpdateProductinfo(pid,pname,plength,pweight,pinfo,singleprice)
	c.Data["json"] = &result
	c.ServeJSON()

}
*/
