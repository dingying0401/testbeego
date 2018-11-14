package controllers

import (
	"github.com/astaxie/beego"

	"testbeego/models"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Prepare() {

}

func (c *LoginController) Get() {
	c.TplName = "login/login.tpl"
}

func (c *LoginController) Post() {
	//o := orm.NewOrm()
	userinfo := models.LoginInfo{}
	userinfo.Username = c.GetString("username")
	userinfo.Password = c.GetString("password")
	m := models.CheckAuth(userinfo)
	if ( m == nil){
		c.Redirect("/home", 301)
	} else {
		c.Redirect("/error", 301)
	}
	return
}
func (c *LoginController) Errorpage(){
	c.TplName = "login/errorpage.tpl"

}


