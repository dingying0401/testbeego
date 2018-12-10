package models


import (
_ "github.com/go-sql-driver/mysql"
"fmt"
)

type LoginUser struct {
	Uid int `orm:"column(uid);pk"`
	Username string
	Password string
}
func init()  {

}
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

func FuncRegisterUser(username,userpwd string) error{
	//o := orm.NewOrm()
	x:= getDBEngine()
	user := LoginUser{Username:username,Password:userpwd}
	affected, err := x.Insert(&user)
	fmt.Println(affected)
	return err
}

