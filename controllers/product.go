package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"math"
	"testbeego/models"
)

//商品页面
type ProductController struct {
	beego.Controller
}

/*用json展示所有的商品列表*/
func (c *ProductController) ListUser() {
	k := models.BoyInfo()
	c.Data["json"] = &k
	c.ServeJSON()
}

/*查询商品（pid)*/
func (c *ProductController) SearchProduct() {
	pid, err := c.GetInt("pid")
	if err != nil {
		println("获取不到商品id")
	}
	SinglePrice, ProductInfo, has := models.SearchProduct(pid)
	if has == true {
		fmt.Println(SinglePrice)
		c.Data["json"] = &ProductInfo
		c.ServeJSON()
	} else {
		fmt.Println("商品不存在")
		c.Ctx.WriteString("您要查询的商品不存在")
	}
}

/*天猫优惠券以及折扣计算*/
func (c *ProductController) CouponCalculate() {
	//优惠券最大值-优惠券最小值（优惠券面额）
	var maxcouponvalue, mincouponvalue float64
	//津贴最大-津贴最小（津贴额度）
	var allowancemax, allowancemin float64
	//产品打折折扣
	var discount float64
	//vip折扣
	var vipdiscount float64
	//用户是否是vip会员
	var ifvip bool
	//原价
	var price float64
	//所有折扣后的价格
	var shopcart float64
	maxcouponvalue, _ = c.GetFloat("maxcouponvalue")
	mincouponvalue, _ = c.GetFloat("mincouponvalue")
	discount, _ = c.GetFloat("discount")
	allowancemax, _ = c.GetFloat("allowancemax")
	allowancemin, _ = c.GetFloat("allowancemin")
	user_id, _ := c.GetInt("uid")
	_, ifvip = models.Checkvip(user_id)
	shopcart, _ = c.GetFloat("shopcartprice")
	if shopcart < maxcouponvalue {
		if ifvip == true {
			vipdiscount = 0.95
		} else {
			vipdiscount = 1
		}
		price = shopcart*vipdiscount*discount - (shopcart/allowancemax)*allowancemin
	} else {
		if ifvip == true {
			vipdiscount = 0.95
		} else {
			vipdiscount = 1
		}
		price = shopcart*vipdiscount*discount - math.Floor(shopcart/allowancemax)*allowancemin - mincouponvalue
	}
	c.Data["json"] = &price
	c.ServeJSON()
}
