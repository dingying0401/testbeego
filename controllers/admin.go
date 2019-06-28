package controllers

import (
	"github.com/astaxie/beego"
	"testbeego/models"
)

type AdminController struct {
	beego.Controller
}

/*删除商品*/
func (c *AdminController) DeleteProduct() {
	product_id, _ := c.GetInt("pid")
	err := models.DeleteProduct(product_id)
	if err == nil {
		c.Ctx.WriteString("删除商品成功")
	} else {
		c.Ctx.WriteString("删除商品失败")
	}
}

/*更新商品信息*/
func (c *AdminController) UpdateProduct() {
	pid, _ := c.GetInt("pid")
	pname := c.GetString("pname")
	pweight, _ := c.GetFloat("pweight")
	plength, _ := c.GetFloat("plength")
	pinfo := c.GetString("pinfo")
	singleprice, _ := c.GetFloat("singleprice")
	result := models.UpdateProductInfo(pid, pname, plength, pweight, pinfo, singleprice)
	c.Data["json"] = &result
	c.ServeJSON()

}

/*上架新的商品*/
func (c *AdminController) AddProduct() {
	pname := c.GetString("pname")
	pweight, _ := c.GetFloat("pweight")
	plength, _ := c.GetFloat("plength")
	pinfo := c.GetString("pinfo")
	singleprice, _ := c.GetFloat("singleprice")
	result := models.FuncAddProduct(pname, pinfo, plength, pweight, singleprice)
	c.Data["json"] = &result
	c.ServeJSON()

}