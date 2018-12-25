package controllers

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
	"strings"
	"testbeego/models"
	"time"
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

func (c *LoginController) ParseToken() (t *jwt.Token, e error) {
	authString := c.Ctx.Input.Header("Authorization")
	beego.Debug("AuthString:", authString)

	kv := strings.Split(authString, " ")
	if len(kv) != 2 || kv[0] != "Bearer" {
		beego.Error("AuthString invalid:", authString)
		test := errors.New("prase token failed")
		c.Ctx.WriteString("login failed")
		return nil,test
	}
	tokenString := kv[1]

	// Parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("mykey"), nil
	})
	if err != nil {
		beego.Error("Parse token:", err)
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				// That's not even a token
				return nil, err
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				// Token is either expired or not active yet
				return nil, err
			} else {
				// Couldn't handle this token
				return nil, err
			}
		} else {
			// Couldn't handle this token
			return nil, err
		}
	}
	if !token.Valid {
		beego.Error("Token invalid:", tokenString)
		flag := errors.New("token invalid")
		return nil, flag
	}
	beego.Debug("Token:", token)
	c.Ctx.WriteString("login success")
	return token, nil

}












