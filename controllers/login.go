package controllers

import (
	"github.com/astaxie/beego"

	"testbeego/models"
)

type LoginController struct {
	beego.Controller
}

type Logininfo struct {
	username string `form："username"`
	password string `form: "password"`
}


func (c *LoginController) Prepare() {

}

func (c *LoginController) Get() {
	c.TplName = "login/login.tpl"
}

func (c *LoginController) Post() {
	//o := orm.NewOrm()
	userinfo := Logininfo{}
	testuser := models.LoginUser{}
	userinfo.username = c.GetString("username")
	userinfo.password = c.GetString("password")
	m := models.Task(&testuser)
	println(m)
	if ( userinfo.username != "" && userinfo.password != ""){
		c.Redirect("/home", 301)
	} else {
		c.Redirect("/error", 301)
	}
	return
}
func (c *LoginController) Errorpage(){
	c.TplName = "login/errorpage.tpl"

}


