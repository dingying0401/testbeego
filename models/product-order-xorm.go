package models

import (
"fmt"
_ "github.com/go-sql-driver/mysql"
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
func SaleProduct(uid,pid int,loseweight float64) error{
	x := getDBEngine()
	productinfo := &ProductInfo{Pid: pid}
	has,err := x.Get(productinfo)
	if  err ==nil {
		if has == true{
			fmt.Println("查找的商品存在")
			fmt.Println(productinfo.Singleprice, productinfo.Weight)
		}else{
			fmt.Println("找不到该商品")
		}
	}else{
		fmt.Println("查找过程中存在错误")
	}
	var orderinfo OrderInfo
	var singleprice float64
	var weight float64
	weight = productinfo.Weight
	singleprice = productinfo.Singleprice
	orderinfo.Loseweight = loseweight
	weight = weight - loseweight
	productinfo.Weight = weight

	affected, err := x.Cols("Weight").Update(&productinfo)
	if  err == nil {
		fmt.Println(affected)
		fmt.Println("更新数值成功")
	}
	orderinfo.User_id = uid
	orderinfo.Product_id = pid
	orderinfo.Totalprice = singleprice * loseweight
	affected, result := x.Insert(&orderinfo)
	if (result == nil) {
		fmt.Println(affected)
	}
	return result
}