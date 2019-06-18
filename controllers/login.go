package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
	"testbeego/models"
	"time"
)

//登陆模块，引用beego.Controller对象
type LoginController struct {
	beego.Controller
}

//注册模块
type RegisterController struct {
	beego.Controller
}

func (c *LoginController) Prepare() {

}

//GET方法，进入登陆界面
func (c *LoginController) Get() {
	c.TplName = "login/login.tpl"
}

//POST，输入用户名密码，并经过jwt验证
//好像没有验证？如果库里没有，则自动跳转到注册页面？？还是用户名密码错误了，跳转到提示
func (c *LoginController) Post() {
	userinfo := models.LoginUser{}
	userinfo.Username = c.GetString("username")
	userinfo.Password = c.GetString("password")
	//m := models.CheckAuth(userinfo)
	//m:=models.CheckAuth(userinfo.Username)
	k := models.FuncUserLogin(userinfo.Username, userinfo.Password)
	if ( k == nil){
		token := jwt.New(jwt.SigningMethodHS256)
		claims := make(jwt.MapClaims)
		claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
		claims["iat"] = time.Now().Unix()
		claims ["username"] = userinfo.Username
		token.Claims = claims
		//获取到tokenString
		tokenString, err := token.SignedString([]byte("mykey"))
		//c.Data["json"] = map[string]interface{}{"status": 200, "message": "login success ", "moreinfo": tokenString}
		c.Data["json"] = tokenString
		c.ServeJSON()
		if err != nil {
			beego.Error("jwt.SignedString:", err)
		}
		/*
		//带权限创建令牌
		v := c.GetSession("authResult")
		if(v == nil){
			c.SetSession("authResult", "success")
		}
		*/
		c.Redirect("/home/product", 301)
	} else {
		c.Redirect("/login/error", 301)
	}
	return

}

//登陆错误
func (c *LoginController) Errorpage(){
	c.TplName = "login/errorpage.tpl"

}

//注册页面
func (c *RegisterController) Get(){
	c.TplName = "login/register.tpl"

}

//POST用户名密码进行注册用户
func (c *RegisterController) Register(){
	member :=models.LoginUser{}
	member.Username = c.GetString("membername")
	member.Password = c.GetString("memberpassword")
	member.Email = c.GetString("memberemail")
	err :=models.FuncRegisterUser(member.Username,member.Password,member.Email)
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












