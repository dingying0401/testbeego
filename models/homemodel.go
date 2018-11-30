package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"fmt"
	)

type ProductInfo struct{
	Id int
	Name string
	Length float32
	Weight float64
	Briefinfo string
	Singleprice float64
}

type OrderInfo struct {
	Id int
	Loseweight float64
	Restweight float64
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
			fmt.Println(m["Id"], m["Name"])

		}
	}
	return &maps
}

