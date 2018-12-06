package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"fmt"
	)

type ProductInfo struct{
	Pid int `orm:"column(pid);pk"` // 设置主键
	Name string
	Length float64
	Weight float64
	Briefinfo string
	Singleprice float64
}

type OrderInfo struct {
	Oid int `orm:"column(oid);pk"` // 设置主键
	Loseweight float64
	Totalprice float64
	Product_id int
	User_id int
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

func CheckOrder(order_id int) OrderInfo {
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
	return orderdetail
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

func UpdateProductinfo(pid int,pname string,plength float64,pweight float64,pinfo string, singleprice float64) ProductInfo {
	o := orm.NewOrm()
	productinfo := ProductInfo{Pid:pid}
	err := o.Read(&productinfo)
	if (err == nil) {
		fmt.Println(productinfo.Name, productinfo.Length, productinfo.Weight, productinfo.Briefinfo, productinfo.Singleprice)

		productinfo.Name = pname
		productinfo.Length = plength
		productinfo.Weight = pweight
		productinfo.Briefinfo = pinfo
		productinfo.Singleprice = singleprice
		if num, err := o.Update(&productinfo); err == nil {
			fmt.Println(num)
		}
	}

/*
	switch {
	case (pname != ""):
		    productinfo.Name = pname
			if _, test := o.Update(&productinfo,"Name"); test == nil {
				fmt.Println("更新数值成功")
			}
		fallthrough

	case (plength != 0):
		    productinfo.Length = plength
			if _, test := o.Update(&productinfo,"Length"); test == nil {
				fmt.Println("更新数值成功")
			}

		fallthrough

	case (pweight != 0):
		productinfo.Weight = pweight
			if _, test := o.Update(&productinfo,"Weight"); test == nil {
			fmt.Println("更新数值成功")
		}
			fallthrough

	case (pinfo != ""):
		productinfo.Briefinfo = pinfo
			if _, test := o.Update(&productinfo,"Briefinfo"); test == nil {
				fmt.Println("更新数值成功")
			}

		fallthrough

	case (singleprice != 0):
		productinfo.Singleprice = singleprice
		if _, test := o.Update(&productinfo,"Singleprice"); test == nil {
			fmt.Println("更新数值成功")
		}

	}*/
	return productinfo
}




