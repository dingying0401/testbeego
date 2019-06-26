package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go" //Golang implementation of JSON Web Tokens (JWT)
	"testbeego/models"
	"time"
)

/*登陆模块，引用beego.Controller对象*/
type LoginController struct {
	beego.Controller
}

/*注册模块*/
type RegisterController struct {
	beego.Controller
}

func (c *LoginController) Prepare() {

}

/*GET方法，进入登陆界面*/
func (c *LoginController) Get() {
	c.TplName = "login/login.tpl"
}

/*POST，输入用户名密码
如果用户名不存在，则需要注册账号
如果用户名存在且密码正确，则可以登陆成功
登陆成功后跳转到购买商品页面
*/
func (c *LoginController) Post() {
	userinfo := models.LoginUser{}
	userinfo.Username = c.GetString("username")
	userinfo.Password = c.GetString("password")
	u, _ := models.Ifuserex(userinfo.Username)
	k, _ := models.FuncUserLogin(userinfo.Username, userinfo.Password)

	if u == true && k == true {
		//生成token
		token := jwt.New(jwt.SigningMethodHS256)
		claims := make(jwt.MapClaims)
		claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
		claims["iat"] = time.Now().Unix()
		claims ["username"] = userinfo.Username
		token.Claims = claims
		//获取到tokenString
		tokenString, err := token.SignedString([]byte("mykey"))
		/*c.Data["json"] = map[string]interface{}{"status": 200, "message": "login success ", "moreinfo": tokenString}
		c.Data["json"] = tokenString
		c.ServeJSON()*/
		fmt.Printf(tokenString)
		if err != nil {
			beego.Error("jwt.SignedString:", err)
		}
		//
		c.Redirect("/home/product", 301)

	} else {
		//登陆失败
		if u == false {
			//用户名不存在
			c.Redirect("/register", 301)
		} else {
			//用户名密码校验失败
			c.Redirect("/login/error", 301)
		}

	}

}

/*登陆错误，错误页面*/
func (c *LoginController) Errorpage() {
	c.TplName = "login/errorpage.tpl"
}

/*GET注册页面*/
func (c *RegisterController) Get() {
	c.TplName = "login/register.tpl"
}

/*POST用户名密码进行用户注册*/
func (c *RegisterController) Register() {
	member := models.LoginUser{}
	//前端获取注册信息
	member.Username = c.GetString("membername")
	member.Password = c.GetString("memberpassword")
	member.Email = c.GetString("memberemail")
	//注册
	err := models.FuncRegisterUser(member.Username, member.Password, member.Email)
	if member.Username != " " && member.Password != " " {
		if err == nil {
			fmt.Println("注册成功")
			c.Redirect("/login", 301)
		} else {
			fmt.Println("注册失败")
			c.Redirect("/register/error", 301)
		}
	} else {
		c.Ctx.WriteString("请输入注册用户数据")
	}
}

func (c *RegisterController) RegistErrorpage() {
	c.TplName = "login/errorpageregister.tpl"
}