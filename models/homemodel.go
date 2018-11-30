package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
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

func RegisterDB4UserInfo() {
	//注册定义的model
	orm.RegisterModel(new(ProductInfo),new(OrderInfo))
}

/*
func BoyInfo() *ProductInfo{
	o := orm.NewOrm()
	qs := o.QueryTable("product_info")
	err := qs.
	if err == nil {
		fmt.Printf("Result Nums: %d\n", num)
		for _, m := range maps {
			fmt.Println(m["Id"], m["Name"])
		}
	}
	return err
}
*/
