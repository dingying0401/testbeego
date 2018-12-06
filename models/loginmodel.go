package models

import (
		"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

type LoginUser struct {
	Uid int `orm:"column(uid);pk"`
	Username string
	Password string
}

func init()  {
	//自定义表名：https://beego.me/docs/mvc/model/models.md
	orm.RegisterModel(new(LoginUser))
}
func UserLogin(username,userpwd string) error{
	o := orm.NewOrm()
	qs := o.QueryTable("login_user")
	m := LoginUser{}
	err := qs.Filter("Username",username).Filter("Password",userpwd).One(&m)
	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
		return err
	}
	return err
}

func RegisterUser(username,userpwd string) error{
	o := orm.NewOrm()
	user := LoginUser{Username:username,Password:userpwd}
	id, err := o.Insert(&user)
	fmt.Println(id)
	return err
}
/*func CheckAuth (data string) (error, LoginUser){
	o := orm.NewOrm()
	//user := LoginUser{Id: 1}
	user := LoginUser{Username: data}
	err := o.Read(&user,"Username")
	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
		fmt.Println(user.Id, user.Username, user.Password)
	}
	return err,user
}
*/




