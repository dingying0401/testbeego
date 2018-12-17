package models

import (
"fmt"
_ "github.com/go-sql-driver/mysql"
	"errors"
	"github.com/xormplus/xorm"
	"log"
	"time"
)
type ProductInfo struct{
	Pid int `xorm:"pk autoincr"`//`orm:"column(pid);pk"` // 设置主键*/
	Name string
	Length float64
	Weight float64
	Briefinfo string
	Singleprice float64
}

type OrderInfo struct {
	Oid int  `xorm:"pk autoincr"` //`orm:"column(oid);pk"` // 设置主键
	Loseweight float64
	Totalprice float64
	Product_id int
	User_id int
	Updatetime time.Time `xorm:"updated"`
}

type OrderDetail struct {
	OrderInfo `xorm:"extends"`
	Name string
	Weight string
	Username string
	//ProductInfo `xorm:"extends"`
	//LoginUser `xorm:"extends"`
}

func getDBEngine() *xorm.Engine {
	//set xorm engine
	var err error
	engine, err := xorm.NewMySQL("mysql", "root:dingying@tcp(10.71.200.21:3306)/test?charset=utf8&loc=Asia%2FShanghai")
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
	}*/
	return engine
}

func BoyInfo() map[int64]ProductInfo {
	x := getDBEngine()
	users := make(map[int64]ProductInfo)
	err := x.Find(&users)
	if err ==nil{
		fmt.Println("有商品列表")
	}else {
		fmt.Println("无商品列表")
	}
	return users
}

func SearchProduct(pid int) (float64, *ProductInfo,bool) {
	x := getDBEngine()
	user := &ProductInfo{Pid:pid}
	has, err := x.Get(user)
	if err ==nil {
		if has == true{
			fmt.Println("查找的商品存在")
		}else{
			fmt.Println("找不到该商品")
		}
	}else{
		fmt.Println("查找过程中存在错误")
	}

	return user.Singleprice,user,has
}
func SaleProduct(uid,pid int,loseweight float64) error {
	x := getDBEngine()
	var flag error
	productinfo := ProductInfo{Pid: pid}
	has, err := x.Get(&productinfo)
	if err == nil {
		if has == true {
			fmt.Println("查找的商品存在")
			fmt.Println(productinfo.Singleprice, productinfo.Weight)
		} else {
			fmt.Println("找不到该商品")
		}
	} else {
		fmt.Println("查找过程中存在错误")
	}
	var orderinfo OrderInfo
	var singleprice float64
	var weight float64
	weight = productinfo.Weight
	singleprice = productinfo.Singleprice
	orderinfo.Loseweight = loseweight
	if (weight >= 0 && loseweight <= weight) {
		weight = weight - loseweight
		productinfo.Weight = weight
		fmt.Println(productinfo.Weight)
		affected, errs := x.Id(pid).Cols("weight").Update(&productinfo)
		if errs == nil {
			flag = errs
			fmt.Println(affected)
			fmt.Println("更新数值成功")
		} else {
			fmt.Println("更新数值失败")
		}
		orderinfo.User_id = uid
		orderinfo.Product_id = pid
		orderinfo.Totalprice = singleprice * loseweight
		affected, result := x.Insert(&orderinfo)
		if (result == nil) {
			flag = result
			fmt.Println(affected)
			return flag
		}
	} else {
		flag = errors.New("low weight alarm")
		fmt.Println("剩余产品库存不足，请选择合适的购买量")
	}
	return flag
}


func ListOrder(user_id int) (map[int64]OrderInfo,error){
	x := getDBEngine()
	orderinfo := make(map[int64]OrderInfo)
	err := x.Where("User_id = ?",user_id).Find(&orderinfo)
	if err ==nil{
		fmt.Println("查询成功")
		fmt.Println(orderinfo)
	}else {
		fmt.Println("查询失败")
	}
	return orderinfo,err
}

func DeleteOrder(order_id int) error{
	x := getDBEngine()
	oderdetail := OrderInfo{}
	affected, err := x.Id(order_id).Delete(&oderdetail)
	if err == nil {
		fmt.Println(affected)
		fmt.Println("删除订单成功")
	}else{
		fmt.Println("删除订单失败")
	}
	return err
}

func DeleteProduct(product_id int) error{
	x := getDBEngine()
	productdetail := ProductInfo{}
	affected, err := x.Id(product_id).Delete(&productdetail)
	if err == nil {
		fmt.Println(affected)
		fmt.Println("删除商品成功")
	}else{
		fmt.Println("删除商品失败")
	}
	return err
}
func UpdateProductinfo(pid int,pname string,plength float64,pweight float64,pinfo string, singleprice float64) *ProductInfo {
	x := getDBEngine()
	productinfo := new(ProductInfo)
	has, err := x.Where("pid=?", pid).Get(productinfo)
	if (has && err == nil) {
		fmt.Println(productinfo.Name, productinfo.Length, productinfo.Weight, productinfo.Briefinfo, productinfo.Singleprice)
		productinfo.Name = pname
		productinfo.Length = plength
		productinfo.Weight = pweight
		productinfo.Briefinfo = pinfo
		productinfo.Singleprice = singleprice
		affected,err := x.Where("pid=?", pid).Update(productinfo)
		if err == nil{
			fmt.Print(affected)
			fmt.Println("更新成功")
		}
	}
	num,result := x.Where("pid=?", pid).Get(productinfo)
	if result == nil{
		fmt.Println(num)
		fmt.Println(productinfo.Name, productinfo.Length, productinfo.Weight, productinfo.Briefinfo, productinfo.Singleprice)
	}
	return productinfo
}


func CheckOrder(order_id int) (error,[]OrderDetail) {
	x := getDBEngine()
	orderdetail := OrderInfo{}
	has,err := x.Id(order_id).Get(&orderdetail)
	if err == nil {
		fmt.Println(has)
		fmt.Println("可以查询")
	} else {
		fmt.Println(orderdetail.Oid)
	}
	orderdetailo := make([]OrderDetail,0)
	errs := x.Table(&OrderInfo{}).Join("LEFT", "product_info", "order_info.product_id=product_info.pid").Join("LEFT","login_user","order_info.user_id = login_user.Uid").Where("oid = ?",orderdetail.Oid).Find(&orderdetailo)
	//sql := `SELECT oid, login_user.username,product_info.name ,loseweight, totalprice from order_info left join product_info on order_info.product_id=product_info.pid left join login_user on order_info.user_id = login_user.Uid where oid = ?`
	//errs := x.SQL("SELECT oid, login_user.username,product_info.name ,loseweight, totalprice from order_info left join product_info on order_info.product_id=product_info.pid left join login_user on order_info.user_id = login_user.Uid where oid = ?",orderdetail.Oid).Find(&orderdetailo)
	return errs,orderdetailo
}

