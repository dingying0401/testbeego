package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"testbeego/models"

)

type OrderController struct {
	beego.Controller
}

func (c *OrderController) HandleOrder() {
	user_id, _ := c.GetInt("uid")
	loseweight, _ := c.GetFloat("loseweight")
	product_id, _ := c.GetInt("pid")
	result := models.SaleProduct(user_id, product_id, loseweight)
	if (result == nil) {
		fmt.Println("订单数据插入成功")
		c.Redirect("/home/product", 301)
	} else {
		fmt.Println("订单数据插入失败")
	}

}

func (c *OrderController) SearchOrder(){
	user_id, _ := c.GetInt("uid")
	orderlist,err := models.ListOrder(user_id)
	if (err ==nil){
		c.Data["json"] = &orderlist
		c.ServeJSON()
	}else {
		c.Ctx.WriteString("抱歉查询不到您的订单")
	}
}

func (c *OrderController) CheckOrder(){
    order_id, _ := c.GetInt("oid")
    orderdetail :=models.CheckOrder(order_id)
	c.Data["json"] = &orderdetail
	c.ServeJSON()
}

func (c *OrderController) DeleteOrder(){
	order_id, _ := c.GetInt("oid")
	err := models.DeleteOrder(order_id)
	if err ==nil {
		c.Ctx.WriteString("删除订单成功")
	}else{
		c.Ctx.WriteString("删除订单失败")
	}
}


