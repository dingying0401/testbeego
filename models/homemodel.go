package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
	"log"
)

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
	}*/
	return engine
}

type ProductInfo struct{
	Pid int `orm:"column(pid);pk"` // 设置主键*/
	Name string
	Length float64
	Weight float64
	Briefinfo string
	Singleprice float64
}

type OrderInfo struct {
	Oid int `orm:"column(oid);pk"` // 设置主键*/
	Loseweight float64
	Totalprice float64
	Product_id int
	User_id int
}

type OrderDetail struct {
	Oid int
	ProductName string
	Loseweight float64
	Totalprice float64
}


func init() {
	//注册定义的model
	orm.RegisterModel(new(ProductInfo),new(OrderInfo))
}

func BoyInfo() *[]orm.Params {
	var maps []orm.Params
	o := orm.NewOrm()
	qs := o.QueryTable("product_info")
	num,err := qs.Values(&maps)
	if err == nil {
		fmt.Printf("Result Nums: %d\n", num)
		for _, m := range maps {
			fmt.Println(m["Pid"], m["Name"])
		}
	}
	return &maps
}
func SearchProduct(id int) (float64, ProductInfo)  {
	o := orm.NewOrm()
	product := ProductInfo{Pid: id}
	err := o.Read(&product)
	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
		fmt.Println(product.Pid, product.Name)

	}
	return product.Singleprice,product
}

func SaleProduct(uid,pid int,loseweight float64) error{
	o := orm.NewOrm()
	productinfo := ProductInfo{Pid: pid}
	err := o.Read(&productinfo)
	if (err ==nil){
		fmt.Println(productinfo.Singleprice, productinfo.Weight)
	}else{
		fmt.Println(err)
	}

	var orderinfo OrderInfo
	var singleprice float64
	var weight float64
	weight = productinfo.Weight
	singleprice = productinfo.Singleprice
	orderinfo.Loseweight = loseweight
	weight = weight - loseweight
	productinfo.Weight = weight
	if num, test := o.Update(&productinfo,"Weight"); test == nil {
		fmt.Println(num)
		fmt.Println("更新数值成功")
	}
	orderinfo.User_id = uid
	orderinfo.Product_id = pid
	orderinfo.Totalprice = singleprice * loseweight
	id, result := o.Insert(&orderinfo)
	if (result == nil) {
		fmt.Println(id)
	}
	return result
}

func ListOrder(user_id int) (*[]OrderInfo,error){
	o := orm.NewOrm()
	var orderinfo []OrderInfo
	num, err := o.QueryTable("order_info").Filter("User_id", user_id).All(&orderinfo)
	if err == nil {
		fmt.Printf("Result Nums: %d\n", num)
	}else{
		fmt.Printf("查询不到")
	}
	return &orderinfo,err
}

func CheckOrder(order_id int) (error,OrderDetail) {
	o := orm.NewOrm()
	orderdetail := OrderInfo{Oid:order_id}
	err := o.Read(&orderdetail)
	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
		fmt.Println(orderdetail.Oid)

	}

	var orderdetailo OrderDetail
	//result,ree:= o.Raw("SELECT oid, login_user.username,product_info.name as productName ,loseweight, totalprice from order_info left join product_info on order_info.product_id=product_info.pid left join login_user on order_info.user_id = login_user.Uid where oid = ?",orderdetail.Oid).Exec()
	result:= o.Raw("SELECT oid, product_info.name as productName ,loseweight, totalprice from order_info left join product_info on order_info.product_id=product_info.pid where oid = ?",orderdetail.Oid).QueryRow(&orderdetailo)
	/*if ree ==nil{
			num, _ := result.RowsAffected()
			fmt.Println("mysql row affected nums: ", num)
		}
	*/
	return result,orderdetailo

	//return orderdetail
}

func DeleteOrder(order_id int) error{
	o := orm.NewOrm()
	oderdetail := OrderInfo{Oid:order_id}
	num, err := o.Delete(&oderdetail)
	if err == nil {
		fmt.Println(num)
		fmt.Println("删除订单成功")
	}else{
		fmt.Println("删除订单失败")
	}
	return err
}


func DeleteProduct(product_id int) error{
	o := orm.NewOrm()
	productdetail := ProductInfo{Pid:product_id}
	num, err := o.Delete(&productdetail)
	if err == nil {
		fmt.Println(num)
		fmt.Println("删除商品成功")
	}else{
		fmt.Println("删除商品失败")
	}
	return err
}

func UpdateProductinfo(pid int,pname string,plength float64,pweight float64,pinfo string, singleprice float64) *ProductInfo {
	/*o := orm.NewOrm()
	productinfo := ProductInfo{Pid:pid}*/
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
		/*if num, err := o.Update(&productinfo); err == nil {
			fmt.Println(num)
		}*/

		affected, err := x.Where("pid=?", pid).Update(productinfo)
		if err == nil{
			fmt.Print(affected)
		}
	}
	return productinfo
}




