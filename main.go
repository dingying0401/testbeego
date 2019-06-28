package main

import (
	"encoding/json"
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/dgrijalva/jwt-go" //Golang implementation of JSON Web Tokens (JWT)
	"strings"
	_ "testbeego/routers"
)

/*jwt服务端*/
var FilterUser = func(ctx *context.Context) {
	//获取request body中的内容
	buf := make([]byte, 1024)
	n, _ := ctx.Request.Body.Read(buf)
	m := ctx.Input.Query("username")
	var requestBody = string(buf[0:n])
	var username interface{}
	if requestBody != "" {
		//将body内容转换成map对象
		var requestMap map[string]interface{}

		//json解析map
		err := json.Unmarshal([]byte(requestBody), &requestMap)
		if (err != nil) {
			beego.Error(err)
			ctx.Redirect(302, "/login")
			return
		}
		//获取request中的username
		username = requestMap["username"]
	} else {
		username = m
	}
	//get token内容
	authString := ctx.Input.Header("Authorization")
	beego.Debug("AuthString:", authString)

	kv := strings.Split(authString, " ")
	if len(kv) != 2 || kv[0] != "Bearer" {
		beego.Error("AuthString invalid:", authString)
		ctx.WriteString("login failed")
		ctx.Redirect(302, "/login")
		return
	}
	tokenString := kv[1]

	// Parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("mykey"), nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				// That's not even a token
				beego.Error("Parse token failed:", err)
				beego.Error("Parse token:", err)
				ctx.Redirect(302, "/login")
				return
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				// Token is either expired or not active yet
				beego.Error("token expired:", err)
				ctx.Redirect(302, "/login")
				return
			} else {
				// Couldn't handle this token
				beego.Error("failed:", err)
				ctx.Redirect(302, "/login")
				return
			}
		} else {
			// Couldn't handle this token
			beego.Error("failed:", err)
			ctx.Redirect(302, "/login")
			return
		}
	}
	if !token.Valid {
		beego.Error("Token invalid:", tokenString)
		ctx.Redirect(302, "/login")
		return
	}
	beego.Debug("Token:", token)

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		errors.New("errPermission")
		ctx.Redirect(302, "/login")
		return
	}
	var user string = claims["username"].(string)
	if user != username {
		beego.Error("no auth to operate other user's content")
		ctx.Redirect(302, "/login")
		return
	}

}

func init() {
	/*
		orm.RegisterDriver("mysql", orm.DRMySQL)
		//orm.RegisterDataBase("default", "mysql", "root:dingying@/test?charset=utf8")
		orm.RegisterDataBase("default", "mysql", "root:dingying@tcp(127.0.0.1:3306)/test?charset=utf8")
	*/
}

func main() {
	//beego.SetStaticPath("/static","public")
	//orm.Debug = true
	//orm.RunSyncdb("default",true, true)
	//open session
	beego.BConfig.WebConfig.Session.SessionOn = true
	//filter
	beego.InsertFilter("/home/shop", beego.BeforeRouter, FilterUser)
	//beego.InsertFilter("/user/*", beego.BeforeRouter, FilterUser)
	//beego.InsertFilter("/login/?:id", beego.BeforeRouter, FilterUser)
	//beego.InsertFilter("/admin/*", beego.BeforeRouter, FilterUser)
	beego.InsertFilter("/test/token", beego.BeforeRouter, FilterUser)
	beego.Run()
	/*
		1.app.conf
		2.hookfunc (函数：AddAPPStartHook)
		3.session or not
		4.是否编译模板：beego会在启动时候根据配置把views目录下的模板进行预编译，然后存在map里面，这样可以有效提高模板运行的效率，无须多次编译
		5.是否开启文档功能：根据 EnableDocs 配置判断是否开启内置的文档路由功能
		6.是否启动管理模块：EnableAdmin
		7.监听服务端口：ListenAndServer
	*/
}
