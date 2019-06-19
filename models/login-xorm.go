package models


import "fmt"

/*
用户登陆以及注册功能模块
登陆函数，实际上是与库内用户账户名以及密码进行查询匹配，如果匹配，则证明用户存在且密码正确
注册函数，实际上是数据库插入用户名，密码邮箱的过程，初次登陆，用户默认非会员，默认角色为普通用户
*/

/*
用户登陆：
用户id 唯一
账户密码 非空
邮箱 注册时必填 非空
是否会员 非会员 会员需要购买后生效
用户角色 默认普通用户"user"
用户电话 默认空 可更新
*/
type LoginUser struct {
	Uid int `xorm:"pk autoincr"`
	Username string
	Password string
	Email string
	Ifvip string
	Address string
	Role string
	Phone string
}

/*登陆函数*/
func FuncUserLogin(username,userpwd string) error{
	x := getDBEngine()
	m := LoginUser{}
	_,err :=x.Where("Username = ?", username).And("Password = ?",userpwd).Get(&m)
	if err ==nil{
		fmt.Println("查询成功")
	}else{
		fmt.Println("查询失败")
	}
	return err
}

/*注册函数*/
func FuncRegisterUser(username,userpwd,email string) error{
	x:= getDBEngine()
	user := LoginUser{Username:username,Password:userpwd,Email:email,Ifvip:"no",Role:"user"}
	affected, err := x.Insert(&user)
	fmt.Println(affected)
	return err
}




