package controllers

import (
	"testbeego/models"
	"fmt"
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}


func (c *UserController) VipOrder() {
	user_id, _ := c.GetInt("Uid")
	err:= models.FuncVipShop(user_id)
	if err == nil {
		fmt.Println("查询成功")
		c.Ctx.WriteString("购买会员成功")
		//c.Data["json"] = &orderdetail
		//c.ServeJSON()
		//if productname == ""{c.Ctx.WriteString("商品不存在，已失效")}
	} else {
		fmt.Println("查询失败")
		c.Ctx.WriteString("购买失败")
	}
}

func (c *UserController) UpdateUserInfo(){
	user_id, _ := c.GetInt("Uid")
	email := c.GetString("email")
	address := c.GetString("address")
	phone := c.GetString("phone")
	k :=models.EditUserinfo(user_id,email,address,phone)
	c.Data["json"] = &k
	c.ServeJSON()

}

