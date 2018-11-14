package controllers

import (
	"github.com/astaxie/beego"

	"testbeego/models"
)

type LoginController struct {
	beego.Controller
}

/*type Logininfo struct {
	username string `form："username"`
	password string `form: "password"`
}
*/

func (c *LoginController) Prepare() {

}

func (c *LoginController) Get() {
	c.TplName = "login/login.tpl"
}

func (c *LoginController) Post() {
	//o := orm.NewOrm()
	userinfo := models.LoginUser{}
	userinfo.Username = c.GetString("username")
	userinfo.Password = c.GetString("password")
	if ( userinfo.Username == models.Task(&userinfo) ){
		c.Redirect("/home", 301)
	} else {
		c.Redirect("/error", 301)
	}
	return
}
func (c *LoginController) Errorpage(){
	c.TplName = "login/errorpage.tpl"

}


