package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type LoginInfo struct {
	Username string `form："username"`
	Password string `form: "password"`
}

type LoginUser struct {
	Id int
	Username string
	Password string
}


/*func RegisterDB() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(LoginUser))
}*/
func init()  {
	orm.RegisterModel(new(LoginUser))
}

func CheckAuth (data LoginInfo) error{
	o := orm.NewOrm()
	user := LoginUser{Id: 1}

	err := o.Read(&user)
	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
		fmt.Println(user.Id, user.Username)
	}
	return err
}

