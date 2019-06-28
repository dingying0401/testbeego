package controllers

import (
	"github.com/astaxie/beego"
	"testbeego/models"
)

type UserController struct {
	beego.Controller
}

/*购买会员，成为vip用户*/
func (c *UserController) VipOrder() {
	user_id, _ := c.GetInt("Uid")
	err := models.FuncVipShop(user_id)
	if err == nil {
		c.Ctx.WriteString("购买会员成功")
	} else {
		c.Ctx.WriteString("购买失败")
	}
}

/*更新用户信息，开始注册默认只有邮箱，用户名信息，其他信息默认为空*/
func (c *UserController) UpdateUserInfo() {
	user_id, _ := c.GetInt("Uid")
	email := c.GetString("email")
	address := c.GetString("address")
	phone := c.GetString("phone")
	k := models.EditUserinfo(user_id, email, address, phone)
	c.Data["json"] = &k
	c.ServeJSON()
}
