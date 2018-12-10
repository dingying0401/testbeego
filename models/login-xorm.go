package models

/*
import (
"github.com/astaxie/beego/orm"
_ "github.com/go-sql-driver/mysql"
"github.com/xormplus/xorm"
"fmt"
	"log"
)

type LoginUser2 struct {
	Uid int
	Username string
	Password string
}

func getDBEngine() *xorm.Engine {
	//set xorm engine
	var err error
	engine, err := xorm.NewMySQL("mysql", "root:dingying@tcp(10.71.200.23:3306)/test?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}

	//连接测试
	if err := engine.Ping(); err!=nil{
		log.Fatal(err)
	}

	//日志打印SQL
	engine.ShowSQL(true)

	//set SqlMap 如果要配置sql的xml文件
	/*err = engine.RegisterSqlMap(xorm.Xml("./sql/oracle", ".xml"))
	if err != nil {
		log.Fatal(err)
	}
	return engine
}


func init()  {
	//自定义表名：https://beego.me/docs/mvc/model/models.md
	orm.RegisterModel(new(LoginUser2))
}
func FuncUserLogin(username,userpwd string) error{
	x := getDBEngine()
	//o := orm.NewOrm()
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

func FuncRegisterUser(username,userpwd string) error{
	o := orm.NewOrm()
	user := LoginUser{Username:username,Password:userpwd}
	id, err := o.Insert(&user)
	fmt.Println(id)
	return err
}
*/