package controllers

import (
	"github.com/astaxie/beego"
	"testbeego/models"
	"fmt"
	)

type LoginController struct {
	beego.Controller
}
type RegisterController struct {
	beego.Controller
}

func (c *LoginController) Prepare() {

}

func (c *LoginController) Get() {
	c.TplName = "login/login.tpl"
}

func (c *LoginController) Post() {
	userinfo := models.LoginUser{}
	userinfo.Username = c.GetString("username")
	userinfo.Password = c.GetString("password")
	//m := models.CheckAuth(userinfo)
	//m:=models.CheckAuth(userinfo.Username)
	k:=models.UserLogin(userinfo.Username,userinfo.Password)
	if ( k == nil){
		c.Redirect("/home", 301)
	} else {
		c.Redirect("/login/error", 301)
	}
	return
}

func (c *LoginController) Errorpage(){
	c.TplName = "login/errorpage.tpl"

}

func (c *RegisterController) Get(){
	c.TplName = "login/register.tpl"

}

func (c *RegisterController) Register(){
	member :=models.LoginUser{}
	member.Username = c.GetString("membername")
	member.Password = c.GetString("memberpassword")
	err :=models.RegisterUser(member.Username,member.Password)
	if (member.Username != " " && member.Password != " "){
		if (err == nil ){
			//c.Data["json"] = &member
			//c.ServeJSON()
			fmt.Println("注册成功")
			c.Redirect("/login",301)
		}else{
			fmt.Println("注册失败")
		}
	}else{
			c.Ctx.WriteString("请输入注册用户数据")
		}
}


func (c *LoginController) FilterUser(){
	_, ok := c.Ctx.Input.Session("uid").(int)
	if !ok && c.Ctx.Request.RequestURI != "/login" {
		c.Ctx.Redirect(302, "/login")
	}
}

var FilterUser = func(ctx *context.Context) {
	_, ok := ctx.Input.Session("uid").(int)
	if !ok && ctx.Request.RequestURI != "/login" {
		ctx.Redirect(302, "/login")
	}
}








