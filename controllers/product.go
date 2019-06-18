package controllers

import (
	"testbeego/models"
	"github.com/astaxie/beego"
	"fmt"
	"math"
)

//商品页面
type ProductController struct {
	beego.Controller
}

func (c *ProductController) ListUser(){
	k := models.BoyInfo()
	c.Data["json"] = &k
	c.ServeJSON()
	//c.Data["Form"] = &kgi
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

func (c *ProductController) CouponCalculate(){
	var maxcouponvalue,mincouponvalue float64
	var allowancemax,allowancemin float64
	var discount float64
	var vipdiscount float64
	var ifvip string
	var price float64
	var shopcart float64
	maxcouponvalue,_ = c.GetFloat("maxcouponvalue")
	mincouponvalue,_ = c.GetFloat("mincouponvalue")
	discount,_ = c.GetFloat("discount")
	allowancemax,_ = c.GetFloat("allowancemax")
	allowancemin,_ = c.GetFloat("allowancemin")
	user_id,_ := c.GetInt("uid")
	_, ifvip = models.Checkvip(user_id)
	shopcart,_ = c.GetFloat("shopcartprice")
	if (shopcart < maxcouponvalue){
		if (ifvip =="yes"){
			vipdiscount = 0.95
			}else{
				vipdiscount = 1
			}
			price = shopcart*vipdiscount*discount -(shopcart/allowancemax)*allowancemin
		}else{
		if (ifvip =="yes"){
			vipdiscount = 0.95
		}else{
			vipdiscount = 1
		}
		price = shopcart*vipdiscount*discount-math.Floor(shopcart/allowancemax)*allowancemin-mincouponvalue
	}
	c.Data["json"] = &price
	c.ServeJSON()
}









